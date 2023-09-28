package db

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	biz "github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
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

func (r *CommentContentRepo) Save(ctx context.Context, g *biz.CommentContentDO) (*biz.CommentContentDO, error) {
	newData := &po.CommentContentPO{
		ObjectMeta:  g.ObjectMeta,
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId: sql.NullInt64{
			Int64: int64(g.RootId),
			Valid: true,
		},
		Content: g.Content,
		Attrs:   g.Attrs,
		EmailState: sql.NullInt32{
			Int32: g.EmailState,
			Valid: true,
		},
	}
	err := r.data.Db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentContentDO{MemberId: g.MemberId, MemberName: g.MemberName}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentContentRepo) Update(ctx context.Context, g *biz.CommentContentDO) (*biz.CommentContentDO, error) {
	newData := &po.CommentContentPO{
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		//RootId: sql.NullInt64{
		//	Int64: int64(g.RootId),
		//	Valid: true,
		//},
		Content: g.Content,
		Attrs:   g.Attrs,
	}
	tData := &po.CommentContentPO{}
	tData.ID = g.ID
	err := r.data.Db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentContentDO{MemberId: g.MemberId, MemberName: g.MemberName}
	data.ID = newData.ID
	return data, nil
}

func (r *CommentContentRepo) UpdaeStateByIDs(cxt context.Context, ids []uint64, state int) error {
	return r.data.Db.Model(&po.CommentContentPO{}).Where("id IN ?", ids).Update("status_flag", state).Error
}
func (r *CommentContentRepo) UpdaeCommentById(cxt context.Context, id uint64, comment string) error {
	return r.data.Db.Model(&po.CommentContentPO{}).Where("id = ?", id).Update("content", comment).Error
}
func (r *CommentContentRepo) UpdaeEmailStateById(cxt context.Context, id uint64, satae int32) error {
	return r.data.Db.Model(&po.CommentContentPO{}).Where("id = ?", id).Update("email_state", satae).Error
}
func (r *CommentContentRepo) DeleteList(ctx context.Context, ids []uint64) error {
	userPO := &po.CommentContentPO{}
	if ids == nil || len(ids) == 0 {
		return nil
	}
	err := r.data.Db.Delete(&userPO, ids).Error
	return err
}
func (r *CommentContentRepo) CountByState(ctx context.Context, state int) (int64, error) {
	var obj int64
	var err error
	if state > 0 {
		err = r.data.Db.Model(&po.CommentContentPO{}).Where("status_flag = ?", state).Count(&obj).Error
	} else {
		err = r.data.Db.Model(&po.CommentContentPO{}).Count(&obj).Error
	}
	if err != nil {
		return 0, err
	}
	return obj, nil
}
func (r *CommentContentRepo) FindByID(cxt context.Context, id uint64) (*biz.CommentContentDO, error) {
	g := &po.CommentContentPO{}
	err := r.data.Db.Where("id = ?", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentContentDO{
		ObjectMeta:  g.ObjectMeta,
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      uint64(g.RootId.Int64),
		Content:     g.Content,
		Attrs:       g.Attrs,
	}
	return data, nil
}
func (r *CommentContentRepo) FindParentByID(cxt context.Context, id uint64) (*biz.CommentContentDO, error) {
	db := r.data.Db
	g := &po.CommentContentPO{}
	err := db.Where("ID = (?)", db.Table("qy_blog_comment_index").Select("parent_id ").Where(" id = ?", id)).Find(&g).Error
	if err != nil {
		return nil, err
	}
	data := &biz.CommentContentDO{
		ObjectMeta:  g.ObjectMeta,
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      uint64(g.RootId.Int64),
		Content:     g.Content,
		Attrs:       g.Attrs,
	}
	return data, nil
}
func (r *CommentContentRepo) FindAllByParentID(cxt context.Context, id uint64, size int) ([]*biz.CommentContentDO, error) {
	db := r.data.Db
	g := make([]*po.CommentContentPO, 0, 0)
	var err error

	if size > 0 {
		err = db.Where("ID in (?)", db.Table("qy_blog_comment_index").Select("id ").Where(" parent_id = ? and status_flag = 1 ", id)).Find(&g).Limit(size).Error
	} else {
		err = db.Where("ID in (?)", db.Table("qy_blog_comment_index").Select("id ").Where(" parent_id = ? and status_flag = 1  ", id)).Find(&g).Error
	}
	if err != nil {
		return nil, err
	}
	data := make([]*biz.CommentContentDO, 0, 0)
	for _, dp := range g {
		data = append(data, doToCommentContentDO(dp))
	}
	return data, nil
}
func (r *CommentContentRepo) ListAll(ctx context.Context, opts biz.CommentContentDOListOption) (*biz.CommentContentDOList, error) {
	ret := &po.CommentContentPOList{}

	var err error
	query := r.data.Db.Model(&po.CommentContentPO{})
	if len(opts.Order) == 0 {
		opts.Order = " created_at desc,id desc "
	}
	if len(opts.Content) > 0 {
		query.Scopes(withFilterKeyLikeValue("content", "%"+opts.Content+"%"))
	}
	if opts.StatusFlag > 0 {
		query.Scopes(withFilterKeyEquarlsValue("status_flag", opts.StatusFlag))
	}
	if opts.AgentId > 0 {
		query.Scopes(withFilterKeyEquarlsValue("agent_id", opts.AgentId))
	}

	if opts.IsWeb {
		if opts.RootId >= 0 {
			query.Scopes(withFilterKeyEquarlsValue("root_id", opts.RootId))
		}
	} else {
		if opts.RootId > 0 {
			query.Scopes(withFilterKeyEquarlsValue("root_id", opts.RootId))
		}
	}
	if opts.EmailState > 0 {
		query.Scopes(withFilterKeyEquarlsValue("email_state", opts.EmailState))
	}

	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := query.
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order(opts.Order).
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := query.
			Find(&ret.Items).
			Count(&ret.TotalCount)
		err = d.Error
	}
	opts.TotalCount = ret.TotalCount
	opts.IsLast()
	ret.FirstFlag = opts.FirstFlag
	ret.Current = opts.Current
	ret.PageSize = opts.PageSize
	ret.LastFlag = opts.LastFlag
	ret.Pages = opts.Pages
	infos := make([]*biz.CommentContentDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.CommentContentDO{
			ObjectMeta:  obj.ObjectMeta,
			MemberId:    obj.MemberId,
			AtMemberIds: obj.AtMemberIds,
			Agent:       obj.Agent,
			MemberName:  obj.MemberName,
			Ip:          obj.Ip,
			Email:       obj.Email,
			Url:         obj.Url,
			RootId:      uint64(obj.RootId.Int64),
			Content:     obj.Content,
			Attrs:       obj.Attrs,
			AgentId:     obj.AgentId,
			EmailState:  int32(obj.EmailState.Int32),
		})
	}
	return &biz.CommentContentDOList{ListMeta: ret.ListMeta, Items: infos}, err
}
func doToCommentContentDO(g *po.CommentContentPO) *biz.CommentContentDO {
	data := &biz.CommentContentDO{
		ObjectMeta:  g.ObjectMeta,
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: g.AtMemberIds,
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      uint64(g.RootId.Int64),
		Content:     g.Content,
		Attrs:       g.Attrs,
	}
	return data
}
