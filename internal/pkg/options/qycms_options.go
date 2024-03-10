package options

import (
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	"github.com/iwinder/qyblog/internal/pkg/options/base"
	"github.com/spf13/pflag"
)

type QycmsOptions struct {
	Token        string             `json:"token" mapstructure:"token"`
	JwtOptions   *base.JwtOptions   `json:"jwt"      mapstructure:"jwt"`
	Log          *log.Options       `json:"log"      mapstructure:"log"`
	EmailOptions *base.EmailOptions `json:"email" mapstructure:"email"`
}

func NewQycmsOptions() *QycmsOptions {
	return &QycmsOptions{
		Token:        "",
		JwtOptions:   base.NewJwtOptions(),
		Log:          log.NewOptions(),
		EmailOptions: base.NewEmailOptions(),
	}
}
func (o *QycmsOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Token, "qycms.token", o.Token, "加密 token.")
	o.JwtOptions.AddFlags(fs)
	o.Log.AddFlags(fs)
	o.EmailOptions.AddFlags(fs)

}

func (o *QycmsOptions) Validate() []error {
	errs := []error{}
	errs = append(errs, o.JwtOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	errs = append(errs, o.EmailOptions.Validate()...)
	return errs
}
