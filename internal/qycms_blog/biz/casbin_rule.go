package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
)

type CasbinRuleDO struct {
	metaV1.ObjectNoInstMeta
	PType string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

type CasbinRuleRepo interface {
	SaveRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error)
	UpdateRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error)
	SavePolicies(ctx context.Context, rules [][]string) (bool, error)
	CleanPolicy(ctx context.Context, p ...string) (bool, error)
	UpdatePolicies(ctx context.Context, oldApi, newApi *ApiDO) (bool, error)
	DeleteRoleForUser(ctx context.Context, user string, domain ...string) (bool, error)
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
