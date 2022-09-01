package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

type casbinRuleRepo struct {
	data *CasbinData
	log  *log.Helper
}

func NewCasbinRuleRepo(data *CasbinData, logger log.Logger) biz.CasbinRuleRepo {
	return &casbinRuleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *casbinRuleRepo) SaveRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error) {
	return r.data.Enf.AddRolesForUser(user, roles, domain...)
}

func (r *casbinRuleRepo) UpdateRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error) {
	flag, err := r.data.Enf.DeleteRolesForUser(user, domain...)
	if flag {
		return r.data.Enf.AddRolesForUser(user, roles, domain...)
	}
	return false, err
}

func (r *casbinRuleRepo) SavePolicies(ctx context.Context, rules [][]string) (bool, error) {
	return r.data.Enf.AddPolicies(rules)
}

func (r *casbinRuleRepo) CleanPolicy(ctx context.Context, p ...string) bool {
	success, _ := r.data.Enf.RemoveFilteredPolicy(0, p...)
	return success
}

//func (r *casbinRuleRepo) SaveRoleForUser(ctx context.Context, obj *biz.RoleDO) (*po.RolePO, error) {
//	//r.data.Enf.AddRoleForUser()
//	//objPO := &po.RolePO{
//	//	ObjectMeta: obj.ObjectMeta,
//	//	Name:       obj.Name,
//	//	Identifier: obj.Identifier,
//	//}
//	//err := r.data.Db.Create(objPO).Error
//	//if err != nil {
//	//	return nil, err
//	//}
//	//return objPO, nil
//}
