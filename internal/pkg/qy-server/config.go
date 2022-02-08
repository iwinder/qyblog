package qy_server

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	// RecommendedHomeDir defines the default directory used to place all iam service configurations.
	RecommendedHomeDir = ".qycms"

	// RecommendedEnvPrefix defines the ENV prefix used by all iam service.
	RecommendedEnvPrefix = "QYCMS"
)

type Config struct {
	Mode            string
	Middlewares     []string
	EnableProfiling bool
	EnableMetrics   bool
	Healthz         bool
	InsecureServing *InsecureServingInfo
	Jwt             *JwtInfo
}

type InsecureServingInfo struct {
	Address string
}

type JwtInfo struct {
	Realm      string
	Key        string
	Timeout    time.Duration
	MaxRefresh time.Duration
}

func NewConfig() *Config {
	return &Config{
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		Healthz:         true,
		EnableProfiling: true,
		EnableMetrics:   true,
		Jwt: &JwtInfo{
			Realm:      "qycms jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
	}
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() (*GenericAPIServer, error) {
	s := &GenericAPIServer{
		Engine:              gin.New(),
		mode:                c.Mode,
		middlewares:         c.Middlewares,
		InsecureServingInfo: c.InsecureServing,
		healthz:             c.Healthz,
		enableMetrics:       c.EnableMetrics,
		enableProfiling:     c.EnableProfiling,
	}
	initGenericAPIServer(s)

	return s, nil
}
