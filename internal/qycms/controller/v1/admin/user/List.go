package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/core"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	v1 "gitee.com/windcoder/qingyucms/internal/qycms/models/v1"
	"github.com/gin-gonic/gin"
)

// List 列表查询
func (u *UserController) List(c *gin.Context) {
	log.L(c).Info("list user function called.")

	var r v1.UserListOption
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	users, err := u.srv.Users().List(c, r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, users)
}
