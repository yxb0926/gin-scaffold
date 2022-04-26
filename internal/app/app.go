package app

import (
	"context"
	"fmt"
	"yxb.com/gin-scaffold/v2/internal/app/config"
)

type options struct {
	ConfigFile string
	Version    string
}
type Options func(*options)

func WithConfigFile(s string) Options {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func WithVersion(v string) Options {
	return func(o *options) {
		o.Version = v
	}
}

func Init(ctx context.Context, opts ...Options) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	// 读取配置文件内容
	config.LoadConf(o.ConfigFile)
	config.PrintWithJson()

	// 初始化日志组件
	InitLogger()
	// 依赖注入
	injector, err := BuildInjector()
	if err != nil {
		fmt.Errorf("%s\n", "BuildInjector Failed, Please")
	}

	addr := fmt.Sprintf("%s:%d", config.C.HttpServer.Host, config.C.HttpServer.Port)
	if err = injector.Engine.Run(addr); err != nil {
		fmt.Errorf("%s\n", "run http server failed, please check")
	}
}

func RunApp(ctx context.Context, opts ...Options) error {
	fmt.Println("Welcome use gin-scaffold project. Play Happy!")
	Init(ctx, opts...)
	return nil
}
