package main

import (
	"context"
	"yxb.com/gin-scaffold/v2/internal/app"
	"yxb.com/gin-scaffold/v2/pkg/logger"
)

const Version = "0.0.1"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")
	configFile := "configs/config.yaml"
	err := app.RunApp(ctx,
		app.WithConfigFile(configFile),
		app.WithVersion(Version))

	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}
