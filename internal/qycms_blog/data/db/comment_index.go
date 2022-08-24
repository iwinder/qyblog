package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	biz "github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type CommentIndexRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentIndexRepo .
func NewCommentIndexRepo(data *Data, logger log.Logger) biz.CommentIndexRepo {
	return &CommentIndexRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommentIndexRepo) Save(ctx context.Context, g *biz.CommentIndexDO) (*po.CommentIndexPO, error) {
	newData := &po.CommentIndexPO{
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
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentIndexRepo) Update(ctx context.Context, g *biz.CommentIndexDO) (*po.CommentIndexPO, error) {
	newData := &po.CommentIndexPO{
		ObjId:     g.ObjId,
		ObjType:   g.ObjType,
		MemberId:  g.MemberId,
		RootId:    g.RootId,
		OarentId:  g.OarentId,
		Floor:     g.Floor,
		Count:     g.Count,
		RootCount: g.RootCount,
		LikeCount: g.LikeCount,
		HateCount: g.HateCount,
		Attrs:     g.Attrs,
	}
	tData := &po.CommentIndexPO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentIndexRepo) FindByID(cxt context.Context, id uint64) (*po.CommentIndexPO, error) {
	tData := &po.CommentIndexPO{}
	err := r.data.Db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}
