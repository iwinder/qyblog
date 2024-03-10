package server

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Config struct {
	HttpServing     *HttpServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}
type HttpServingInfo struct {
	Address string
	Timeout time.Duration
}
type JwtInfo struct {
	// defaults to "iam jwt"
	Realm string
	// defaults to empty
	Key string
	// defaults to one hour
	Timeout time.Duration
	// defaults to zero
	MaxRefresh time.Duration
}

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		Jwt: &JwtInfo{
			Realm:      "iam jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
	}
}
