package service

import (
	"github.com/dcalsky/kong_service_demo/internal/adapter"
	"github.com/dcalsky/kong_service_demo/internal/config"
	"github.com/dcalsky/kong_service_demo/internal/service/kong_service"
	"github.com/dcalsky/kong_service_demo/internal/service/organization"
	"github.com/dcalsky/kong_service_demo/internal/service/session"
)

var (
	KongServiceSvc  kong_service.IKongService
	SessionSvc      session.ISessionService
	OrganizationSvc organization.IOrganizationService
)

func MustInit() {
	KongServiceSvc = kong_service.NewKongService(adapter.KongServiceRepo, adapter.AccountRepo, adapter.RepoHelper)
	SessionSvc = session.NewSessionService(adapter.AccountRepo, config.Conf.KongSecret)
	OrganizationSvc = organization.NewOrganizationService(adapter.AccountRepo, adapter.OrganizationRepo)
}
