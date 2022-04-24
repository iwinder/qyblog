package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

func (u *UserController) List(c *gin.Context) {
	log.L(c).Info("list user function called.")

	var r metav1.ListOptions
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
