package main

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/dcalsky/kong_service_demo/internal/config"
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

	v1 := r.Group("/api/v1")
	// no auth
	{
		v1.POST("/Register", handler.Register)
		v1.POST("/Login", handler.Login)
	}
	jwtValidator := middleware.JwtValidator(config.Conf.KongSecret)
	v1.Use(jwtValidator)
	// kong service
	{
		v1.POST("/CreateKongService", handler.CreateKongService)
		v1.POST("/DescribeKongService", handler.DescribeKongService)
		v1.POST("/DeleteKongService", handler.DeleteKongService)
		v1.POST("/ListKongServices", handler.ListKongServices)
		v1.POST("/UpdateKongService", handler.UpdateKongService)
		v1.POST("/CreateKongServiceVersion", handler.CreateKongServiceVersion)
		v1.POST("/SwitchKongServiceVersion", handler.SwitchKongServiceVersion)
	}
	// organization
	{
		v1.POST("/CreateOrganization", handler.CreateOrganization)
		v1.POST("/DescribeOrganization", handler.DescribeOrganization)
		v1.POST("/AddAccountToOrganization", handler.AddAccountToOrganization)
		v1.POST("/RemoveAccountFromOrganization", handler.RemoveAccountFromOrganization)
	}
}
