package qy_server

import "github.com/gin-gonic/gin"

type Config struct {
	Mode    string
	Healthz bool
}

type InsecureSeringInfo struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Mode:    gin.ReleaseMode,
		Healthz: true,
	}
}
