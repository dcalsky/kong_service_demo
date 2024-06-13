package repo

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type IOrganizationRepo interface {
	AddAccountToOrganization(ctx context.Context, accountId entity.AccountId, organization entity.OrganizationId) error
	RemoveAccountFromOrganization(ctx context.Context, accountId entity.AccountId, organization entity.OrganizationId) error
	CreateOrganization(ctx context.Context, organization *entity.Organization) error
	DescribeOrganizationById(ctx context.Context, organizationId entity.OrganizationId) (*entity.Organization, error)
	ListOrganizationMembers(ctx context.Context, organizationId entity.OrganizationId) []entity.AccountId
}

type organizationRepo struct {
	db *gorm.DB
}

func NewOrganizationRepo(db *gorm.DB) IOrganizationRepo {
	s := &organizationRepo{
		db: db,
	}
	return s
}

func (s *organizationRepo) AddAccountToOrganization(ctx context.Context, accountId entity.AccountId, organization entity.OrganizationId) error {
	now := time.Now()
	return s.db.WithContext(ctx).Create(&entity.OrganizationAccountMapping{
		ID:             0,
		AccountId:      accountId,
		OrganizationId: organization,
		CreatedAt:      now,
		UpdatedAt:      now,
	}).Error
}

func (s *organizationRepo) RemoveAccountFromOrganization(ctx context.Context, accountId entity.AccountId, organization entity.OrganizationId) error {
	return s.db.WithContext(ctx).Where("AccountId = ? and OrganizationId = ?", accountId, organization).Delete(&entity.OrganizationAccountMapping{}).Error
}

func (s *organizationRepo) CreateOrganization(ctx context.Context, organization *entity.Organization) error {
	return s.db.WithContext(ctx).Create(organization).Error
}

func (s *organizationRepo) DescribeOrganizationById(ctx context.Context, organizationId entity.OrganizationId) (*entity.Organization, error) {
	var organization entity.Organization
	err := s.db.WithContext(ctx).Where("ID = ?", organizationId).First(&organization).Error
	return &organization, err
}

func (s *organizationRepo) ListOrganizationMembers(ctx context.Context, organizationId entity.OrganizationId) []entity.AccountId {
	var accountIds []entity.AccountId
	s.db.Model(&entity.OrganizationAccountMapping{}).WithContext(ctx).Where("OrganizationId = ?", organizationId).Pluck("AccountId", &accountIds)
	return accountIds
}
