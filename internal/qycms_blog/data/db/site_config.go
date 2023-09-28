package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/qycms_blog/biz"
	"github.com/iwinder/qyblog/internal/qycms_blog/data/po"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
)

var fileSiteConfigCacheKey = func() string {
	return "sys_site_config"
}
var siteConfigCache sync.Map

type siteConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewSiteConfigRepo .
func NewSiteConfigRepo(data *Data, logger log.Logger) biz.SiteConfigRepo {
	return &siteConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *siteConfigRepo) Save(ctx context.Context, data *biz.SiteConfigDO) (*biz.SiteConfigDO, error) {
	objPO := &po.SiteConfigPO{
		ObjectMeta:  data.ObjectMeta,
		ConfigKey:   data.ConfigKey,
		ConfigValue: data.ConfigValue,
		ConfigName:  data.ConfigName,
		ConfigTip:   data.ConfigTip,
		Ftype:       data.Ftype,
	}
	err := r.data.Db.Omit(clause.Associations).Create(objPO).Error
	if err != nil {
		return nil, err
	}
	objDO := &biz.SiteConfigDO{ConfigName: objPO.ConfigName}
	objDO.ID = objPO.ID
	return objDO, nil
}

func (r *siteConfigRepo) Update(ctx context.Context, data *biz.SiteConfigDO) (*biz.SiteConfigDO, error) {
	objPO := &po.SiteConfigPO{
		ObjectMeta:  data.ObjectMeta,
		ConfigKey:   data.ConfigKey,
		ConfigValue: data.ConfigValue,
		ConfigName:  data.ConfigName,
		ConfigTip:   data.ConfigTip,
		Ftype:       data.Ftype,
	}
	err := r.data.Db.Model(&po.SiteConfigPO{}).Where("id=?", data.ID).Updates(&objPO).Error
	if err != nil {
		return nil, err
	}

	dataDO := &biz.SiteConfigDO{ConfigName: objPO.ConfigName}
	dataDO.ID = objPO.ID
	return dataDO, nil
}
func (r *siteConfigRepo) UpdateInBatches(ctx context.Context, datas []*biz.SiteConfigDO) error {
	//objPOs := make([]*po.SiteConfigPO, 0, len(datas))
	var err error
	for _, data := range datas {
		objPO := &po.SiteConfigPO{
			ObjectMeta:  data.ObjectMeta,
			ConfigKey:   data.ConfigKey,
			ConfigValue: data.ConfigValue,
			ConfigName:  data.ConfigName,
			ConfigTip:   data.ConfigTip,
			Ftype:       data.Ftype,
		}
		d := r.data.Db.Model(&objPO).Updates(&objPO).Error
		if d != nil {
			log.Error("批量更新失败", d)
			err = d
		}
	}

	return err
}

func (r *siteConfigRepo) GetValueByKey(key string) (any, bool) {
	return siteConfigCache.Load(key)
}

//func (r *siteConfigRepo) FindByMd5(ctx context.Context, fmd5 string) (*biz.SiteConfigDO, error) {
//	obj := &po.SiteConfigPO{}
//	err := r.data.Db.Where("fmd5 = ?", fmd5).First(&obj).Error
//	if err != nil {
//		return nil, err
//	}
//	objDo := &biz.SiteConfigDO{
//		ObjectMeta:  obj.ObjectMeta,
//		ConfigKey:   obj.ConfigKey,
//		ConfigValue: obj.ConfigValue,
//		ConfigName:  obj.ConfigName,
//		ConfigTip:   obj.ConfigTip,
//		Ftype:       obj.Ftype,
//	}
//	return objDo, nil
//}

//func (r *siteConfigRepo) CountByMd5(ctx context.Context, fmd5 string) (int64, error) {
//	var obj int64
//	err := r.data.Db.Model(&po.SiteConfigPO{}).Where("fmd5 = ?", fmd5).Count(&obj).Error
//	if err != nil {
//		return 0, err
//	}
//	return obj, nil
//}
//func (r *siteConfigRepo) DeleteList(c context.Context, ids []uint64) error {
//	objPO := &po.SiteConfigPO{}
//	err := r.data.Db.Delete(&objPO, ids).Error
//	return err
//}

func (r *siteConfigRepo) InitLocalSiteConfigCache(ctx context.Context) {
	alen := 0
	siteConfigCache.Range(func(k, v interface{}) bool {
		alen++
		return true
	})
	if alen == 0 {
		ops := biz.SiteConfigDOListOption{}
		r.ListAll(ctx, ops)
	}
}

func (r *siteConfigRepo) ListAll(ctx context.Context, opts biz.SiteConfigDOListOption) ([]*biz.SiteConfigDO, error) {

	ret := &po.SiteConfigPOList{}
	var err error
	ret.Items, err = r.getSiteConfigFromCache(ctx)
	if err != nil || ret.Items == nil || len(ret.Items) == 0 {
		d := r.data.Db.Model(&po.SiteConfigPO{}).Order("id ").Find(&ret.Items).Count(&ret.TotalCount)
		err = d.Error
		if err != nil {
			return nil, err
		}
		r.setSiteConfigCache(ctx, ret.Items)
	}

	var infos []*biz.SiteConfigDO
	if len(opts.Types) > 0 {
		str := opts.Types + ","
		infos = make([]*biz.SiteConfigDO, 0, 5)
		for _, obj := range ret.Items {
			key := fmt.Sprintf("%d,", obj.Ftype)
			if strings.Contains(str, key) {
				infos = append(infos, &biz.SiteConfigDO{
					ObjectMeta:  obj.ObjectMeta,
					ConfigKey:   obj.ConfigKey,
					ConfigValue: obj.ConfigValue,
					ConfigName:  obj.ConfigName,
					ConfigTip:   obj.ConfigTip,
					Ftype:       obj.Ftype,
				})
			}
		}
	} else {
		alen := 0
		siteConfigCache.Range(func(k, v interface{}) bool {
			alen++
			return true
		})
		needInit := alen == 0
		infos = make([]*biz.SiteConfigDO, 0, len(ret.Items))
		for _, obj := range ret.Items {
			infos = append(infos, &biz.SiteConfigDO{
				ObjectMeta:  obj.ObjectMeta,
				ConfigKey:   obj.ConfigKey,
				ConfigValue: obj.ConfigValue,
				ConfigName:  obj.ConfigName,
				ConfigTip:   obj.ConfigTip,
				Ftype:       obj.Ftype,
			})

		}
		if needInit {
			for _, obj := range ret.Items {
				siteConfigCache.Store(obj.ConfigKey, obj.ConfigValue)
			}
		}
	}

	return infos, nil
}

func (r *siteConfigRepo) getSiteConfigFromCache(ctx context.Context) ([]*po.SiteConfigPO, error) {
	key := fileSiteConfigCacheKey()
	result, err := r.data.RedisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = make([]*po.SiteConfigPO, 0, 0)
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *siteConfigRepo) setSiteConfigCache(ctx context.Context, data []*po.SiteConfigPO) {
	key := fileSiteConfigCacheKey()
	marshal, err := json.Marshal(data)
	if err != nil {
		r.log.Errorf("fail to set data cache:json.Marshal(%v) error(%v)", data, err)
	}
	err = r.data.RedisCli.Set(ctx, key, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set data cache:redis.Set(%v) error(%v)", data, err)
	}
	for _, obj := range data {
		siteConfigCache.Store(obj.ConfigKey, obj.ConfigValue)
	}
}
