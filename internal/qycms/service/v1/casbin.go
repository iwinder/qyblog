package v1

import (
	"gitee.com/windcoder/qingyucms/internal/qycms/store"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

type CasbinSrv interface {
}
type CasbinService struct {
	store store.Factory
}

func (c *CasbinService) GetCasbinOr() {
	Apter, _ := gormadapter.NewAdapterByDB(store.GetClient().CommonDB().GetCommonDB())
	//syncedEnforcer, _ := casbin.NewEnforcer("", a)
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
}
