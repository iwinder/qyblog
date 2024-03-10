package filter

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	qyfhttp "github.com/iwinder/qyblog/internal/pkg/filter/http"
)

func NewFilter() http.ServerOption {
	return http.Filter(qyfhttp.LocalHttpRequestFilter())

}
