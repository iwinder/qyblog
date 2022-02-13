package user

import (
	v1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-api/qysystem/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qy-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qy-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Update(c *gin.Context) {
	log.L(c).Info("update user function called.")
	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	user, err := u.srv.Users().Get(c, c.Param("username"), metav1.GetOptions{Unscoped: true})

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.Extend = r.Extend
	//if errs := user.ValidateUpdate(); len(errs) != 0 {
	//	core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)
	//
	//	return
	//}

	if err := u.srv.Users().Update(c, user, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
