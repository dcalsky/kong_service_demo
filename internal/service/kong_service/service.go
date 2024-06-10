package kong_service

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type IKongService interface {
	CreateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceRequest) dto.CreateKongServiceResponse
	ListKongServices(ctx context.Context, kongArgs base.KongArgs, req dto.ListKongServicesRequest) dto.ListKongServicesResponse
	DescribeKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeKongServiceRequest) dto.DescribeKongServiceResponse
	UpdateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.UpdateKongServiceRequest) dto.UpdateKongServiceResponse
	DeleteKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DeleteKongServiceRequest)
}

type kongService struct {
	validator
	mapper
	kongServiceRepo repo.IKongServiceRepo
}

func NewKongService(kongServiceRepo repo.IKongServiceRepo) IKongService {
	s := &kongService{
		validator:       validator{},
		mapper:          mapper{},
		kongServiceRepo: kongServiceRepo,
	}
	return s
}

func (s *kongService) CreateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceRequest) dto.CreateKongServiceResponse {
	s.validator.ValidateCreateKongService(req)
	ks := entity.NewKongService()
	err := s.kongServiceRepo.CreateService(ctx, &ks)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	return dto.CreateKongServiceResponse{
		Service: ks.ToForDetail(),
	}
}

func (s *kongService) ListKongServices(ctx context.Context, kongArgs base.KongArgs, req dto.ListKongServicesRequest) dto.ListKongServicesResponse {
	services, err := s.kongServiceRepo.ListServices(ctx, repo.ListServicesRequest{
		Pagination: req.Pagination,
		SortBy:     req.SortBy,
		All:        false,
	})
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	return dto.ListKongServicesResponse{
		Services: lo.Map(services, func(svc entity.KongService, index int) dto.KongServiceForDetail {
			return svc.ToForDetail()
		}),
	}
}

func (s *kongService) DescribeKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeKongServiceRequest) dto.DescribeKongServiceResponse {
	s.validator.ValidateDescribeKongService(req)

	ks, err := s.kongServiceRepo.DescribeService(ctx, entity.KongServiceId(req.Id))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	base.PanicIf(ks == nil, base.NotFoundKongService.WithRawError(fmt.Errorf("kong service not found: %d", req.Id)))
	// TODO: permission
	return dto.DescribeKongServiceResponse{
		Service: ks.ToForDetail(),
	}
}

func (s *kongService) UpdateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.UpdateKongServiceRequest) dto.UpdateKongServiceResponse {
	s.validator.ValidateUpdateKongService(req)

	// TODO: permission
	// TODO: transaction
	ks, err := s.kongServiceRepo.DescribeService(ctx, entity.KongServiceId(req.Id))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(ks == nil, base.NotFoundKongService.WithRawError(fmt.Errorf("kong service not found: %d", req.Id)))

	s.mapper.UpdateKongServiceByUpdateReq(ks, req)

	err = s.kongServiceRepo.ReplaceService(ctx, *ks)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	return dto.UpdateKongServiceResponse{}
}

func (s *kongService) DeleteKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DeleteKongServiceRequest) {
	s.validator.ValidateDeleteKongService(req)

	// TODO: permission
	err := s.kongServiceRepo.DeleteService(ctx, entity.KongServiceId(req.Id))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
}
