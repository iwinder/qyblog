package user

import (
	"gitee.com/windcoder/qingyucms/internal/qysystem/config"
	srvv1 "gitee.com/windcoder/qingyucms/internal/qysystem/service/v1"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
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
