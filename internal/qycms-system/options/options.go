package options

import (
	"encoding/json"
	genericoption "gitee.com/windcoder/qingyucms/internal/pkg/qycms-options"
	server "gitee.com/windcoder/qingyucms/internal/pkg/qycms-server"
	cliflag "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/cli/flag"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/idUtil"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
)

type Options struct {
	GenericServerRunOptions *genericoption.ServerRunOptions       `json:"server" mapstructure:"server"`
	InsecureServing         *genericoption.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	GRPCOptions             *genericoption.GRPCOptions            `json:"grpc" mapstructure:"grpc"`
	MySQLOptions            *genericoption.MySQLOptions           `json:"mysql" mapstructure:"mysql"`
	RedisOptions            *genericoption.RedisOptions           `json:"redis" mapstructure:"redis"`
	JwtOptions              *genericoption.JwtOptions             `json:"jwt" mapstructure:"jwt"`
	Log                     *log.Options                          `json:"log" mapstructure:"log"`
	QyOptions               *genericoption.QyOptions              `json:"qy" mapstructure:"qy"`
}

func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoption.NewServerRunOptions(),
		InsecureServing:         genericoption.NewInsecureServingOptions(),
		GRPCOptions:             genericoption.NewGRPCOptions(),
		MySQLOptions:            genericoption.NewMySQLOptions(),
		RedisOptions:            genericoption.NewRedisOptions(),
		JwtOptions:              genericoption.NewJwtOptions(),
		Log:                     log.NewOptions(),
		QyOptions:               genericoption.NewQyOptions(),
	}
	return &o
}

func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.JwtOptions.AddFlags(fss.FlagSet("jwt"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	o.QyOptions.AddFlags(fss.FlagSet("qy"))
	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

func (o *Options) Complete() error {
	if o.JwtOptions.Key == "" {
		o.JwtOptions.Key = idUtil.NewSecretKey()
	}
	return nil
}
