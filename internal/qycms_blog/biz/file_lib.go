package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
)

var (
	// ErrFileLibNotFound is file lib not found.
	ErrFileLibNotFound = errors.NotFound("109404", "file lib not found")
)

type FileLibDO struct {
	metaV1.ObjectMeta
	OriginFileName string // 原始名称
	Fname          string // 文件名
	Fsize          uint64 // 文件大小
	Extention      string // 文件扩展名
	MimeType       string // 资源的 MIME 类型
	Fhash          string // 文件的HASH值
	Fmd5           string // 文件md5值
	RelativePath   string // 通过浏览器访问的相对路径
	Ftype          int    // 文件的存储类型，0为普通存储，1为低频存储
	EndUser        string // 文件上传时设置的endUser 七牛
	Domain         string // 域名
	DefUrl         string // 文件链接
}
type FileLibDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*FileLibDO `json:"items"`
}

type FileLibDOListOption struct {
	metaV1.ListOptions `json:"page"`
	FileLibDO          `json:"item"`
}

type FileLibRepo interface {
	Save(ctx context.Context, data *FileLibDO) (*FileLibDO, error)
	Update(context.Context, *FileLibDO) (*FileLibDO, error)
	FindByMd5(context.Context, string) (*FileLibDO, error)
	CountByMd5(ctx context.Context, fmd5 string) (int64, error)
	DeleteList(c context.Context, uids []uint64) error
	ListAll(c context.Context, opts FileLibDOListOption) (*FileLibDOList, error)
}

type FileLibUsecase struct {
	repo FileLibRepo
	log  *log.Helper
}

// NewFileLibUsecase new a UserDO usecase.
func NewFileLibUsecase(repo FileLibRepo, logger log.Logger) *FileLibUsecase {
	return &FileLibUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *FileLibUsecase) Save(ctx context.Context, data *FileLibDO) (*FileLibDO, error) {
	uc.log.WithContext(ctx).Infof("Save: %v", data.Fname)
	dataDO, err := uc.repo.Save(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *FileLibUsecase) Update(ctx context.Context, data *FileLibDO) (*FileLibDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", data.Fname)
	dataDO, err := uc.repo.Update(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibNotFound
		}
		return nil, err
	}

	return dataDO, nil
}
func (uc *FileLibUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}
func (uc *FileLibUsecase) FindByMd5(ctx context.Context, fmd5 string) (*FileLibDO, error) {
	uc.log.WithContext(ctx).Infof("FindByMd5: %v", fmd5)
	return uc.repo.FindByMd5(ctx, fmd5)
}
func (uc *FileLibUsecase) CountByMd5(ctx context.Context, fmd5 string) int64 {
	uc.log.WithContext(ctx).Infof("CountByMd5: %v", fmd5)
	count, _ := uc.repo.CountByMd5(ctx, fmd5)
	return count
}

func (uc *FileLibUsecase) ListAll(ctx context.Context, opts FileLibDOListOption) (*FileLibDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFileLibNotFound
		}
		return nil, err
	}

	return dataDOs, nil
}
