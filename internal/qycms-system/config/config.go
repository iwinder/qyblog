package config

import "gitee.com/windcoder/qingyucms/internal/qycms-system/options"

type Config struct {
	*options.Options
}

func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
