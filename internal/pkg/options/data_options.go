package options

import (
	"github.com/iwinder/qyblog/internal/pkg/options/base"
	"github.com/spf13/pflag"
)

type DataOptions struct {
	MySQLOptions *base.MySQLOptions `json:"server"   mapstructure:"server"`
	RedisOptions *base.RedisOptions `json:"redis"    mapstructure:"redis"`
}

func NewDataOptions() *DataOptions {
	return &DataOptions{
		MySQLOptions: base.NewMySQLOptions(),
		RedisOptions: base.NewRedisOptions(),
	}
}
func (o *DataOptions) AddFlags(fs *pflag.FlagSet) {
	o.MySQLOptions.AddFlags(fs)
	o.RedisOptions.AddFlags(fs)
}

func (o *DataOptions) Validate() []error {
	errs := []error{}
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.RedisOptions.Validate()...)
	return errs
}
