package session

import (
	"context"
	"testing"

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

var exampleSecret = "kong_test"

func TestLogin(t *testing.T) {
	suite := newTestSuite(t)

	svc := NewSessionService(suite.mockAccountRepo, exampleSecret)

	ctx := context.Background()

	cases := []struct {
		name            string
		email           string
		password        string
		correctPassword string
		except          *base.Exception
	}{
		{
			name:            "login successfully",
			email:           "one@example.com",
			password:        "123456",
			correctPassword: "123456",
		},
		{
			name:            "password not match",
			email:           "one@example.com",
			password:        "12345678",
			correctPassword: "123456",
			except:          &base.Unauthorized,
		},
	}
	for _, one := range cases {
		hashedPassword, err := hashPassword(one.correctPassword)
		require.NoError(t, err)
		suite.mockAccountRepo.EXPECT().DescribeAccountByEmail(gomock.Any(), gomock.Any()).Return(&entity.Account{
			ID:       1,
			Email:    one.email,
			Password: hashedPassword,
		}, nil)

		except := base.Catch(func() {
			resp := svc.Login(ctx, dto.LoginRequest{
				Email:    one.email,
				Password: one.password,
			})
			require.Equal(t, one.email, resp.Account.Email)
			require.NotEmpty(t, resp.Token)
		})
		if one.except == nil {
			require.Nil(t, except)
		} else {
			require.NotNil(t, except)
			require.Equal(t, one.except.Message, except.Message)
		}
	}
}

func TestRegister(t *testing.T) {
	suite := newTestSuite(t)

	svc := NewSessionService(suite.mockAccountRepo, exampleSecret)

	ctx := context.Background()
	email := "kong@example.com"
	suite.mockAccountRepo.EXPECT().DescribeAccountByEmail(gomock.Any(), gomock.Any()).Return(nil, nil)
	suite.mockAccountRepo.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, account *entity.Account) {
		account.ID = 1
		account.Email = email
	})
	resp := svc.Register(ctx, dto.RegisterRequest{
		Email:    email,
		NickName: "kong_test",
		Password: "123456789",
	})
	require.Equal(t, email, resp.Account.Email)
	require.NotEmpty(t, resp.Token)

}
