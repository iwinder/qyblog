package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	biz "github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type CommentContentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentIndexRepo .
func NewCommentContentRepo(data *Data, logger log.Logger) biz.CommentContentRepo {
	return &CommentContentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommentContentRepo) Save(ctx context.Context, g *biz.CommentContentDO) (*po.CommentContentPO, error) {
	newData := &po.CommentContentPO{
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
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentContentRepo) Update(ctx context.Context, g *biz.CommentContentDO) (*po.CommentContentPO, error) {
	newData := &po.CommentContentPO{
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
	tData := &po.CommentContentPO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentContentRepo) FindByID(cxt context.Context, id uint64) (*po.CommentContentPO, error) {
	tData := &po.CommentContentPO{}
	err := r.data.Db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}
