package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/gormutil"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
)

type tagsRepo struct {
	data *Data
	log  *log.Helper
}

// NewTagsRepo .
func NewTagsRepo(data *Data, logger log.Logger) biz.TagsRepo {
	return &tagsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Save 创建
func (r *tagsRepo) Save(ctx context.Context, obj *biz.TagsDO) (*biz.TagsDO, error) {
	objPO := &po.TagsPO{
		ObjectMeta:  obj.ObjectMeta,
		Name:        obj.Name,
		Identifier:  obj.Identifier,
		Description: obj.Description,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.TagsDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Update 更新
func (r *tagsRepo) Update(ctx context.Context, obj *biz.TagsDO) (*biz.TagsDO, error) {
	objPO := &po.TagsPO{
		Name:        obj.Name,
		Identifier:  obj.Identifier,
		Description: obj.Description,
	}
	//tObj := &po.TagsPO{}
	//tObj.ID = obj.ID
	err := r.data.Db.Model(&objPO).Where("id=?", obj.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.TagsDO{Name: objPO.Name}
	objDO.ID = objPO.ID
	return objDO, nil
}

// Delete 根据ID删除
func (r *tagsRepo) Delete(c context.Context, id uint64) error {
	objPO := &po.TagsPO{}
	objPO.ID = id
	err := r.data.Db.Delete(&objPO).Error
	return err
}

// DeleteList 根据ID批量删除
func (r *tagsRepo) DeleteList(c context.Context, ids []uint64) error {
	objPO := &po.TagsPO{}
	err := r.data.Db.Delete(&objPO, ids).Error
	return err
}

// FindByID 根据ID查询
func (r *tagsRepo) FindByID(ctx context.Context, id uint64) (*biz.TagsDO, error) {
	obj := &po.TagsPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.TagsDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}
func (r *tagsRepo) FindOneByName(ctx context.Context, name string) (*biz.TagsDO, error) {
	obj := &po.TagsPO{}
	err := r.data.Db.Where("name = ?", name).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.TagsDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}
func (r *tagsRepo) FindOneByIdentifier(ctx context.Context, name string) (*biz.TagsDO, error) {
	obj := &po.TagsPO{}
	err := r.data.Db.Where("identifier = ?", name).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.TagsDO{
		ObjectMeta: obj.ObjectMeta,
		Name:       obj.Name,
		Identifier: obj.Identifier,
	}
	return objDO, nil
}
func (r *tagsRepo) CountByIdentifier(ctx context.Context, str string) (int64, error) {
	var obj int64
	err := r.data.Db.Model(&po.TagsPO{}).Where("identifier like ?", str+"%").Count(&obj).Error
	if err != nil {
		return 0, err
	}
	return obj, nil
}

// ListAll 批量查询
func (r *tagsRepo) ListAll(c context.Context, opts biz.TagsDOListOption) (*biz.TagsDOList, error) {
	ret := &po.TagsPOList{}

	where := &po.TagsPO{}
	var err error
	queryDB := r.data.Db.Model(where)
	if len(opts.Name) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("name", "%"+opts.Name+"%"))
	}
	if len(opts.Identifier) > 0 {
		queryDB.Scopes(withFilterKeyLikeValue("identifier", "%"+opts.Identifier+"%"))
	}
	if opts.PageFlag {
		ol := gormutil.Unpointer(opts.Offset, opts.Limit)
		d := queryDB.Where(where).
			Offset(ol.Offset).
			Limit(ol.Limit).
			Order("id").
			Find(&ret.Items).
			Offset(-1).
			Limit(-1).
			Count(&ret.TotalCount)
		err = d.Error
	} else {
		d := queryDB.Where(where).
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

	infos := make([]*biz.TagsDO, 0, len(ret.Items))
	for _, obj := range ret.Items {
		infos = append(infos, &biz.TagsDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID:         obj.ID,
				InstanceID: obj.InstanceID,
				Extend:     obj.Extend,
				CreatedAt:  obj.CreatedAt,
				UpdatedAt:  obj.UpdatedAt,
			},
			Name:       obj.Name,
			Identifier: obj.Identifier,
		})
	}
	return &biz.TagsDOList{ListMeta: ret.ListMeta, Items: infos}, err
}

func (r *tagsRepo) FindAllByArticleID(ctx context.Context, articleId uint64) ([]*biz.TagsDO, error) {
	tagsPO := make([]*po.TagsPO, 0, 0)
	db := r.data.Db
	e := db.Where("ID in (?)", db.Table("qy_blog_article_tags").Select("tag_id ").Where(" article_id = ?", articleId)).Find(&tagsPO)
	if e.Error != nil {
		return nil, e.Error
	}
	infos := make([]*biz.TagsDO, 0, len(tagsPO))
	for _, obj := range tagsPO {
		infos = append(infos, &biz.TagsDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: obj.ID,
			},
			Name:       obj.Name,
			Identifier: obj.Identifier,
		})
	}
	return infos, nil
}
