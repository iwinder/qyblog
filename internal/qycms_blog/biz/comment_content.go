package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"gorm.io/gorm"
)

type CommentContentDO struct {
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
	Content        string
	Attrs          string
	ParentUserName string
	ObjTitle       string
	ObjLink        string
	Avatar         string
	Count          int32
	RootCount      int32
	EmailState     int32
	Children       []*CommentContentDO
}

type CommentCount struct {
	Total        int64 // 总共
	NowTotal     int64 // 当前已审核
	PendingTotal int64 // 待审核

}
type CommentContentDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CommentContentDO `json:"items"`
}

type CommentContentDOListOption struct {
	metaV1.ListOptions `json:"page"`
	CommentContentDO   `json:"item"`
	IsWeb              bool
	IsChild            bool
}

// CommentAgentRepo is a Greater repo.
type CommentContentRepo interface {
	Save(context.Context, *CommentContentDO) (*CommentContentDO, error)
	Update(context.Context, *CommentContentDO) (*CommentContentDO, error)
	UpdaeStateByIDs(context.Context, []uint64, int) error
	UpdaeCommentById(context.Context, uint64, string) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*CommentContentDO, error)
	FindParentByID(context.Context, uint64) (*CommentContentDO, error)
	ListAll(context.Context, CommentContentDOListOption) (*CommentContentDOList, error)
	CountByState(context.Context, int) (int64, error)
	UpdaeEmailStateById(context.Context, uint64, int32) error
	FindAllByParentID(context.Context, uint64, int) ([]*CommentContentDO, error)
}

// CommentContentUsecase is a CommentAgentDO usecase.
type CommentContentUsecase struct {
	repo CommentContentRepo
	log  *log.Helper
	au   *ArticleUsecase
	ci   *CommentIndexUsecase
	uu   *UserUsecase
}

// NewCommentContentUsecase new a CommentContentDO usecase.
func NewCommentContentUsecase(repo CommentContentRepo, logger log.Logger,
	au *ArticleUsecase, ci *CommentIndexUsecase, uu *UserUsecase,
) *CommentContentUsecase {
	return &CommentContentUsecase{repo: repo, log: log.NewHelper(logger),
		au: au, ci: ci, uu: uu}
}

// CreateCommentContent creates a CommentContentDO, and returns the new CommentContentDO.
func (uc *CommentContentUsecase) CreateCommentContent(ctx context.Context, g *CommentContentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateCommentContent: %v-%v", g.MemberId, g.MemberName)
	data, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	g.ID = data.ID
	//uc.ci.CreateCommentIndexByContent(ctx, g)
	return data, nil
}

// Update 更新
func (uc *CommentContentUsecase) Update(ctx context.Context, g *CommentContentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.MemberId, g.MemberName)
	data, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return data, nil
}
func (uc *CommentContentUsecase) UpdateStateByIDs(ctx context.Context, ids []uint64, state int) error {
	uc.log.WithContext(ctx).Infof("UpdaeStateByIDs: %v", ids)
	err := uc.repo.UpdaeStateByIDs(ctx, ids, state)
	return err
}
func (uc *CommentContentUsecase) UpdaeCommentById(ctx context.Context, id uint64, comment string) error {
	uc.log.WithContext(ctx).Infof("UpdaeCommentById: %v", id)
	err := uc.repo.UpdaeCommentById(ctx, id, comment)
	return err
}
func (uc *CommentContentUsecase) UpdaeEmailStateById(ctx context.Context, id uint64, state int32) error {
	uc.log.WithContext(ctx).Infof("UpdaeCommentById: %v", id)
	err := uc.repo.UpdaeEmailStateById(ctx, id, state)
	return err
}
func (uc *CommentContentUsecase) FindAllByParentID(ctx context.Context, id uint64, size int) ([]*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("UpdaeCommentById: %v", id)
	return uc.repo.FindAllByParentID(ctx, id, size)
}

func (uc *CommentContentUsecase) CountAll(ctx context.Context) *CommentCount {
	data := &CommentCount{}
	all, _ := uc.repo.CountByState(ctx, 0)
	success, _ := uc.repo.CountByState(ctx, 1)
	need, _ := uc.repo.CountByState(ctx, 2)
	data.Total = all
	data.NowTotal = success
	data.PendingTotal = need
	return data
}

func (uc *CommentContentUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	err := uc.repo.DeleteList(ctx, ids)
	return err
}

// FindByID 根据ID查询
func (uc *CommentContentUsecase) FindByID(ctx context.Context, id uint64) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return g, nil
}

func (uc *CommentContentUsecase) ListAll(ctx context.Context, opts CommentContentDOListOption) (*CommentContentDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	for i, _ := range dataDOs.Items {
		id := dataDOs.Items[i].ID
		aid := dataDOs.Items[i].AgentId
		uid := dataDOs.Items[i].MemberId
		if dataDOs.Items[i].RootId > 0 {
			parent, _ := uc.repo.FindParentByID(ctx, id)
			if parent != nil {
				dataDOs.Items[i].ParentUserName = parent.MemberName
			}
		}
		if !opts.IsWeb {
			aobj, _ := uc.au.FindOneByAgentID(ctx, aid)
			if aobj != nil {
				dataDOs.Items[i].ObjTitle = aobj.Title
				dataDOs.Items[i].ObjLink = aobj.PermaLink
			}
		}
		if uid > 0 {
			uobj, _ := uc.uu.FindOneByID(ctx, uid)
			if uobj != nil {
				dataDOs.Items[i].Avatar = uobj.Avatar
			} else {
				dataDOs.Items[i].Avatar = stringUtil.MD5ByStr(dataDOs.Items[i].Email)
			}
		}
		if opts.IsWeb {
			idx, _ := uc.ci.FindByID(ctx, id)
			if idx != nil {
				dataDOs.Items[i].Count = idx.Count
				dataDOs.Items[i].RootCount = idx.RootCount
			}
		}
	}

	return dataDOs, nil
}
