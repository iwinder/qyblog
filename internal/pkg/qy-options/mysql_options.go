package qy_options

import "time"

type MySQLOptions struct {
	Host                  string        `json:"host,omitempty" mapstructure:"host"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Passowrd              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdlConnections     int           `json:"max-idl-connections,omitempty" mapstructure:"max-idl-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level" mapstructure:"log-level"`
}
