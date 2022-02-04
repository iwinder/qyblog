package options

import (
	"encoding/json"
	cliflag "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/cli/flag"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	genericoption "gitee.com/windcoder/qingyucms/internal/pkg/qy-options"
	server "gitee.com/windcoder/qingyucms/internal/pkg/qy-server"
)

type Options struct {
	GenericServerRunOptions *genericoption.ServerRunOptions       `json:"server" mapstructure: "server"`
	InsecureServing         *genericoption.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	MySQLOptions            *genericoption.MySQLOptions           `json:"mysql" mapstructure:"mysql"`
	Log                     *log.Options                          `json:"log" mapstructure:"log"`
}

func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoption.NewServerRunOptions(),
		InsecureServing:         genericoption.NewInsecureServingOptions(),
		MySQLOptions:            genericoption.NewMySQLOptions(),
		Log:                     log.NewOptions(),
	}
	return &o
}

func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
