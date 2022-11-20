package http

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	nhttp "net/http"
	"strings"
)

func LocalHttpRequestFilter() http.FilterFunc {
	return func(next nhttp.Handler) nhttp.Handler {
		return nhttp.HandlerFunc(func(w nhttp.ResponseWriter, req *nhttp.Request) {
			req.Header.Add("X-RemoteAddr", strings.Split(req.RemoteAddr, ":")[0])

			next.ServeHTTP(w, req)
		})
	}
}
