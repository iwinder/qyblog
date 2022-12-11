package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
)

type CommentDO struct {
	metaV1.ObjectMeta
	AgentId        uint64
	MemberId       uint64
	AtMemberIds    string
	Agent          string
	MemberName     string
	Ip             string
	Email          string
	Url            string
	RootId         uint64
	ParentId       uint64
	Content        string
	Meta           string
	ParentUserName string
	ObjTitle       string
	ObjLink        string
	Avatar         string
}

type CommentDOList struct {
	metaV1.ListMeta
	Agent *CommentAgentDO
	Items []*CommentContentDO
}

type CommentUsecase struct {
	log *log.Helper
	au  *ArticleUsecase
	ca  *CommentAgentUsecase
	ci  *CommentIndexUsecase
	cc  *CommentContentUsecase
	uu  *UserUsecase
}

func NewCommentUsecase(logger log.Logger, ca *CommentAgentUsecase,
	ci *CommentIndexUsecase, uu *UserUsecase, au *ArticleUsecase,
	cc *CommentContentUsecase,
) *CommentUsecase {
	return &CommentUsecase{log: log.NewHelper(logger),
		au: au, ca: ca, ci: ci, cc: cc, uu: uu}
}

// CreateComment 新增评论
func (uc *CommentUsecase) CreateComment(ctx context.Context, g *CommentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateComment: %v-%v", g.MemberId, g.MemberName)
	// 如果有父级，先查询父级
	g.RootId = 0
	if g.ParentId > 0 {
		parent, err := uc.ci.FindByID(ctx, g.ParentId)
		if err != nil {
			log.Error(err)
		}
		// 其父级的根目录为0，则父级为根评论
		if parent.RootId == 0 {
			g.RootId = parent.ID
			// TODO: 计算改为定时任务
		} else {
			g.RootId = parent.RootId
		}
	}

	// 创建index
	ci := &CommentIndexDO{
		AgentId:  g.AgentId,
		MemberId: g.MemberId,
		RootId:   g.RootId,
		ParentId: g.ParentId,
	}
	ci.StatusFlag = g.StatusFlag
	cidata, cierr := uc.ci.CreateCommentIndex(ctx, ci)
	if cierr != nil {
		log.Error(cierr)
	}
	// 创建内容
	cc := &CommentContentDO{
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: "",
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      g.RootId,
		Content:     g.Content,
	}
	cc.StatusFlag = g.StatusFlag
	if cidata != nil {
		cc.ID = cidata.ID
	}
	data, err := uc.cc.CreateCommentContent(ctx, cc)
	if err != nil {
		return nil, err
	}
	g.ID = data.ID
	if cc.StatusFlag == 1 {
		// 更新计数
		uc.ca.UpdateAddCountById(ctx, g.AgentId, g.RootId == 0)
		if g.ParentId > 0 {
			uc.ci.UpdateAddCountById(ctx, g.ParentId, g.RootId == 0)
		}
	} else {
		// TODO：推送待审核消息
	}
	return data, nil
}

// UpdateCommentComent 更新评论内容
func (uc *CommentUsecase) UpdateCommentComent(ctx context.Context, g *CommentDO) error {
	return uc.cc.UpdaeCommentById(ctx, g.ID, g.Content)
}

// UpdateCommentState 更新状态
func (uc *CommentUsecase) UpdateCommentState(ctx context.Context, ids []uint64, state int) error {
	err := uc.cc.UpdateStateByIDs(ctx, ids, state)
	if err != nil {
		return err
	}
	err = uc.ci.UpdateStateByIDs(ctx, ids, state)
	if err != nil {
		log.Error(err)
	}

	for _, id := range ids {
		//
		idx, _ := uc.ci.FindByID(ctx, id)
		if idx != nil {
			// 更新
			if state == 1 {
				uc.UpdateAddCount(ctx, idx)
			} else {
				uc.UpdateMinusCount(ctx, idx)
			}
		}
	}
	return nil
}
func (uc *CommentUsecase) UpdateAddCount(ctx context.Context, idx *CommentIndexDO) {
	uc.ca.UpdateAddCountById(ctx, idx.AgentId, idx.RootId == 0)
	if idx.ParentId > 0 {
		uc.ci.UpdateAddCountById(ctx, idx.ParentId, idx.RootId == 0)
	}
}
func (uc *CommentUsecase) UpdateMinusCount(ctx context.Context, idx *CommentIndexDO) {
	uc.ca.UpdateMinusCountById(ctx, idx.AgentId, idx.RootId == 0)
	if idx.ParentId > 0 {
		uc.ci.UpdateMinusCountById(ctx, idx.ParentId, idx.RootId == 0)
	}
}
func (uc *CommentUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	err := uc.cc.DeleteList(ctx, ids)
	if err != nil {
		return err
	}
	err = uc.ci.DeleteList(ctx, ids)
	if err != nil {
		log.Error(err)
	}
	for _, id := range ids {
		idx, _ := uc.ci.FindByID(ctx, id)
		if idx != nil {
			// 更新
			uc.UpdateMinusCount(ctx, idx)
		}
	}
	return err
}
func (uc *CommentUsecase) ListAllForWeb(ctx context.Context, opts CommentContentDOListOption) *CommentDOList {
	result := &CommentDOList{Agent: &CommentAgentDO{Count: 0}}
	if !opts.IsChild && opts.Current == 1 {
		agent, err := uc.ca.FindByID(ctx, opts.AgentId)
		if err != nil {
			log.Error(err)
			result.Agent = &CommentAgentDO{}
			result.Items = make([]*CommentContentDO, 0, 0)
			return result
		}
		result.Agent = agent
	}
	objList, err := uc.cc.ListAll(ctx, opts)
	if err != nil {
		log.Error(err)
		result.Items = make([]*CommentContentDO, 0, 0)
		return result
	}
	result.Items = objList.Items
	result.ListMeta = objList.ListMeta
	return result
}
