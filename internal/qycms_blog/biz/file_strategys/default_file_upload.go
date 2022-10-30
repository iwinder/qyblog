package file_strategys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/fileUtil"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"io/ioutil"
	"mime/multipart"
	"strings"
	"sync"
)

type IUploadStrategy interface {
	Upload(context.Context, multipart.File, *multipart.FileHeader, uint64) (*biz.FileLibDO, error)
	ListAll(context.Context, *v1.ListQyAdminFileRequest) (*biz.FileLibDOList, error)
}
type DefUpload struct {
	fic *biz.FileLibConfigUsecase
	fi  *biz.FileLibUsecase
}

var (
	defUpload *DefUpload
	once      sync.Once
)

func NewDefUpload(fic *biz.FileLibConfigUsecase, fi *biz.FileLibUsecase) *DefUpload {
	if fic == nil {
		log.Fatalf("NewDefUpload Failed : fic NotFound")
	}
	once.Do(func() {
		defUpload = &DefUpload{
			fic: fic,
			fi:  fi,
		}
	})
	if defUpload == nil {
		log.Fatalf("NewDefUpload Failed: init Failed")
	}

	return defUpload
}
func (u *DefUpload) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, id uint64) (*biz.FileLibDO, error) {
	config, err := u.fic.FindByTypeId(ctx, id)
	if err != nil {
		return nil, err
	}

	b, _ := ioutil.ReadAll(file)
	fmd5 := stringUtil.MD5(b)
	num := u.fi.CountByMd5(ctx, fmd5)
	if num > 0 {
		return nil, nil
	}
	//把文件保存到指定位置
	//config.Prefix+
	var path strings.Builder
	ext := fileUtil.GetExt(header.Filename)
	fname := stringUtil.GetShortUuid() + "." + ext
	parentPath, relativePath := fileUtil.GetLocalPrefixPath(config.Prefix, fname)
	err = fileUtil.CheckAndMkdirAll(parentPath)
	if err != nil {
		return nil, err
	}
	path.WriteString(parentPath)
	path.WriteString(header.Filename)
	err = ioutil.WriteFile(path.String(), b, 0777)
	if err != nil {
		return nil, err
	}

	fileDo := &biz.FileLibDO{
		OriginFileName: header.Filename,
		Fname:          fname,
		Fsize:          uint64(header.Size),
		Extention:      ext,
		MimeType:       header.Header.Get("Content-Type"),
		Fmd5:           fmd5,
		Fhash:          fmd5,
		RelativePath:   relativePath,
	}
	return fileDo, nil
}
func (u *DefUpload) ListAll(ctx context.Context, in *v1.ListQyAdminFileRequest) (*biz.FileLibDOList, error) {
	opts := biz.FileLibDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = 0
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.OriginFileName = in.SearchText
	opts.ListOptions.Init()
	return u.fi.ListAll(ctx, opts)
}
