package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/data/po"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type CommentIndexDO struct {
	metaV1.ObjectMeta
	ObjId     int64
	ObjType   int32
	MemberId  int64
	RootId    int64
	OarentId  int64
	Floor     int32
	Count     int32
	RootCount int32
	LikeCount int32
	HateCount int32
	Attrs     int32
}

// CommentAgentRepo is a Greater repo.
type CommentIndexRepo interface {
	Save(context.Context, *CommentIndexDO) (*po.CommentIndexPO, error)
	Update(context.Context, *CommentIndexDO) (*po.CommentIndexPO, error)
	FindByID(context.Context, uint64) (*po.CommentIndexPO, error)
}

// CommentIndexUsecase is a CommentAgentDO usecase.
type CommentIndexUsecase struct {
	repo CommentIndexRepo
	log  *log.Helper
}

// NewCommentIndexUsecase new a CommentIndexDO usecase.
func NewCommentIndexUsecase(repo CommentIndexRepo, logger log.Logger) *CommentIndexUsecase {
	return &CommentIndexUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCommentIndex creates a CommentIndexDO, and returns the new CommentIndexDO.
func (uc *CommentIndexUsecase) CreateCommentIndex(ctx context.Context, g *CommentIndexDO) (*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("CreateCommentIndex: %v-%v", g.ObjId, g.ObjType)
	dataPO, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	data := &CommentIndexDO{ObjId: dataPO.ObjId, ObjType: g.ObjType}
	data.ID = dataPO.ID
	return data, nil
}

// Update 更新
func (uc *CommentIndexUsecase) Update(ctx context.Context, g *CommentIndexDO) (*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.ObjId, g.ObjType)
	dataPO, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &CommentIndexDO{ObjId: dataPO.ObjId, ObjType: g.ObjType}
	data.ID = dataPO.ID
	return data, nil
}

// FindByID 根据ID查询
func (uc *CommentIndexUsecase) FindByID(ctx context.Context, id uint64) (*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	data := &CommentIndexDO{
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId:   g.MemberId,
		RootId:     g.RootId,
		OarentId:   g.OarentId,
		Floor:      g.Floor,
		Count:      g.Count,
		RootCount:  g.RootCount,
		LikeCount:  g.LikeCount,
		HateCount:  g.HateCount,
		Attrs:      g.Attrs,
	}
	return data, nil
}
