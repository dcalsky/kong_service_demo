package repo

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

//go:generate mockgen -package repo --build_flags=--mod=mod  --destination account_mock.go . IAccountRepo
type IAccountRepo interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	DescribeAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
	DescribeAccountById(ctx context.Context, id entity.AccountId) (*entity.Account, error)
	IsAccountInOrganization(ctx context.Context, accountId entity.AccountId, organizationId entity.OrganizationId) (bool, error)
	ListAccounts(ctx context.Context, req ListAccountsRequest) ([]entity.Account, error)
}

type ListAccountsRequest struct {
	Ids []entity.AccountId
}

type accountRepo struct {
	db IRepoHelper
}

func NewAccountRepo(db IRepoHelper) IAccountRepo {
	s := &accountRepo{
		db: db,
	}
	return s
}

func (s *accountRepo) CreateAccount(ctx context.Context, account *entity.Account) error {
	return s.db.WithContext(ctx).Create(account).Error
}

func (s *accountRepo) DescribeAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	err := s.db.WithContext(ctx).Where("Email = ?", email).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (s *accountRepo) DescribeAccountById(ctx context.Context, id entity.AccountId) (*entity.Account, error) {
	var account entity.Account
	err := s.db.WithContext(ctx).Where("ID = ?", id).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (s *accountRepo) IsAccountInOrganization(ctx context.Context, accountId entity.AccountId, organizationId entity.OrganizationId) (bool, error) {
	var res int64
	err := s.db.WithContext(ctx).Model(&entity.OrganizationAccountMapping{}).Where("AccountId = ? and OrganizationId = ?", accountId, organizationId).Count(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return res >= 1, nil
}

func (s *accountRepo) ListAccounts(ctx context.Context, req ListAccountsRequest) ([]entity.Account, error) {
	var accounts []entity.Account
	err := s.db.WithContext(ctx).Where("ID in ?", req.Ids).Find(&accounts).Error
	return accounts, err
}
