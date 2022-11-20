package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
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
	Meta           string
	ParentUserName string
	ObjTitle       string
	ObjLink        string
	Avatar         string
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
}

// CommentAgentRepo is a Greater repo.
type CommentContentRepo interface {
	Save(context.Context, *CommentContentDO) (*CommentContentDO, error)
	Update(context.Context, *CommentContentDO) (*CommentContentDO, error)
	UpdaeStateByIDs(context.Context, []uint64, int) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*CommentContentDO, error)
	ListAll(context.Context, CommentContentDOListOption) (*CommentContentDOList, error)
	CountByState(context.Context, int) (int64, error)
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
	uc.ci.CreateCommentIndexByContent(ctx, g)
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
func (uc *CommentContentUsecase) UpdaeStateByIDs(ctx context.Context, ids []uint64, state int) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	err := uc.repo.UpdaeStateByIDs(ctx, ids, state)
	return err
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
		aid := dataDOs.Items[i].AgentId
		pid := dataDOs.Items[i].RootId
		uid := dataDOs.Items[i].MemberId
		if pid > 0 {
			parent, _ := uc.repo.FindByID(ctx, pid)
			if parent != nil {
				dataDOs.Items[i].ParentUserName = parent.MemberName
			}
			aobj, _ := uc.au.FindOneByAgentID(ctx, aid)
			if aobj != nil {
				dataDOs.Items[i].ObjTitle = aobj.Title
				dataDOs.Items[i].ObjLink = aobj.PermaLink
			}
			if uid > 0 {
				uobj, _ := uc.uu.FindOneByID(ctx, uid)
				if uobj != nil {
					dataDOs.Items[i].Avatar = uobj.Avatar
				}
			}
		}

	}

	return dataDOs, nil
}
