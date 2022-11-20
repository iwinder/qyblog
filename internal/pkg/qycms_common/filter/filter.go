package filter

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)
import qyfhttp "github.com/iwinder/qingyucms/internal/pkg/qycms_common/filter/http"

func NewFilter() http.ServerOption {
	return http.Filter(qyfhttp.LocalHttpRequestFilter())

}
