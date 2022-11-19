package biz

import (
	"context"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/iwinder/qingyucms/api/qycms_user/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// CommentAgentDO is a CommentAgentDO model.
type CommentAgentDO struct {
	metaV1.ObjectMeta
	ObjId     uint64
	ObjType   int32
	MemberId  uint64
	Count     int32
	RootCount int32
	AllCount  int32
	State     int8
	Attrs     int32
}

// CommentAgentRepo is a Greater repo.
type CommentAgentRepo interface {
	Save(context.Context, *CommentAgentDO) (*po.CommentAgentPO, error)
	Update(context.Context, *CommentAgentDO) (*po.CommentAgentPO, error)
	FindByID(context.Context, uint64) (*po.CommentAgentPO, error)
}

// CommentIndexUsecase is a CommentAgentDO usecase.
type CommentAgentUsecase struct {
	repo CommentAgentRepo
	log  *log.Helper
}

// NewCommentAgentUsecase new a CommentAgentDO usecase.
func NewCommentAgentUsecase(repo CommentAgentRepo, logger log.Logger) *CommentAgentUsecase {
	return &CommentAgentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCommentAgent creates a CommentAgentDO, and returns the new CommentAgentDO.
func (uc *CommentAgentUsecase) CreateCommentAgent(ctx context.Context, g *CommentAgentDO) (*CommentAgentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateCommentAgent: %v-%v", g.ObjId, g.ObjType)
	dataPO, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	data := &CommentAgentDO{ObjId: dataPO.ObjId, ObjType: g.ObjType}
	data.ID = dataPO.ID
	return data, nil
}

// Update 更新
func (uc *CommentAgentUsecase) Update(ctx context.Context, g *CommentAgentDO) (*CommentAgentDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.ObjId, g.ObjType)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &CommentAgentDO{ObjId: dataPO.ObjId, ObjType: g.ObjType}
	data.ID = dataPO.ID
	return data, nil
}

// FindByID 根据ID查询
func (uc *CommentAgentUsecase) FindByID(ctx context.Context, id uint64) (*CommentAgentDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &CommentAgentDO{
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId:   g.MemberId,
		Count:      g.Count,
		RootCount:  g.RootCount,
		AllCount:   g.AllCount,
		Attrs:      g.Attrs,
	}
	return data, nil
}
