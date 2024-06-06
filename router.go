package main

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/dcalsky/kong_service_demo/internal/handler"
	"github.com/dcalsky/kong_service_demo/internal/middleware"
)

func RegisterRoutes() {
	r := server.Default(
		server.WithHostPorts("0.0.0.0:8002"),
	)
	r.Use(middleware.TrafficLogger())
	r.Use(middleware.ExceptionGuard())
	registerHttp(r)

	r.Spin()
}

func registerHttp(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(http.StatusNotFound, nil)
	})
}
