package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	biz "github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
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

func (r *CommentIndexRepo) Save(ctx context.Context, g *biz.CommentIndexDO) (*biz.CommentIndexDO, error) {
	newData := &po.CommentIndexPO{
		AgentId:    g.AgentId,
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId: sql.NullInt64{
			Int64: int64(g.MemberId),
			Valid: true,
		},
		RootId: sql.NullInt64{
			Int64: int64(g.RootId),
			Valid: true,
		},
		ParentId: sql.NullInt64{
			Int64: int64(g.ParentId),
			Valid: true,
		},
		Floor:     g.Floor,
		Count:     g.Count,
		RootCount: g.RootCount,
		LikeCount: g.LikeCount,
		HateCount: g.HateCount,
		Attrs:     g.Attrs,
	}
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentIndexDO{ObjId: newData.ObjId, ObjType: g.ObjType}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentIndexRepo) Update(ctx context.Context, g *biz.CommentIndexDO) (*biz.CommentIndexDO, error) {
	newData := &po.CommentIndexPO{
		AgentId: g.AgentId,
		ObjId:   g.ObjId,
		ObjType: g.ObjType,
		MemberId: sql.NullInt64{
			Int64: int64(g.MemberId),
			Valid: true,
		},
		RootId: sql.NullInt64{
			Int64: int64(g.RootId),
			Valid: true,
		},
		ParentId: sql.NullInt64{
			Int64: int64(g.ParentId),
			Valid: true,
		},
		Floor: g.Floor,
		//Count:     g.Count,
		//RootCount: g.RootCount,
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
	data := &biz.CommentIndexDO{ObjId: newData.ObjId, ObjType: g.ObjType}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentIndexRepo) UpdaeStateByIDs(cxt context.Context, ids []uint64, state int) error {
	return r.data.Db.Model(&po.CommentIndexPO{}).Where("id IN ?", ids).Update("status_flag", state).Error
}

func (r *CommentIndexRepo) UpdateAddCountById(ctx context.Context, id uint64, isRoot bool) error {
	//r.data.Db.Model(&po.CommentAgentPO{}).Where("id = ?", id).Update("name = 1")
	str := "UPDATE qy_blog_comment_index set count = count+1 "
	if isRoot {
		str += " , root_count = root_count +1 "
	}
	str += " where id = " + fmt.Sprintf("%d", id)
	return r.data.Db.Exec(str).Error
}
func (r *CommentIndexRepo) UpdateMinusCountById(ctx context.Context, id uint64, isRoot bool) error {
	//r.data.Db.Model(&po.CommentAgentPO{}).Where("id = ?", id).Update("name = 1")
	str := "UPDATE qy_blog_comment_index set count = count-1 "
	if isRoot {
		str += " , root_count = root_count - 1 "
	}
	str += " where id = " + fmt.Sprintf("%d", id)
	return r.data.Db.Exec(str).Error
}
func (r *CommentIndexRepo) UpdateObjIdByAgentIds(ctx context.Context) error {
	str := " UPDATE qy_blog_comment_index  ci INNER JOIN ( SELECT id, obj_id, obj_type FROM qy_blog_comment_agent ca WHERE ca.id  )  n on ci.agent_id  = n.id SET ci.obj_id = n.obj_id, ci.obj_type=n.obj_type where  ci.obj_id = 0;"
	return r.data.Db.Exec(str).Error
}
func (r *CommentIndexRepo) DeleteList(ctx context.Context, ids []uint64) error {
	userPO := &po.CommentIndexPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	return err
}
func (r *CommentIndexRepo) FindByID(cxt context.Context, id uint64) (*biz.CommentIndexDO, error) {
	g := &po.CommentIndexPO{}
	err := r.data.Db.Where("id = ?", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := doToCommentIndexDO(g)
	return data, nil
}
func (r *CommentIndexRepo) FindAllByParentID(cxt context.Context, id uint64) ([]*biz.CommentIndexDO, error) {
	g := make([]*po.CommentIndexPO, 0, 0)
	err := r.data.Db.Where("parent_id = ?", id).Find(&g).Error
	if err != nil {
		return nil, err
	}
	data := make([]*biz.CommentIndexDO, 0, 0)
	for _, dp := range g {
		data = append(data, doToCommentIndexDO(dp))
	}
	return data, nil
}

func doToCommentIndexDO(g *po.CommentIndexPO) *biz.CommentIndexDO {
	data := &biz.CommentIndexDO{
		AgentId:    g.AgentId,
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId:   uint64(g.MemberId.Int64),
		RootId:     uint64(g.RootId.Int64),
		ParentId:   uint64(g.ParentId.Int64),
		Floor:      g.Floor,
		Count:      g.Count,
		RootCount:  g.RootCount,
		LikeCount:  g.LikeCount,
		HateCount:  g.HateCount,
		Attrs:      g.Attrs,
	}
	return data
}
