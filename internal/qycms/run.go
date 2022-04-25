package qycms

import "gitee.com/windcoder/qingyucms/internal/qycms/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
