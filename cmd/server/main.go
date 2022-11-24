package main

import (
	"context"
	"flag"

	"github.com/xiaoyan648/project-layout/internal/pkg/conf"
	"github.com/xiaoyan648/project-layout/internal/service"

	"github.com/go-leo/leo"
	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/config/medium/file"
	"github.com/go-leo/leo/config/parser"
	"github.com/go-leo/leo/config/valuer"
	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/log/zap"
)

var (
	configPath    string
	logLevel      string
	logFormatJSON bool
)

func init() {
	flag.StringVar(&configPath, "conf", "/configs/nacos-test.yaml", "config path, eg: -conf config.yaml")
	flag.StringVar(&logLevel, "log", "debug", "log level, eg: -log [debug|info|warn|error]")
	flag.BoolVar(&logFormatJSON, "log_json", false, "log format, eg: -log_json [ture|false]")
}

const (
	// ID app id.
	ID = "10000"
	// Name app name.
	Name = "github.com/xiaoyan648/project-layout"
	// Version app version.
	Version = "v1"
)

func main() {
	flag.Parse()
	// scan config.
	appConf, err := GetAppConfig()
	if err != nil {
		panic(err)
	}

	// new logger.
	logFormatOpt := zap.PlainText()
	if logFormatJSON {
		logFormatOpt = zap.JSON()
	}
	logger := zap.New(zap.LevelAdapt(log.Level(logLevel)), zap.Console(true), logFormatOpt)

	// new app.
	app, cleanup, err := initApp(appConf.Server, appConf.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	panic(app.Run(context.Background()))
}

func newApp(
	logger log.Logger, serverCfg *conf.Server, data *conf.Data,
	server *service.GreeterService,
) *leo.App {
	// init server
	httpOpt := NewHttpServer(serverCfg, server, logger)
	// grpcOpt := NewGrpcServer(serverCfg, server, logger)

	// new app
	return leo.NewApp(
		leo.ID(ID), leo.Name(Name), leo.Version(Version),
		leo.Logger(logger),
		leo.HTTP(httpOpt),
		// leo.GRPC(grpcOpt),
	)
}

func GetAppConfig() (conf.Content, error) {
	appConfig := new(conf.Content)
	fileConfMgr := config.NewManager(
		config.WithLoader(file.NewLoader(configPath)),
		config.WithParser(parser.NewYamlParser()),
		config.WithValuer(valuer.NewTrieTreeValuer()),
	)
	if err := fileConfMgr.ReadConfig(); err != nil {
		panic(err)
	}

	if err := fileConfMgr.Unmarshal(appConfig); err != nil {
		panic(err)
	}

	return *appConfig, nil
}
