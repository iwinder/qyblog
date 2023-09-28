package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/auth/auth_constants"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"gorm.io/gorm"
	"strconv"
)

var (
	// ErrApiNotFound is api not found.
	ErrApiNotFound = errors.NotFound("101404", "api not found")
)

type ApiDO struct {
	metaV1.ObjectMeta
	ApiGroup    string
	Method      string
	Path        string
	Description string
	Identifier  string
	GroupId     uint64
}
type ApiTreeDO struct {
	ID          string
	ApiGroup    string
	Method      string
	Path        string
	Description string
	Identifier  string
	GroupId     uint64
	Children    []*ApiTreeDO
}
type ApiDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ApiDO `json:"items"`
}

type ApiDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ApiDO              `json:"item"`
}

type ApiRepo interface {
	Save(context.Context, *ApiDO) (*ApiDO, error)
	Update(context.Context, *ApiDO) (*ApiDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*ApiDO, error)
	ListAll(c context.Context, opts ApiDOListOption) (*ApiDOList, error)
}

type ApiUsecase struct {
	repo     ApiRepo
	acg      *ApiGroupUsecase
	casbRepo CasbinRuleRepo
	log      *log.Helper
}

func NewApiUsecase(repo ApiRepo, acg *ApiGroupUsecase, casbRepo CasbinRuleRepo, logger log.Logger) *ApiUsecase {
	return &ApiUsecase{repo: repo, acg: acg, casbRepo: casbRepo, log: log.NewHelper(logger)}
}

func (uc *ApiUsecase) Create(ctx context.Context, obj *ApiDO) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", obj.Path)
	objDO, err := uc.repo.Save(ctx, obj)
	if err != nil {
		return nil, err
	}
	return objDO, nil
}

// Update 更新用户
func (uc *ApiUsecase) Update(ctx context.Context, obj *ApiDO) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", obj.Path)
	oldDO, err := uc.FindOneByID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	_, cerr := uc.casbRepo.UpdatePolicies(ctx, oldDO, obj)
	if cerr != nil {
		return nil, err
	}

	objDO, err := uc.repo.Update(ctx, obj)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiNotFound
		}
		return nil, err
	}
	return objDO, nil
}

// Delete 根据ID删除用户
func (uc *ApiUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除用户
func (uc *ApiUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *ApiUsecase) FindOneByID(ctx context.Context, id uint64) (*ApiDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	obj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiNotFound
		}
		return nil, err
	}
	return obj, nil
}

// ListAll 批量查询
func (uc *ApiUsecase) ListAll(ctx context.Context, opts ApiDOListOption) (*ApiDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	objDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrApiNotFound
		}
		return nil, err
	}

	return objDOs, nil
}
func (uc *ApiUsecase) TreeAll(ctx context.Context) ([]*ApiTreeDO, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	// 获取分组
	acgOpt := ApiGroupDOListOption{}
	acgOpt.PageFlag = false
	acgData, acgErr := uc.acg.ListAll(ctx, acgOpt)
	if acgErr != nil {
		return nil, acgErr
	}
	opts := ApiDOListOption{}
	opts.PageFlag = false
	var result []*ApiTreeDO
	acgLen := len(acgData.Items)
	if acgLen > 0 {
		result = make([]*ApiTreeDO, 0, acgLen)
		for _, obj := range acgData.Items {
			parent := &ApiTreeDO{
				ID:          auth_constants.PrefixApiGroup + strconv.FormatUint(obj.ID, 10),
				Description: obj.Name,
			}
			opts.GroupId = obj.ID
			objDOs, err := uc.repo.ListAll(ctx, opts)
			if err != nil {
				uc.log.Error(err.Error())
				parent.Children = make([]*ApiTreeDO, 0, 0)
				continue
			}
			child := make([]*ApiTreeDO, 0, len(objDOs.Items))
			for _, cobj := range objDOs.Items {
				child = append(child, &ApiTreeDO{
					ID:          strconv.FormatUint(cobj.ID, 10),
					Path:        cobj.Path,
					Method:      cobj.Method,
					Identifier:  cobj.Identifier,
					Description: cobj.Description,
				})
			}
			parent.Children = child
			result = append(result, parent)
		}
	} else {
		result = make([]*ApiTreeDO, 0, 0)
	}

	return result, nil
}
