package organization

import (
	"context"
	"fmt"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type IOrganizationService interface {
	AddAccountToOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.JoinOrganizationRequest)
	RemoveAccountFromOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.QuitOrganizationRequest)
	CreateOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.CreateOrganizationRequest) dto.CreateOrganizationResponse
	DescribeOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeOrganizationRequest) dto.DescribeOrganizationResponse
}

type organizationService struct {
	accountRepo repo.IAccountRepo
	orgRepo     repo.IOrganizationRepo
}

func NewOrganizationService(accountRepo repo.IAccountRepo, orgRepo repo.IOrganizationRepo) IOrganizationService {
	s := &organizationService{
		accountRepo: accountRepo,
		orgRepo:     orgRepo,
	}
	return s
}

func (s *organizationService) assertAccountCanOperateOrganization(ctx context.Context, accountId entity.AccountId, orgId entity.OrganizationId) {
	isMember, err := s.accountRepo.IsAccountInOrganization(ctx, accountId, orgId)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(!isMember, base.PermissionDenied.WithRawError(fmt.Errorf("account not in organization, account id: %d, organization id: %d", accountId, orgId)))
}

func (s *organizationService) AddAccountToOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.JoinOrganizationRequest) {

	currentAccountId := entity.AccountId(kongArgs.AccountId)
	targetAccountId := entity.AccountId(req.AccountId)

	// verify whether the current account is the member of the organization
	s.assertAccountCanOperateOrganization(ctx, currentAccountId, entity.OrganizationId(req.OrganizationId))

	// verify whether the current account has joined the organization
	joined, err := s.accountRepo.IsAccountInOrganization(ctx, targetAccountId, entity.OrganizationId(req.OrganizationId))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(joined, base.HasJoinedOrganization.WithRawError(fmt.Errorf("account has joined the organization, id: %d, organization: %d", req.AccountId, req.OrganizationId)))

	err = s.orgRepo.AddAccountToOrganization(ctx, targetAccountId, entity.OrganizationId(req.OrganizationId))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
}

func (s *organizationService) RemoveAccountFromOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.QuitOrganizationRequest) {
	currentAccountId := entity.AccountId(kongArgs.AccountId)
	targetAccountId := entity.AccountId(req.AccountId)

	// verify whether the current account is the member of the organization
	s.assertAccountCanOperateOrganization(ctx, currentAccountId, entity.OrganizationId(req.OrganizationId))

	err := s.orgRepo.RemoveAccountFromOrganization(ctx, targetAccountId, entity.OrganizationId(req.OrganizationId))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
}

func (s *organizationService) CreateOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.CreateOrganizationRequest) dto.CreateOrganizationResponse {
	org := entity.NewOrganization(req.Name)
	err := s.orgRepo.CreateOrganization(ctx, &org)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	err = s.orgRepo.AddAccountToOrganization(ctx, entity.AccountId(kongArgs.AccountId), org.ID)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	return dto.CreateOrganizationResponse{
		Organization: org.ToForDetail(make([]entity.Account, 0)),
	}
}

func (s *organizationService) DescribeOrganization(ctx context.Context, kongArgs base.KongArgs, req dto.DescribeOrganizationRequest) dto.DescribeOrganizationResponse {
	// verify whether the current account is the member of the organization
	s.assertAccountCanOperateOrganization(ctx, entity.AccountId(kongArgs.AccountId), entity.OrganizationId(req.Id))

	organization, err := s.orgRepo.DescribeOrganizationById(ctx, entity.OrganizationId(req.Id))
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	// get the members of the organization
	memberIds := s.orgRepo.ListOrganizationMembers(ctx, entity.OrganizationId(req.Id))
	members, err := s.accountRepo.ListAccounts(ctx, repo.ListAccountsRequest{
		Ids: memberIds,
	})
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	return dto.DescribeOrganizationResponse{
		Organization: organization.ToForDetail(members),
	}
}
