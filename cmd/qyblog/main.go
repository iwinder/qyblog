package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	qy_logger "github.com/iwinder/qyblog/internal/pkg/qycms_common/qy-logger"
	"github.com/iwinder/qyblog/internal/qycms_blog/conf"
	"github.com/iwinder/qyblog/internal/qycms_blog/job"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, job *job.QyCronJob) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			//gs,
			hs,
			job,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	//f, lerr := os.OpenFile(bc.Qycms.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if lerr != nil {
	//	panic(lerr)
	//}
	//writeSyncer := zapcore.AddSync(f)
	//defLevel := zapcore.DebugLevel
	//if len(bc.Qycms.Logger.Level) > 0 {
	//	defLevel.Set(bc.Qycms.Logger.Level)
	//}
	cof := zap.NewProductionEncoderConfig()
	cof.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(cof)
	encoder.AddString("service.id", id)
	encoder.AddString("service.name", Name)
	encoder.AddString("service.version", Version)
	encoder.AddReflected("trace.id", tracing.TraceID())
	encoder.AddReflected("span.id", tracing.SpanID())
	logCOnfig := qy_logger.LoggerConfig{
		Mode:       bc.Qycms.Mode,
		Path:       bc.Qycms.Logger.Path,
		Level:      bc.Qycms.Logger.Level,
		MaxSize:    bc.Qycms.Logger.MaxSize,
		MaxBackups: bc.Qycms.Logger.MaxBackups,
		MaxAge:     bc.Qycms.Logger.MaxAge,
		Compress:   bc.Qycms.Logger.Compress,
	}
	logger := qy_logger.Logger(encoder, logCOnfig)
	//core := zapcore.NewCore(encoder, writeSyncer, defLevel)
	//z := zap.New(core)
	//logger := kratoszap.NewLogger(z)
	//logger := log.With(log.NewHelper(),
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", id,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"trace.id", tracing.TraceID(),
	//	"span.id", tracing.SpanID(),
	//)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Qycms, bc.Auth, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
