package options

import (
	"encoding/json"
	cliflag "github.com/iwinder/qyblog/internal/pkg/cli/flag"
	genericoptions "github.com/iwinder/qyblog/internal/pkg/options"
	"github.com/iwinder/qyblog/internal/pkg/server"
	"github.com/iwinder/qyblog/internal/pkg/utils/idUtil"
)

type Options struct {
	ServingOptions *genericoptions.ServingOptions `json:"server"   mapstructure:"server"`
	DataOptions    *genericoptions.DataOptions    `json:"store"    mapstructure:"store"`
	QycmsOptions   *genericoptions.QycmsOptions   `json:"qycms"      mapstructure:"qycms"`
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.ServingOptions.Validate()...)
	errs = append(errs, o.DataOptions.Validate()...)
	errs = append(errs, o.QycmsOptions.Validate()...)
	return errs
}

func NewOptions() *Options {
	o := Options{
		ServingOptions: genericoptions.NewServingOptions(),
		DataOptions:    genericoptions.NewDataOptions(),
		QycmsOptions:   genericoptions.NewQycmsOptions(),
	}

	return &o
}
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ServingOptions.AddFlags(fss.FlagSet("server"))
	o.DataOptions.AddFlags(fss.FlagSet("store"))
	o.QycmsOptions.AddFlags(fss.FlagSet("qycms"))
	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Complete set default Options.
func (o *Options) Complete() error {
	if o.QycmsOptions.JwtOptions.Key == "" {
		o.QycmsOptions.JwtOptions.Key = idUtil.NewSecretKey()
	}

	return o.ServingOptions.HttpOptions.Complete()
}
