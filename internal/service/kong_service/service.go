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

//go:generate mockgen -package kong_service --build_flags=--mod=mod  --destination service_mock.go . IKongService
type IKongService interface {
	CreateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceRequest) dto.CreateKongServiceResponse
	ListKongServices(ctx context.Context, kongArgs base.KongArgs, req dto.ListKongServicesRequest) dto.ListKongServicesResponse
	DescribeKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeKongServiceRequest) dto.DescribeKongServiceResponse
	UpdateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.UpdateKongServiceRequest) dto.UpdateKongServiceResponse
	DeleteKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DeleteKongServiceRequest)
	CreateKongServiceVersion(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceVersionRequest) dto.CreateKongServiceVersionResponse
	SwitchKongServiceVersion(ctx context.Context, kongArgs base.KongArgs, req dto.SwitchKongServiceVersionRequest)
}

type kongService struct {
	validator
	mapper
	kongServiceRepo repo.IKongServiceRepo
	accountRepo     repo.IAccountRepo
	repoHelper      repo.IRepoHelper
}

func NewKongService(kongServiceRepo repo.IKongServiceRepo, accountRepo repo.IAccountRepo, repoHelper repo.IRepoHelper) IKongService {
	s := &kongService{
		validator:       validator{},
		mapper:          mapper{},
		kongServiceRepo: kongServiceRepo,
		accountRepo:     accountRepo,
		repoHelper:      repoHelper,
	}
	return s
}

func (s *kongService) mustSwitchServiceVersion(ctx context.Context, ks *entity.KongService, versionId entity.KongServiceVersionId) {
	s.mapper.UpdateKongServiceVersion(ks, versionId)
	err := s.kongServiceRepo.ReplaceService(ctx, *ks)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
}

func (s *kongService) assertAccountCanOperateService(ctx context.Context, kongArgs base.KongArgs, ks entity.KongService) {
	owned, err := s.accountRepo.IsAccountInOrganization(ctx, entity.AccountId(kongArgs.AccountId), ks.OrganizationId)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(!owned, base.NoPermissionToOperateKongService.WithRawError(fmt.Errorf("account id: %d, kong service id: %d, organization id: %d", kongArgs.AccountId, ks.ID, ks.OrganizationId)))
}

func (s *kongService) mustGetKongService(ctx context.Context, id entity.KongServiceId) entity.KongService {
	ks, err := s.kongServiceRepo.DescribeService(ctx, id)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(ks == nil, base.NotFoundKongService.WithRawError(fmt.Errorf("kong service not found: %d", id)))
	return *ks
}

func (s *kongService) CreateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceRequest) dto.CreateKongServiceResponse {
	s.validator.ValidateCreateKongService(req)

	var account *entity.Account
	var ks entity.KongService
	var err error

	s.repoHelper.MustTransaction(ctx, func(ctx context.Context) {
		// get the detail of current account
		account, err = s.accountRepo.DescribeAccountById(ctx, entity.AccountId(kongArgs.AccountId))
		base.PanicIfErr(err, base.InternalError.WithRawError(err))
		base.PanicIf(account == nil, base.Unauthorized.WithRawError(fmt.Errorf("account not found, id: %d", kongArgs.AccountId)))

		// verify whether the current account is the member of the organization
		accountIsOrgMember, err := s.accountRepo.IsAccountInOrganization(ctx, entity.AccountId(kongArgs.AccountId), entity.OrganizationId(req.OrganizationId))
		base.PanicIfErr(err, base.InternalError.WithRawError(err))
		base.PanicIf(!accountIsOrgMember, base.PermissionDenied.WithRawError(fmt.Errorf("account not in organization, account id: %d, organization id: %d", kongArgs.AccountId, req.OrganizationId)))

		ks = entity.NewKongService(entity.AccountId(kongArgs.AccountId), entity.OrganizationId(req.OrganizationId), req.Name, req.Description)
		err = s.kongServiceRepo.CreateService(ctx, &ks)
		base.PanicIfErr(err, base.InternalError.WithRawError(err))
	})

	return dto.CreateKongServiceResponse{
		Service: ks.ToForDetail(lo.FromPtr(account), make([]entity.KongServiceVersion, 0)),
	}
}

func (s *kongService) ListKongServices(ctx context.Context, kongArgs base.KongArgs, req dto.ListKongServicesRequest) dto.ListKongServicesResponse {
	services, pageRes, err := s.kongServiceRepo.ListServices(ctx, repo.ListServicesRequest{
		Pagination:  req.Pagination,
		SortBy:      req.SortBy,
		All:         false,
		Name:        req.Name,
		Description: req.Description,
		Fuzzy:       req.Fuzzy,
	})
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	serviceIds := lo.Map(services, func(svc entity.KongService, index int) entity.KongServiceId {
		return svc.ID
	})
	// count how many versions belong to each service
	serviceId2VersionAmount, err := s.kongServiceRepo.CountServicesVersionAmount(ctx, serviceIds)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	return dto.ListKongServicesResponse{
		Services: lo.Map(services, func(svc entity.KongService, index int) dto.KongServiceForList {
			return svc.ToForList(serviceId2VersionAmount[svc.ID])
		}),
		Pagination: pageRes,
	}
}

