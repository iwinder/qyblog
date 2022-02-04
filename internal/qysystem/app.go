package qysystem

import (
	app "gitee.com/windcoder/qingyucms/internal/pkg/qy-app"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	"gitee.com/windcoder/qingyucms/internal/qysystem/config"
	"gitee.com/windcoder/qingyucms/internal/qysystem/options"
)

const commandDesc = "QYCMS SYSTEM"

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(" QYCMS SYSTEM Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
