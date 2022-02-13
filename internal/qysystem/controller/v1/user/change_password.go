package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qy-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qy-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
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
