package entity

import (
	"time"

	"github.com/rs/xid"
	"github.com/samber/lo"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type KongServiceId uint
type KongServiceVersionId uint

type KongService struct {
	ID             KongServiceId         `gorm:"primaryKey"`
	OrganizationId OrganizationId        `gorm:"index:OrganizationId"` // belong organization
	CreatorId      AccountId             `gorm:"index:CreatorId"`      // created by who
	VersionId      *KongServiceVersionId // current version, nil means no version
	Name           string                `gorm:"index:Name"`
	Description    string                `gorm:"index:Description"`
	CreatedAt      time.Time
	UpdatedAt      time.Time `gorm:"index:UpdatedAt"`
}

type KongServiceVersion struct {
	Id            KongServiceVersionId `gorm:"primaryKey"`
	KongServiceId KongServiceId        `gorm:"index:kong_service_id"` // belong service
	Hash          string               // the unique hash for version
	Changelog     string               // changelog message
	CreatorId     AccountId            // the creator of this version
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s KongService) TableName() string {
	return "kong_service"
}

func (s KongServiceVersion) TableName() string {
	return "kong_service_version"
}

func NewKongService(creatorId AccountId, organizationId OrganizationId, name, description string) KongService {
	now := time.Now()
	return KongService{
		ID:             0,
		OrganizationId: organizationId,
		CreatorId:      creatorId,
		VersionId:      nil,
		Name:           name,
		Description:    description,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func NewKongServiceVersion(kongServiceId KongServiceId, changelog string, creatorId AccountId) KongServiceVersion {
	now := time.Now()
	return KongServiceVersion{
		Id:            0,
		KongServiceId: kongServiceId,
		Hash:          xid.New().String(),
		Changelog:     changelog,
		CreatorId:     creatorId,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func (s KongService) ToForList(versionAmount int) dto.KongServiceForList {
	return dto.KongServiceForList{
		Id:               uint(s.ID),
		Name:             s.Name,
		Description:      s.Description,
		VersionAmount:    versionAmount,
		CurrentVersionId: (*uint)(s.VersionId),
		CreatedAt:        s.CreatedAt.Unix(),
		UpdatedAt:        s.UpdatedAt.Unix(),
	}
}

func (s KongService) ToForDetail(creator Account, versions []KongServiceVersion) dto.KongServiceForDetail {
	return dto.KongServiceForDetail{
		Id:               uint(s.ID),
		Name:             s.Name,
		Description:      s.Description,
		CurrentVersionId: (*uint)(s.VersionId),
		Versions: lo.Map(versions, func(version KongServiceVersion, index int) dto.KongServiceVersionForDetail {
			return version.ToForDetail(creator)
		}),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func (s KongServiceVersion) ToForDetail(creator Account) dto.KongServiceVersionForDetail {
	return dto.KongServiceVersionForDetail{
		Id:            uint(s.Id),
		KongServiceId: uint(s.KongServiceId),
		Creator:       creator.ToForKongServiceVersion(),
		Changelog:     s.Changelog,
		CreatedAt:     s.CreatedAt.Unix(),
		UpdatedAt:     s.UpdatedAt.Unix(),
	}
}
