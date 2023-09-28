package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
	"time"
)

var fileLibTypeConfigCacheKey = func(typeId string) string {
	return "file_lib_config_" + typeId
}
var fileLibTypeTokenCacheKey = func(typeId string) string {
	return "file_lib_token_" + typeId
}

type fileLibTypeConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileLibConfigRepo .
func NewFileLibConfigRepo(data *Data, logger log.Logger) biz.FileLibConfigRepo {
	return &fileLibTypeConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *fileLibTypeConfigRepo) Save(ctx context.Context, data *biz.FileLibConfigDO) (*biz.FileLibConfigDO, error) {
	objPO := &po.FileLibConfigPO{
		ObjectMeta: data.ObjectMeta,
		AccessKey:  data.AccessKey,
		SecretKey:  data.SecretKey,
		Bucket:     data.Bucket,
		Prefix:     data.Prefix,
		Domain:     data.Domain,
		Endpoint:   data.Endpoint,
		TypeId:     data.TypeId,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	cacheKey := fileLibTypeConfigCacheKey(fmt.Sprintf("%d", objPO.TypeId))
	r.setFileLibConfigCache(ctx, objPO, cacheKey)
	objDO := &biz.FileLibConfigDO{TypeId: objPO.TypeId}
	objDO.ID = objPO.ID
	return objDO, nil
}

func (r *fileLibTypeConfigRepo) Update(ctx context.Context, data *biz.FileLibConfigDO) (*biz.FileLibConfigDO, error) {
	objPO := &po.FileLibConfigPO{
		ObjectMeta: data.ObjectMeta,
		AccessKey:  data.AccessKey,
		SecretKey:  data.SecretKey,
		Bucket:     data.Bucket,
		Prefix:     data.Prefix,
		Domain:     data.Domain,
		Endpoint:   data.Endpoint,
		TypeId:     data.TypeId,
	}
	err := r.data.Db.Model(&po.FileLibConfigPO{}).Where("id=?", data.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}
	cacheKey := fileLibTypeConfigCacheKey(fmt.Sprintf("%d", objPO.TypeId))
	r.setFileLibConfigCache(ctx, objPO, cacheKey)
	dataDO := &biz.FileLibConfigDO{TypeId: objPO.TypeId}
	dataDO.ID = objPO.ID
	return dataDO, nil
}
func (r *fileLibTypeConfigRepo) FindByID(ctx context.Context, id uint64) (*biz.FileLibConfigDO, error) {
	obj := &po.FileLibConfigPO{}
	err := r.data.Db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.FileLibConfigDO{
		ObjectMeta: obj.ObjectMeta,
		AccessKey:  obj.AccessKey,
		SecretKey:  obj.SecretKey,
		Bucket:     obj.Bucket,
		Prefix:     obj.Prefix,
		Domain:     obj.Domain,
		Endpoint:   obj.Endpoint,
		TypeId:     obj.TypeId,
	}
	return objDO, nil
}

func (r *fileLibTypeConfigRepo) FindByTypeId(ctx context.Context, id uint64) (*biz.FileLibConfigDO, error) {
	cacheKey := fileLibTypeConfigCacheKey(fmt.Sprintf("%d", id))
	obj, err := r.getFileLibConfigFromCache(ctx, cacheKey)
	if err != nil {
		obj = &po.FileLibConfigPO{}
		err = r.data.Db.Where("type_id = ?", id).First(&obj).Error
		if err != nil {
			return nil, err
		}
		r.setFileLibConfigCache(ctx, obj, cacheKey)
	}

	objDO := &biz.FileLibConfigDO{
		ObjectMeta: obj.ObjectMeta,
		AccessKey:  obj.AccessKey,
		SecretKey:  obj.SecretKey,
		Bucket:     obj.Bucket,
		Prefix:     obj.Prefix,
		Domain:     obj.Domain,
		Endpoint:   obj.Endpoint,
		TypeId:     obj.TypeId,
	}
	return objDO, nil
}

func (r *fileLibTypeConfigRepo) getFileLibConfigFromCache(ctx context.Context, key string) (*po.FileLibConfigPO, error) {
	result, err := r.data.RedisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &po.FileLibConfigPO{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *fileLibTypeConfigRepo) setFileLibConfigCache(ctx context.Context, data *po.FileLibConfigPO, key string) {
	marshal, err := json.Marshal(data)
	if err != nil {
		r.log.Errorf("fail to set data cache:json.Marshal(%v) error(%v)", data, err)
	}
	err = r.data.RedisCli.Set(ctx, key, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:redis.Set(%v) error(%v)", data, err)
	}
}

func (r *fileLibTypeConfigRepo) SetTokenByTypeId(ctx context.Context, id uint64, token string) (bool, error) {

	key := fileLibTypeTokenCacheKey(fmt.Sprintf("%d", id))
	err := r.data.RedisCli.Set(ctx, key, token, 3000*time.Second).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:(%v) error(%v)", token, err)
		return false, err
	}
	return true, nil
}
func (r *fileLibTypeConfigRepo) GetTokenByTypeId(ctx context.Context, id uint64) (string, error) {

	key := fileLibTypeTokenCacheKey(fmt.Sprintf("%d", id))
	result, err := r.data.RedisCli.Get(ctx, key).Result()
	if err != nil {
		r.log.Errorf("fail to get data cache:(%v) error(%v)", key, err)
		return "", err
	}
	return result, nil
}
