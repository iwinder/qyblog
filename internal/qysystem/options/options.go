package options

import (
	genericoption "gitee.com/windcoder/qingyucms/internal/pkg/qy-options"
)

type Options struct {
	GenericServerRunOptions *genericoption.ServerRunOptions `json:"server" mapstructure: "server"`
}
