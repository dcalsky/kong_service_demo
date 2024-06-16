package kong_service

import (
	"context"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type testSuite struct {
	ctrl            *gomock.Controller
	mockAccountRepo *repo.MockIAccountRepo
	mockServiceRepo *repo.MockIKongServiceRepo
	mockOrgRepo     *repo.MockIOrganizationRepo
	mockRepoHelper  *repo.MockIRepoHelper
}

func newTestSuite(t *testing.T) testSuite {
	ctrl := gomock.NewController(t)
	return testSuite{
		ctrl:            ctrl,
		mockAccountRepo: repo.NewMockIAccountRepo(ctrl),
		mockServiceRepo: repo.NewMockIKongServiceRepo(ctrl),
		mockOrgRepo:     repo.NewMockIOrganizationRepo(ctrl),
		mockRepoHelper:  repo.NewMockIRepoHelper(ctrl),
	}
}

func TestCreateKongService(t *testing.T) {
	suite := newTestSuite(t)
	svc := NewKongService(suite.mockServiceRepo, suite.mockAccountRepo, suite.mockRepoHelper)
	kongArgs := base.KongArgs{
		AccountId: 10,
	}
	cases := []struct {
		name            string
		orgId           entity.OrganizationId
		accountId       entity.AccountId
		kongServiceName string
	}{
		{
			name:            "create kong service successfully",
			accountId:       10,
			orgId:           1,
			kongServiceName: "test",
		},
	}
	for _, one := range cases {
		t.Run(one.name, func(t *testing.T) {
			ctx := context.Background()
			suite.mockAccountRepo.EXPECT().DescribeAccountById(gomock.Any(), gomock.Any()).Return(&entity.Account{
				ID: one.accountId,
			}, nil)
			suite.mockAccountRepo.EXPECT().IsAccountInOrganization(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			suite.mockRepoHelper.EXPECT().MustTransaction(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, cb func(ctx2 context.Context)) {
				cb(ctx)
			})
			suite.mockServiceRepo.EXPECT().CreateService(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, service *entity.KongService) {
				service.ID = entity.KongServiceId(1)
				service.Name = one.kongServiceName
			})
			resp := svc.CreateKongService(ctx, kongArgs, dto.CreateKongServiceRequest{
				Name:           one.kongServiceName,
				Description:    "example",
				OrganizationId: uint(one.orgId),
			})
			require.Equal(t, one.kongServiceName, resp.Service.Name)
			require.NotEmpty(t, resp.Service.Id)
		})
	}
}

func TestUpdateKongService(t *testing.T) {
	suite := newTestSuite(t)
	svc := NewKongService(suite.mockServiceRepo, suite.mockAccountRepo, suite.mockRepoHelper)
	kongArgs := base.KongArgs{
		AccountId: 10,
	}
	cases := []struct {
		name                   string
		orgId                  entity.OrganizationId
		accountId              entity.AccountId
		kongServiceName        string
		updatedKongServiceName string
	}{
		{
			name:                   "update kong service successfully",
			accountId:              10,
			orgId:                  1,
			kongServiceName:        "test",
			updatedKongServiceName: "updated_test",
		},
	}
	for _, one := range cases {
		t.Run(one.name, func(t *testing.T) {
			ctx := context.Background()
			suite.mockServiceRepo.EXPECT().DescribeService(gomock.Any(), gomock.Any()).Return(&entity.KongService{
				ID:             1,
				OrganizationId: one.orgId,
				CreatorId:      one.accountId,
				Name:           one.kongServiceName,
			}, nil)
			suite.mockAccountRepo.EXPECT().DescribeAccountById(gomock.Any(), gomock.Any()).Return(&entity.Account{
				ID: one.accountId,
			}, nil)
			suite.mockServiceRepo.EXPECT().ListVersionsByServiceId(gomock.Any(), gomock.Any()).Return(nil, nil)
			suite.mockAccountRepo.EXPECT().IsAccountInOrganization(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			suite.mockRepoHelper.EXPECT().MustTransaction(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, cb func(ctx2 context.Context)) {
				cb(ctx)
			})
			suite.mockServiceRepo.EXPECT().ReplaceService(gomock.Any(), gomock.Any()).Return(nil)
			resp := svc.UpdateKongService(ctx, kongArgs, dto.UpdateKongServiceRequest{
				Id:   1,
				Name: lo.ToPtr(one.updatedKongServiceName),
			})
			require.Equal(t, one.updatedKongServiceName, resp.Service.Name)
		})
	}
}

func TestListKongServices(t *testing.T) {
	suite := newTestSuite(t)
	svc := NewKongService(suite.mockServiceRepo, suite.mockAccountRepo, suite.mockRepoHelper)
	kongArgs := base.KongArgs{
		AccountId: 10,
	}
	cases := []struct {
		name            string
		orgId           entity.OrganizationId
		accountId       entity.AccountId
		kongServiceName string
		versionAmount   int
	}{
		{
			name:            "list kong services successfully",
			accountId:       10,
			orgId:           1,
			kongServiceName: "test",
			versionAmount:   10,
		},
	}
	for _, one := range cases {
		t.Run(one.name, func(t *testing.T) {
			ctx := context.Background()
			suite.mockServiceRepo.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return([]entity.KongService{
				{
					ID:             1,
					OrganizationId: one.orgId,
					CreatorId:      one.accountId,
					Name:           one.kongServiceName,
				},
			}, dto.PagingResult{Total: 1}, nil)
			suite.mockServiceRepo.EXPECT().CountServicesVersionAmount(gomock.Any(), gomock.Any()).Return(map[entity.KongServiceId]int{
				1: one.versionAmount,
			}, nil)
			resp := svc.ListKongServices(ctx, kongArgs, dto.ListKongServicesRequest{})
			require.Len(t, resp.Services, 1)
			require.Equal(t, one.versionAmount, resp.Services[0].VersionAmount)
			require.Equal(t, int64(len(resp.Services)), resp.Pagination.Total)
		})
	}
}

