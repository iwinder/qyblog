package biz

import (
	"context"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrCommentAgentNotFound is common agent not found.
	ErrCommentAgentNotFound = errors.NotFound("106404", "common agent not found")
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
	Attrs     string
}

// CommentAgentRepo is a Greater repo.
type CommentAgentRepo interface {
	Save(context.Context, *CommentAgentDO) (*CommentAgentDO, error)
	Update(context.Context, *CommentAgentDO) (*CommentAgentDO, error)
	FindByID(context.Context, uint64) (*CommentAgentDO, error)
	UpdateAddCountById(context.Context, uint64, bool) error
	UpdateMinusCountById(context.Context, uint64, bool) error
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
	data, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Update 更新
func (uc *CommentAgentUsecase) Update(ctx context.Context, g *CommentAgentDO) (*CommentAgentDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.ObjId, g.ObjType)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCommentAgentNotFound
		}
		return nil, err
	}
	data := &CommentAgentDO{ObjId: dataPO.ObjId, ObjType: g.ObjType}
	data.ID = dataPO.ID
	return data, nil
}
func (uc *CommentAgentUsecase) UpdateAddCountById(ctx context.Context, id uint64, isRoot bool) error {
	uc.log.WithContext(ctx).Infof("UpdateAddCountById:  %v-%v", id, isRoot)
	err := uc.repo.UpdateAddCountById(ctx, id, isRoot)
	return err
}
func (uc *CommentAgentUsecase) UpdateMinusCountById(ctx context.Context, id uint64, isRoot bool) error {
	uc.log.WithContext(ctx).Infof("UpdateAddCountById:  %v-%v", id, isRoot)
	err := uc.repo.UpdateAddCountById(ctx, id, isRoot)
	return err
}

// FindByID 根据ID查询
func (uc *CommentAgentUsecase) FindByID(ctx context.Context, id uint64) (*CommentAgentDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCommentAgentNotFound
		}
		return nil, err
	}

	return g, nil
}
