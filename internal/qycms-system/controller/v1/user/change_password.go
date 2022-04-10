package user

import (
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/meta/v1"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	// Old password.
	// Required: true
	OldPassword string `json:"oldPassword" binding:"omitempty"`

	// New password.
	// Required: true
	NewPassword string `json:"newPassword" binding:"password"`
}

func (u *UserController) ChangePassword(c *gin.Context) {
	log.L(c).Info("change password function called.")

	var r ChangePasswordRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	user, err := u.srv.Users().Get(c, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if err := user.Compare(r.OldPassword); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrPasswordIncorrect, err.Error()), nil)

		return
	}

	user.Password = r.NewPassword
	if err := u.srv.Users().ChangePassword(c, user); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
