package qycms_system

import "gitee.com/windcoder/qingyucms/internal/qycms-system/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
