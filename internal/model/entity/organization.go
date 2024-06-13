package entity

import (
	"time"

	"github.com/samber/lo"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type OrganizationId uint
type OrganizationAccountMappingId uint

type Organization struct {
	ID        OrganizationId `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationAccountMapping struct {
	ID             OrganizationAccountMappingId `gorm:"primaryKey"`
	AccountId      AccountId                    `gorm:"index:AccountId_OrganizationId"`
	OrganizationId OrganizationId               `gorm:"index:AccountId_OrganizationId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s Organization) TableName() string {
	return "organization"
}

func (s OrganizationAccountMapping) TableName() string {
	return "organization_account_mapping"
}

func NewOrganization(name string) Organization {
	now := time.Now()
	return Organization{
		ID:        0,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewOrganizationAccountMapping(accountId AccountId, organizationId OrganizationId) OrganizationAccountMapping {
	now := time.Now()
	return OrganizationAccountMapping{
		ID:             0,
		AccountId:      accountId,
		OrganizationId: organizationId,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func (s Organization) ToForDetail(members []Account) dto.OrganizationForDetail {
	return dto.OrganizationForDetail{
		Id:   uint(s.ID),
		Name: s.Name,
		Members: lo.Map(members, func(member Account, index int) dto.AccountForDetail {
			return member.ToForDetail()
		}),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}
