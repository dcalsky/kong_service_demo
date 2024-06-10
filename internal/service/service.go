package service

import (
	"github.com/dcalsky/kong_service_demo/internal/adapter"
	"github.com/dcalsky/kong_service_demo/internal/config"
	"github.com/dcalsky/kong_service_demo/internal/service/kong_service"
	"github.com/dcalsky/kong_service_demo/internal/service/session"
)

var (
	KongServiceSvc kong_service.IKongService
	SessionSvc     session.ISessionService
)

func MustInit() {
	KongServiceSvc = kong_service.NewKongService(adapter.KongServiceRepo)
	SessionSvc = session.NewSessionService(adapter.AccountRepo, config.Conf.KongSecret)
}
