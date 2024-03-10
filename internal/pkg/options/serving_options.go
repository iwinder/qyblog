package options

import (
	"github.com/iwinder/qyblog/internal/pkg/options/base"
	"github.com/spf13/pflag"
)

type ServingOptions struct {
	HttpOptions *base.HttpServingOptions `json:"http" mapstructure:"http"`
}

func NewServingOptions() *ServingOptions {
	return &ServingOptions{
		HttpOptions: base.NewHttpServingOptions(),
	}
}
func (s *ServingOptions) Validate() []error {
	var errors []error
	errors = s.HttpOptions.Validate()
	if errors != nil {
		return errors
	}
	return nil
}
func (o *ServingOptions) AddFlags(fs *pflag.FlagSet) {
	o.HttpOptions.AddFlags(fs)
}
