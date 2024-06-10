package service

import (
	"github.com/dcalsky/kong_service_demo/internal/adapter"
	"github.com/dcalsky/kong_service_demo/internal/service/kong_service"
)

var (
	KongServiceSvc kong_service.IKongService
)

func MustInit() {
	KongServiceSvc = kong_service.NewKongService(adapter.KongServiceRepo)
}
