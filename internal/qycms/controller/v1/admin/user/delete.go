package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

// Delete 删除用户
func (u *UserController) Delete(c *gin.Context) {
	log.L(c).Info("delete user function called.")

	if err := u.srv.Users().Delete(c, c.Param("username"), metav1.DeleteOptions{
		Unscoped: false,
	}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}

//// DeleteCollection 真删除
//func (u *UserController) DeleteCollection(c *gin.Context) {
//	log.L(c).Info("batch delete user function called.")
//
//	usernames := c.QueryArray("name")
//
//	if err := u.srv.Users().DeleteCollection(c, usernames, metav1.DeleteOptions{}); err != nil {
//		core.WriteResponse(c, err, nil)
//
//		return
//	}
//
//	core.WriteResponse(c, nil, nil)
//}
