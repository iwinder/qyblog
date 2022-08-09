package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/data/po"
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
	err := r.data.db.Create(newData).Error
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
	err := r.data.db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentContentRepo) FindByID(cxt context.Context, id uint64) (*po.CommentContentPO, error) {
	tData := &po.CommentContentPO{}
	err := r.data.db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}
