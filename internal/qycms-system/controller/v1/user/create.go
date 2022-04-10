package user

import (
	v1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-api/qycms-system/v1"
	core "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/meta/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Create(c *gin.Context) {
	log.L(c).Info("user create function called.")
	token := u.qycnf.GetToken()
	var r v1.User
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	r.Salt = token
	if err := u.srv.Users().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	//if errs :=
	core.WriteResponse(c, nil, r)

}
