package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type casbinRuleRepo struct {
	cdata *CasbinData
	log   *log.Helper
}

func NewCasbinRuleRepo(cdata *CasbinData, logger log.Logger) biz.CasbinRuleRepo {
	return &casbinRuleRepo{
		cdata: cdata,
		log:   log.NewHelper(logger),
	}
}

func (r *casbinRuleRepo) SaveRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error) {
	return r.cdata.Enf.AddRolesForUser(user, roles, domain...)
}

func (r *casbinRuleRepo) UpdateRoleForUser(ctx context.Context, user string, roles []string, domain ...string) (bool, error) {
	flag, err := r.cdata.Enf.DeleteRolesForUser(user, domain...)

	if flag || (err == nil && !flag) {
		return r.cdata.Enf.AddRolesForUser(user, roles, domain...)
	}
	return false, err
}

func (r *casbinRuleRepo) SavePolicies(ctx context.Context, rules [][]string) (bool, error) {
	return r.cdata.Enf.AddPolicies(rules)
}

func (r *casbinRuleRepo) CleanPolicy(ctx context.Context, p ...string) (bool, error) {
	success, err := r.cdata.Enf.RemoveFilteredPolicy(0, p...)
	return success, err
}

func (r *casbinRuleRepo) UpdatePolicies(ctx context.Context, oldApi, newApi *biz.ApiDO) (bool, error) {
	newPO := &po.CasbinRulePO{
		V0: oldApi.Identifier,
		V1: oldApi.Path,
		V2: oldApi.Method,
	}
	err := r.cdata.data.Db.Model(newPO).Where("v0 = ? AND v1 = ? AND v2 = ?", oldApi.ApiGroup, oldApi.Path, oldApi.Method).Updates(newPO).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *casbinRuleRepo) DeleteRoleForUser(ctx context.Context, user string, domain ...string) (bool, error) {
	flag, err := r.cdata.Enf.DeleteRolesForUser(user, domain...)

	if !flag || err != nil {
		return false, err
	}
	return true, nil
}

//func (r *casbinRuleRepo) SaveRoleForUser(ctx context.Context, obj *biz.RoleDO) (*po.RolePO, error) {
//	//r.cdata.Enf.AddRoleForUser()
//	//objPO := &po.RolePO{
//	//	ObjectMeta: obj.ObjectMeta,
//	//	Name:       obj.Name,
//	//	Identifier: obj.Identifier,
//	//}
//	//err := r.cdata.Db.Create(objPO).Error
//	//if err != nil {
//	//	return nil, err
//	//}
//	//return objPO, nil
//}
