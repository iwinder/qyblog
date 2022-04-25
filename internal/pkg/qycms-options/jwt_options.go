package qycms_options

import (
	"fmt"
	server "gitee.com/windcoder/qingyucms/internal/pkg/qycms-server"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/pflag"
	"time"
)

type JwtOptions struct {
	Realm      string        `json:"realm" mapstructure:"realm"`
	Key        string        `json:"key" mapstructure:"key"`
	Timeout    time.Duration `json:"timeout" mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

func NewJwtOptions() *JwtOptions {
	defaults := server.NewConfig()

	return &JwtOptions{
		Realm:      defaults.Jwt.Realm,
		Key:        defaults.Jwt.Key,
		Timeout:    defaults.Jwt.Timeout,
		MaxRefresh: defaults.Jwt.MaxRefresh,
	}
}

func (j *JwtOptions) ApplyTo(c *server.Config) error {
	c.Jwt = &server.JwtInfo{
		Realm:      j.Realm,
		Key:        j.Key,
		Timeout:    j.Timeout,
		MaxRefresh: j.MaxRefresh,
	}
	return nil
}

func (j *JwtOptions) Validate() []error {
	var errs []error
	if !govalidator.StringLength(j.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("--secret-key must larger than 5 and little than 33"))
	}
	return errs
}

func (j *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}
	fs.StringVar(&j.Realm, "jwt.realm", j.Realm, "Realm name to display to the user.")
	fs.StringVar(&j.Key, "jwt.key", j.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&j.Timeout, "jwt.timeout", j.Timeout, "JWT token timeout.")

	fs.DurationVar(&j.MaxRefresh, "jwt.max-refresh", j.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")
}
