package options

type Config struct {
	*Options
}

func CreateConfigFromOptions(opts *Options) (*Config, error) {
	return &Config{opts}, nil
}
