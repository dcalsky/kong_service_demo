package entity

import (
	"time"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type KongServiceId uint

type KongService struct {
	Id        KongServiceId `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s KongService) TableName() string {
	return "kong_service"
}

func NewKongService() KongService {
	now := time.Now()
	return KongService{
		Id:        0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s KongService) ToForDetail() dto.KongServiceForDetail {
	return dto.KongServiceForDetail{
		Id:        uint(s.Id),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}
