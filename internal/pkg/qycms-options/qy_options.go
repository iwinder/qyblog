package qycms_options

import "github.com/spf13/pflag"

type QyOptions struct {
	Token string `json:"token,omitempty" mapstructure:"token"`
}

func NewDefaultQyOptions() *QyOptions {
	return &QyOptions{
		Token: "",
	}
}

func (o *QyOptions) Validate() []error {
	errs := []error{}
	return errs
}

func (o *QyOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Token, "qycms.token", o.Token, ""+
		"token for access to QYCMS service.(QYCMS 服务的 token)")
}
