package base

import (
	"fmt"
	"github.com/iwinder/qyblog/internal/pkg/server"
	"github.com/spf13/pflag"
	"strconv"
	"strings"
	"time"
)

type HttpServingOptions struct {
	Addr    string        `json:"addr" mapstructure:"addr"`
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

func NewHttpServingOptions() *HttpServingOptions {
	return &HttpServingOptions{
		Addr: "127.0.0.1:8080",
	}
}
func (s *HttpServingOptions) ApplyTo(c *server.Config) error {
	c.HttpServing = &server.HttpServingInfo{
		Address: s.Addr,
		Timeout: s.Timeout,
	}
	return nil
}
func (s *HttpServingOptions) Complete() error {
	return nil
}
func (s *HttpServingOptions) Validate() []error {
	var errors []error

	i := strings.LastIndex(s.Addr, ":")
	// 如果没有设置端口号，自动赋值默认端口
	if i < 0 {
		s.Addr = s.Addr + ":8080"
		i = strings.LastIndex(s.Addr, ":")
	}
	portStr := s.Addr[i+1:]

	port, err := strconv.Atoi(portStr)
	if err != nil {
		errors = append(
			errors,
			err,
		)
	}

	if port < 0 || port > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--http.addr %v 's port must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
				s.Addr,
			),
		)
	}

	return errors
}

func (s *HttpServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Addr, "server.http.addr", s.Addr, ""+
		"通过 --server.http.addr 设置 ip:port 形式的地址"+
		"(set to 0.0.0.0:8080 for all IPv4 interfaces).")
	fs.DurationVar(&s.Timeout, "server.http.timeout", s.Timeout, "通过 --server.http.timeout 设置 HttpOptions request timeout.")
}
