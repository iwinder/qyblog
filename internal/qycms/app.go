package qycms

import (
	app "gitee.com/windcoder/qingyucms/internal/pkg/qycms-app"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"gitee.com/windcoder/qingyucms/internal/qycms/config"
	"gitee.com/windcoder/qingyucms/internal/qycms/options"
)

const commandDesc = "QYCMS SYSTEM"

func NewApp(basename string) *app.App {
	// 获取配置信息
	opts := options.NewDefaultOptions()
	application := app.NewApp(" QYCMS SYSTEM Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)
	return application
}

// 初始化定制方法
func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		// 日志初始化
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