func TestCreateKongServiceVersion(t *testing.T) {
	suite := newTestSuite(t)
	svc := NewKongService(suite.mockServiceRepo, suite.mockAccountRepo, suite.mockRepoHelper)
	kongArgs := base.KongArgs{
		AccountId: 10,
	}
	cases := []struct {
		name      string
		orgId     entity.OrganizationId
		accountId entity.AccountId
		changelog string
		versions  []entity.KongServiceVersion
	}{
		{
			name:      "auto switch",
			accountId: 10,
			orgId:     1,
			versions:  nil,
			changelog: strings.Repeat("ok", 10),
		},
		{
			name:      "no switch",
			accountId: 10,
			orgId:     1,
			changelog: strings.Repeat("ok1", 10),
			versions: []entity.KongServiceVersion{
				{
					KongServiceId: 1,
				},
			},
		},
	}
	for _, one := range cases {
		t.Run(one.name, func(t *testing.T) {
			ctx := context.Background()
			suite.mockServiceRepo.EXPECT().DescribeService(gomock.Any(), gomock.Any()).Return(&entity.KongService{
				ID:             1,
				OrganizationId: one.orgId,
				CreatorId:      one.accountId,
			}, nil)
			suite.mockAccountRepo.EXPECT().DescribeAccountById(gomock.Any(), gomock.Any()).Return(&entity.Account{
				ID: one.accountId,
			}, nil)
			suite.mockAccountRepo.EXPECT().IsAccountInOrganization(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			suite.mockServiceRepo.EXPECT().ListVersionsByServiceId(gomock.Any(), gomock.Any()).Return(one.versions, nil)
			suite.mockRepoHelper.EXPECT().MustTransaction(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, cb func(ctx2 context.Context)) {
				cb(ctx)
			})
			if len(one.versions) == 0 {
				suite.mockServiceRepo.EXPECT().ReplaceService(gomock.Any(), gomock.Any())
			}
			suite.mockServiceRepo.EXPECT().CreateServiceVersion(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, version *entity.KongServiceVersion) {
				version.Id = 1
				version.Changelog = one.changelog
			}).Return(nil)

			resp := svc.CreateKongServiceVersion(ctx, kongArgs, dto.CreateKongServiceVersionRequest{
				KongServiceId: 1,
				Changelog:     one.changelog,
			})
			require.Equal(t, uint(1), resp.Version.Id)
			require.Equal(t, one.changelog, resp.Version.Changelog)
		})
	}
}

func TestSwitchKongServiceVersion(t *testing.T) {
	suite := newTestSuite(t)
	svc := NewKongService(suite.mockServiceRepo, suite.mockAccountRepo, suite.mockRepoHelper)
	kongArgs := base.KongArgs{
		AccountId: 10,
	}
	cases := []struct {
		name          string
		orgId         entity.OrganizationId
		accountId     entity.AccountId
		versionId     entity.KongServiceVersionId
		kongServiceId entity.KongServiceId
		except        *base.Exception
	}{
		{
			name:          "switch successfully",
			accountId:     10,
			orgId:         1,
			kongServiceId: 1,
			versionId:     1,
		},
		{
			name:          "switch failed",
			accountId:     10,
			orgId:         1,
			kongServiceId: 2,
			versionId:     2,
			except:        &base.ForbidSwitchVersionNotBelongToService,
		},
	}
	for _, one := range cases {
		t.Run(one.name, func(t *testing.T) {
			ctx := context.Background()
			suite.mockAccountRepo.EXPECT().IsAccountInOrganization(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			suite.mockServiceRepo.EXPECT().DescribeService(gomock.Any(), gomock.Any()).Return(&entity.KongService{
				ID:             1,
				OrganizationId: one.orgId,
				CreatorId:      one.accountId,
			}, nil)
			suite.mockServiceRepo.EXPECT().DescribeServiceVersion(gomock.Any(), gomock.Any()).Return(&entity.KongServiceVersion{
				Id:            one.versionId,
				KongServiceId: one.kongServiceId,
			}, nil)

			if one.kongServiceId == entity.KongServiceId(1) {
				suite.mockServiceRepo.EXPECT().ReplaceService(gomock.Any(), gomock.Any())
			}
			except := base.Catch(func() {
				svc.SwitchKongServiceVersion(ctx, kongArgs, dto.SwitchKongServiceVersionRequest{
					KongServiceId: 1,
					VersionId:     uint(one.versionId),
				})
			})
			if one.except != nil {
				require.NotNil(t, except)
				require.Equal(t, one.except.Message, except.Message)
			}
		})
	}
}
