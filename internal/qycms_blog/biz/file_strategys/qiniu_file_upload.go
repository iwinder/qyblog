package file_strategys

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/fileUtil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"mime/multipart"
	"sync"
	"time"
)

type QiNiuUpload struct {
	fic *biz.FileLibConfigUsecase
}

var (
	qiNiuUpload *QiNiuUpload
	qiNiuCnce   sync.Once
)

func NewQiNiuUpload(fic *biz.FileLibConfigUsecase) *QiNiuUpload {
	if fic == nil {
		log.Fatalf("NewQiNiuUpload Failed : fic NotFound")
	}
	qiNiuCnce.Do(func() {
		qiNiuUpload = &QiNiuUpload{
			fic: fic,
		}
	})
	if qiNiuUpload == nil {
		log.Fatalf("NewQiNiuUpload Failed: init Failed")
	}
	return qiNiuUpload
}
func (u *QiNiuUpload) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, id uint64) (*biz.FileLibDO, error) {
	config, err := u.fic.FindByTypeId(ctx, id)
	if err != nil {
		return nil, err
	}
	upToken, _ := u.fic.GetTokenByTypeId(ctx, id)
	if len(upToken) == 0 {
		accessKey := config.AccessKey
		secretKey := config.SecretKey
		bucket := config.Bucket
		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}
		mac := qbox.NewMac(accessKey, secretKey)
		upToken = putPolicy.UploadToken(mac)
		u.fic.SetTokenByTypeId(ctx, id, upToken)
	}
	data, _ := ioutil.ReadAll(file)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": header.Filename,
		},
	}
	dataLen := int64(len(data))
	key := fileUtil.GetOssKey(config.Prefix, header.Filename)
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	durl := fileUtil.GetOssKey(config.Domain, ret.Key)
	return &biz.FileLibDO{
		ObjectMeta:     metaV1.ObjectMeta{ID: 0},
		OriginFileName: header.Filename,
		Fname:          ret.Key,
		Fhash:          ret.Hash,
		DefUrl:         durl,
	}, nil
}
func (u *QiNiuUpload) ListAll(ctx context.Context, in *v1.ListQyAdminFileRequest) (*biz.FileLibDOList, error) {
	if in.LastFlag || in.Marker == "-1" {
		return &biz.FileLibDOList{
			ListMeta: metaV1.ListMeta{
				TotalCount: 0,
				PageSize:   0,
				Current:    0,
				Pages:      0,
				FirstFlag:  false,
				LastFlag:   false,
				Marker:     "-1",
			},
			Items: make([]*biz.FileLibDO, 0, 0),
		}, nil
	}
	config, err := u.fic.FindByTypeId(ctx, in.TypeId)
	if err != nil {
		return nil, err
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	delimiter := ""
	//初始列举marker为空
	marker := in.Marker
	prefix := in.SearchText
	if len(config.Prefix) > 0 {
		prefix = fileUtil.GetOssKey(config.Prefix, in.SearchText)
	}
	entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(config.Bucket, prefix, delimiter, marker, int(in.PageSize))
	lastFlag := !hasNext
	dto := &biz.FileLibDOList{
		ListMeta: metaV1.ListMeta{
			TotalCount: in.Total,
			PageSize:   in.PageSize,
			Current:    in.Current,
			Pages:      in.Pages,
			FirstFlag:  false,
			LastFlag:   lastFlag,
			Marker:     nextMarker,
		},
	}
	item := make([]*biz.FileLibDO, 0, len(entries))
	for _, entry := range entries {
		durl := fileUtil.GetOssKey(config.Domain, entry.Key)
		startDate := time.UnixMilli(entry.PutTime * 100 / 1e6)
		item = append(item, &biz.FileLibDO{
			ObjectMeta:     metaV1.ObjectMeta{ID: 0, UpdatedAt: &startDate},
			OriginFileName: entry.Key,
			Fname:          entry.Key,
			Fsize:          uint64(entry.Fsize),
			Extention:      "",
			MimeType:       entry.MimeType,
			Fmd5:           "",
			Fhash:          entry.Hash,
			RelativePath:   "",
			Ftype:          entry.Type,
			EndUser:        entry.EndUser,
			Domain:         "",
			DefUrl:         durl,
		})
	}
	dto.Items = item
	return dto, nil
}
