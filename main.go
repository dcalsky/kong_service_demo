package main

import (
	"go.uber.org/zap/zapcore"

	"github.com/dcalsky/kong_service_demo/internal/adapter"
	"github.com/dcalsky/kong_service_demo/internal/common/logs"
	"github.com/dcalsky/kong_service_demo/internal/config"
	"github.com/dcalsky/kong_service_demo/internal/service"
)

func main() {
	logs.MustInit(zapcore.InfoLevel)
	config.MustInit()
	adapter.MustInit()
	service.MustInit()
	RegisterRoutes()
	adapter.Release()
}
