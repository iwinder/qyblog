package user

import (
	"gitee.com/windcoder/qingyucms/internal/qycms/config"
	srvv1 "gitee.com/windcoder/qingyucms/internal/qycms/service/v1"
	"gitee.com/windcoder/qingyucms/internal/qycms/store"
)

type UserController struct {
	srv   srvv1.Service
	qycnf config.QyComConfig
}

func NewUserController(store store.Factory) *UserController {
	aqycnf, _ := config.GetQyComConfigOr(nil)

	return &UserController{srv: srvv1.NewService(store),
		qycnf: aqycnf,
	}
}
