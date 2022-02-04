package qy_server

import "github.com/gin-gonic/gin"

type Config struct {
	Mode            string
	Middlewares     []string
	EnableProfiling bool
	EnableMetrics   bool
	Healthz         bool
	InsecureServing *InsecureServingInfo
}

type InsecureServingInfo struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		Healthz:         true,
		EnableProfiling: true,
		EnableMetrics:   true,
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
