package user

import (
	srvv1 "gitee.com/windcoder/qingyucms/internal/qysystem/service/v1"
	"gitee.com/windcoder/qingyucms/internal/qysystem/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{srv: srvv1.NewService(store)}
}
