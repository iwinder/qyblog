package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"gorm.io/gorm"
)

type TagsDO struct {
	metaV1.ObjectMeta
	Name        string
	Identifier  string
	Description string
}

type TagsDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*TagsDO `json:"items"`
}

type TagsDOListOption struct {
	metaV1.ListOptions `json:"page"`
	TagsDO             `json:"item"`
}

type TagsRepo interface {
	Save(context.Context, *TagsDO) (*TagsDO, error)
	Update(context.Context, *TagsDO) (*TagsDO, error)
	//Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*TagsDO, error)
	FindOneByName(context.Context, string) (*TagsDO, error)
	CountByIdentifier(ctx context.Context, str string) (int64, error)
	ListAll(c context.Context, opts TagsDOListOption) (*TagsDOList, error)
	FindAllByArticleID(ctx context.Context, articleId uint64) ([]*TagsDO, error)
	FindOneByIdentifier(ctx context.Context, name string) (*TagsDO, error)
}

type TagsUsecase struct {
	repo TagsRepo
	log  *log.Helper
}

func NewTagsUsecase(repo TagsRepo, logger log.Logger) *TagsUsecase {
	return &TagsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TagsUsecase) Create(ctx context.Context, obj *TagsDO) (*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Name)
	if obj.Identifier == "" || len(obj.Identifier) == 0 {
		obj.Identifier = stringUtil.PinyinConvert(obj.Name)
	}
	count, _ := uc.repo.CountByIdentifier(ctx, obj.Identifier)
	if count > 0 {
		obj.Identifier = fmt.Sprintf("%s-%d", obj.Identifier, count)
	}
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	return objDO, nil
}

// Update 更新用户
func (uc *TagsUsecase) Update(ctx context.Context, obj *TagsDO) (*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Name)
	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除用户
//func (uc *TagsUsecase) Delete(ctx context.Context, id uint64) error {
//	uc.log.WithContext(ctx).Infof("Delete: %v", id)
//	return uc.repo.Delete(ctx, id)
//}

// DeleteList 根据ID批量删除用户
func (uc *TagsUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}
func (uc *TagsUsecase) FindOneByName(ctx context.Context, name string) (*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByName: %v", name)
	obj, err := uc.repo.FindOneByName(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return obj, nil
}
func (uc *TagsUsecase) FindOneByIdentifier(ctx context.Context, name string) (*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByIdentifier: %v", name)
	obj, err := uc.repo.FindOneByIdentifier(ctx, name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return obj, nil
}

// FindOneByID 根据ID查询用户信息
func (uc *TagsUsecase) FindOneByID(ctx context.Context, id uint64) (*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (uc *TagsUsecase) ListAll(ctx context.Context, opts TagsDOListOption) (*TagsDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return objDOs, nil
}
func (uc *TagsUsecase) FindAllByArticleID(ctx context.Context, articleId uint64) ([]*TagsDO, error) {
	uc.log.WithContext(ctx).Infof("FindAllByArticleID")
	objDOs, err := uc.repo.FindAllByArticleID(ctx, articleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return objDOs, nil
}
