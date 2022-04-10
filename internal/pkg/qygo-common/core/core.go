package core

import (
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
