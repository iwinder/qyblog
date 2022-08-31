package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
)

type CasbinRuleDO struct {
	metaV1.ObjectMeta
	Name       string
	Identifier string
}

type CasbinRuleRepo interface {
	//Save(context.Context, *RoleDO) (*po.RolePO, error)
	//Update(context.Context, *RoleDO) (*po.RolePO, error)
	//Delete(context.Context, uint64) error
	//DeleteList(c context.Context, uids []uint64) error
	//FindByID(context.Context, uint64) (*po.RolePO, error)
	//FindByKey(c context.Context, key string) (*po.RolePO, error)
	//ListAll(c context.Context, opts UserDOListOption) (*po.UserPOList, error)
}

type CasbinRuleUsecase struct {
	repo CasbinRuleRepo
	log  *log.Helper
}

func NewCasbinRuleUsecase(repo CasbinRuleRepo, logger log.Logger) *CasbinRuleUsecase {
	return &CasbinRuleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 增加API权限 _A

// 增加菜单权限 _M
