package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

type CommentIndexDO struct {
	metaV1.ObjectMeta
	AgentId   uint64
	ObjId     uint64
	ObjType   int32
	MemberId  uint64
	RootId    uint64
	ParentId  uint64
	Floor     int32
	Count     int32
	RootCount int32
	LikeCount int32
	HateCount int32
	Attrs     string
}

// CommentAgentRepo is a Greater repo.
type CommentIndexRepo interface {
	Save(context.Context, *CommentIndexDO) (*CommentIndexDO, error)
	Update(context.Context, *CommentIndexDO) (*CommentIndexDO, error)
	UpdaeStateByIDs(context.Context, []uint64, int) error
	UpdateAddCountById(context.Context, uint64, bool) error
	UpdateMinusCountById(context.Context, uint64, bool) error
	FindByID(context.Context, uint64) (*CommentIndexDO, error)
	FindAllByParentID(context.Context, uint64) ([]*CommentIndexDO, error)
	DeleteList(c context.Context, uids []uint64) error
	UpdateObjIdByAgentIds(context.Context) error
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
	data, err := uc.repo.Save(ctx, g)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *CommentIndexUsecase) CreateCommentIndexByContent(ctx context.Context, g *CommentContentDO) (*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("CreateCommentIndex: %v-%v", g.AgentId, g.ID)
	gd := &CommentIndexDO{
		ObjectMeta: metaV1.ObjectMeta{
			ID:         g.ID,
			StatusFlag: g.StatusFlag,
		},
		AgentId:  g.AgentId,
		MemberId: g.MemberId,
		ParentId: g.RootId,
	}
	data, err := uc.CreateCommentIndex(ctx, gd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Update 更新
func (uc *CommentIndexUsecase) Update(ctx context.Context, g *CommentIndexDO) (*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("Update:  %v-%v", g.ObjId, g.ObjType)
	data, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return data, nil
}
func (uc *CommentIndexUsecase) UpdateStateByIDs(ctx context.Context, ids []uint64, state int) error {
	uc.log.WithContext(ctx).Infof("UpdaeStateByIDs: %v", ids)
	err := uc.repo.UpdaeStateByIDs(ctx, ids, state)
	return err
}
func (uc *CommentIndexUsecase) UpdateAddCountById(ctx context.Context, id uint64, isRoot bool) error {
	uc.log.WithContext(ctx).Infof("UpdateAddCountById:  %v-%v", id, isRoot)
	err := uc.repo.UpdateAddCountById(ctx, id, isRoot)
	return err
}
func (uc *CommentIndexUsecase) UpdateMinusCountById(ctx context.Context, id uint64, isRoot bool) error {
	uc.log.WithContext(ctx).Infof("UpdateAddCountById:  %v-%v", id, isRoot)
	err := uc.repo.UpdateMinusCountById(ctx, id, isRoot)
	return err
}

func (uc *CommentIndexUsecase) UpdateObjIdByAgentIds(ctx context.Context) error {
	uc.log.WithContext(ctx).Infof("UpdateObjIdByAgentIds: ")
	err := uc.repo.UpdateObjIdByAgentIds(ctx)
	return err
}

func (uc *CommentIndexUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
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

	return g, nil
}
func (uc *CommentIndexUsecase) FindAllByParentID(ctx context.Context, id uint64) ([]*CommentIndexDO, error) {
	uc.log.WithContext(ctx).Infof("FindByParentID: %v", id)
	g, err := uc.repo.FindAllByParentID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return g, nil
}
