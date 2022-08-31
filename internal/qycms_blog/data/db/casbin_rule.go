package db

import (
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