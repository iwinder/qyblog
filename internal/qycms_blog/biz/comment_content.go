package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"
)

type CommentContentDO struct {
	metaV1.ObjectMeta
	MemberId    int64
	AtMemberIds string
	Agent       string
	MemberName  string
	Ip          int32
	Email       int32
	Url         int32
	RootId      int64
	Content     string
	Meta        string
}

// CommentAgentRepo is a Greater repo.
type CommentContentRepo interface {
	Save(context.Context, *CommentContentDO) (*po.CommentContentPO, error)
	Update(context.Context, *CommentContentDO) (*po.CommentContentPO, error)
	FindByID(context.Context, uint64) (*po.CommentContentPO, error)
}

// CommentContentUsecase is a CommentAgentDO usecase.
type CommentContentUsecase struct {
	repo CommentContentRepo
	log  *log.Helper
}

// NewCommentContentUsecase new a CommentContentDO usecase.
func NewCommentContentUsecase(repo CommentContentRepo, logger log.Logger) *CommentContentUsecase {
	return &CommentContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCommentContent creates a CommentContentDO, and returns the new CommentContentDO.
func (uc *CommentContentUsecase) CreateCommentContent(ctx context.Context, g *CommentContentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateCommentContent: %v-%v", g.MemberId, g.MemberName)
	dataPO, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	data := &CommentContentDO{MemberId: g.MemberId, MemberName: g.MemberName}
	data.ID = dataPO.ID
	return data, nil
}

// Update 更新
func (uc *CommentContentUsecase) Update(ctx context.Context, g *CommentContentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.MemberId, g.MemberName)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &CommentContentDO{MemberId: g.MemberId, MemberName: g.MemberName}
	data.ID = dataPO.ID
	return data, nil
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
	data := &CommentContentDO{
		ObjectMeta:  g.ObjectMeta,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      g.RootId,
		Content:     g.Content,
		Meta:        g.Meta,
	}
	return data, nil
}
