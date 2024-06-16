package main

import (
	"github.com/dcalsky/kong_service_demo/internal/adapter"
	"github.com/dcalsky/kong_service_demo/internal/config"
	"github.com/dcalsky/kong_service_demo/internal/service"
)

func main() {
	config.MustInit()
	adapter.MustInit()
	service.MustInit()
	RegisterRoutes()
	adapter.Release()
}
