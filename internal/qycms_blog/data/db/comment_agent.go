package db

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	biz "github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type CommentAgentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentAgentRepo .
func NewCommentAgentRepo(data *Data, logger log.Logger) biz.CommentAgentRepo {
	return &CommentAgentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommentAgentRepo) Save(ctx context.Context, g *biz.CommentAgentDO) (*biz.CommentAgentDO, error) {
	newData := &po.CommentAgentPO{
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId:   g.MemberId,
		Count:      g.Count,
		RootCount:  g.RootCount,
		AllCount:   g.AllCount,
		Attrs:      g.Attrs,
	}
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentAgentDO{ObjId: newData.ObjId, ObjType: g.ObjType}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentAgentRepo) Update(ctx context.Context, g *biz.CommentAgentDO) (*biz.CommentAgentDO, error) {
	newData := &po.CommentAgentPO{
		ObjId:     g.ObjId,
		ObjType:   g.ObjType,
		MemberId:  g.MemberId,
		Count:     g.Count,
		RootCount: g.RootCount,
		AllCount:  g.AllCount,
		Attrs:     g.Attrs,
	}
	tData := &po.CommentAgentPO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentAgentDO{ObjId: newData.ObjId, ObjType: g.ObjType}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentAgentRepo) UpdateAddCountById(ctx context.Context, id uint64, isRoot bool) error {
	//r.data.Db.Model(&po.CommentAgentPO{}).Where("id = ?", id).Update("name = 1")
	str := "UPDATE qy_blog_comment_agent set count = count+1 "
	if isRoot {
		str += " , root_count = root_count +1 "
	}
	str += " where id = " + fmt.Sprintf("%d", id)
	return r.data.Db.Exec(str).Error
}
func (r *CommentAgentRepo) UpdateMinusCountById(ctx context.Context, id uint64, isRoot bool) error {
	//r.data.Db.Model(&po.CommentAgentPO{}).Where("id = ?", id).Update("name = 1")
	str := "UPDATE qy_blog_comment_agent set count = count-1 "
	if isRoot {
		str += " , root_count = root_count - 1 "
	}
	str += " where id = " + fmt.Sprintf("%d", id)
	return r.data.Db.Exec(str).Error
}

func (r *CommentAgentRepo) FindByID(cxt context.Context, id uint64) (*biz.CommentAgentDO, error) {
	g := &po.CommentAgentPO{}
	err := r.data.Db.Where("id = ?", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentAgentDO{
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
