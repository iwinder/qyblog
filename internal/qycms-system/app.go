package qycms_system

import (
	"github.com/iwinder/qyblog/internal/pkg/app"
	log "github.com/iwinder/qyblog/internal/pkg/logger"
	"github.com/iwinder/qyblog/internal/qycms-system/options"
)

const commandDesc = `ff`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithOptions(opts),
		app.WithDesc(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}
func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.QycmsOptions.Log)
		defer log.Flush()

		cfg, err := options.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
