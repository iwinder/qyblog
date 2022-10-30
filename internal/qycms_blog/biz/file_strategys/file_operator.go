package file_strategys

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"mime/multipart"
)

type UploadOperator struct {
	strategy IUploadStrategy
}

func NewUploadOperator(id uint64, fic *biz.FileLibConfigUsecase, fi *biz.FileLibUsecase) *UploadOperator {
	var strategy IUploadStrategy
	switch id {
	case 1:
		strategy = NewDefUpload(fic, fi)
	case 2:
		strategy = NewQiNiuUpload(fic)
	}
	return &UploadOperator{
		strategy: strategy,
	}
}

// 设置策略
func (operator *UploadOperator) setStrategy(strategy IUploadStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *UploadOperator) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, id uint64) (*biz.FileLibDO, error) {
	return operator.strategy.Upload(ctx, file, header, id)
}
func (operator *UploadOperator) ListAll(ctx context.Context, in *v1.ListQyAdminFileRequest) (*biz.FileLibDOList, error) {
	return operator.strategy.ListAll(ctx, in)
}