func (s *kongService) kongServiceToDetail(ctx context.Context, ks entity.KongService) dto.KongServiceForDetail {
	versions, err := s.kongServiceRepo.ListVersionsByServiceId(ctx, ks.ID)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	account, err := s.accountRepo.DescribeAccountById(ctx, ks.CreatorId)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	return ks.ToForDetail(lo.FromPtr(account), versions)
}

func (s *kongService) DescribeKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeKongServiceRequest) dto.DescribeKongServiceResponse {
	s.validator.ValidateDescribeKongService(req)

	ks := s.mustGetKongService(ctx, entity.KongServiceId(req.Id))
	s.assertAccountCanOperateService(ctx, kongArgs, ks)

	return dto.DescribeKongServiceResponse{
		Service: s.kongServiceToDetail(ctx, ks),
	}
}

func (s *kongService) UpdateKongService(ctx context.Context, kongArgs base.KongArgs, req dto.UpdateKongServiceRequest) dto.UpdateKongServiceResponse {
	s.validator.ValidateUpdateKongService(req)

	var ks entity.KongService
	s.repoHelper.MustTransaction(ctx, func(ctx context.Context) {
		ks = s.mustGetKongService(ctx, entity.KongServiceId(req.Id))
		s.assertAccountCanOperateService(ctx, kongArgs, ks)

		s.mapper.UpdateKongServiceByUpdateReq(&ks, req)

		err := s.kongServiceRepo.ReplaceService(ctx, ks)
		base.PanicIfErr(err, base.InternalError.WithRawError(err))
	})

	return dto.UpdateKongServiceResponse{
		Service: s.kongServiceToDetail(ctx, ks),
	}
}

func (s *kongService) DeleteKongService(ctx context.Context, kongArgs base.KongArgs, req dto.DeleteKongServiceRequest) {
	s.validator.ValidateDeleteKongService(req)

	ks := s.mustGetKongService(ctx, entity.KongServiceId(req.Id))
	s.assertAccountCanOperateService(ctx, kongArgs, ks)

	err := s.kongServiceRepo.DeleteService(ctx, entity.KongServiceId(req.Id))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
}

func (s *kongService) CreateKongServiceVersion(ctx context.Context, kongArgs base.KongArgs, req dto.CreateKongServiceVersionRequest) dto.CreateKongServiceVersionResponse {
	s.validator.ValidateCreateVersion(req)

	var newVersion entity.KongServiceVersion
	var ks entity.KongService

	s.repoHelper.MustTransaction(ctx, func(ctx context.Context) {
		ks = s.mustGetKongService(ctx, entity.KongServiceId(req.KongServiceId))
		s.assertAccountCanOperateService(ctx, kongArgs, ks)

		versions, err := s.kongServiceRepo.ListVersionsByServiceId(ctx, ks.ID)
		base.PanicIfErr(err, base.InternalError.WithRawError(err))

		if len(versions) == 0 {
			req.SwitchToNewVersion = lo.ToPtr(true)
		}

		newVersion = s.mapper.CreateKongServiceVersionByCreateReq(entity.AccountId(kongArgs.AccountId), req)
		err = s.kongServiceRepo.CreateServiceVersion(ctx, &newVersion)
		base.PanicIfErr(err, base.InternalError.WithRawError(err))

		// if there are only one versions, or if user want to make the new version the current version of service -> switch version
		if req.SwitchToNewVersion != nil && *req.SwitchToNewVersion {
			s.mustSwitchServiceVersion(ctx, &ks, newVersion.Id)
		}
	})

	account, err := s.accountRepo.DescribeAccountById(ctx, ks.CreatorId)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	return dto.CreateKongServiceVersionResponse{
		Version: newVersion.ToForDetail(lo.FromPtr(account)),
	}

}

func (s *kongService) SwitchKongServiceVersion(ctx context.Context, kongArgs base.KongArgs, req dto.SwitchKongServiceVersionRequest) {
	ks := s.mustGetKongService(ctx, entity.KongServiceId(req.KongServiceId))
	s.assertAccountCanOperateService(ctx, kongArgs, ks)

	serviceVersion, err := s.kongServiceRepo.DescribeServiceVersion(ctx, entity.KongServiceVersionId(req.VersionId))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(serviceVersion == nil, base.ResourceNotFound.WithRawError(fmt.Errorf("kong service version not found, id: %d", req.VersionId)))

	if ks.ID != serviceVersion.KongServiceId {
		panic(base.ForbidSwitchVersionNotBelongToService.WithRawError(fmt.Errorf("service id: %d, version id: %d", ks.ID, serviceVersion.Id)))
	}
	s.mustSwitchServiceVersion(ctx, &ks, serviceVersion.Id)
}
