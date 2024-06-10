package repo

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type IKongServiceRepo interface {
	ListServices(ctx context.Context, req ListServicesRequest) ([]entity.KongService, error)
	DescribeService(ctx context.Context, id entity.KongServiceId) (*entity.KongService, error)
	CreateService(ctx context.Context, service *entity.KongService) error
	ReplaceService(ctx context.Context, service entity.KongService) error
	DeleteService(ctx context.Context, id entity.KongServiceId) error
}

type kongServiceRepo struct {
	table *gorm.DB
}

type ListServicesRequest struct {
	CreatorAccountId *string
	Name             *string
	Description      *string
	Pagination       *dto.PagingOption
	SortBy           []string
	All              bool
}

func NewKongServiceRepo(table *gorm.DB) IKongServiceRepo {
	s := &kongServiceRepo{
		table: table,
	}
	return s
}

func (s *kongServiceRepo) ListServices(ctx context.Context, req ListServicesRequest) ([]entity.KongService, error) {
	tx := s.table.WithContext(ctx)
	if req.CreatorAccountId != nil {
		tx.Where("CreatorAccountId = ?", *req.CreatorAccountId)
	}
	if req.Name != nil {
		tx.Where("Name like ?", *req.Name)
	}
	if req.Description != nil {
		tx.Where("Description like ?", *req.Description)
	}
	BuildGormSortBy(tx, req.SortBy)
	BuildGormPagination(tx, req.Pagination)
	var res []entity.KongService
	err := tx.Find(&res).Error
	return res, err
}

func (s *kongServiceRepo) DescribeService(ctx context.Context, id entity.KongServiceId) (*entity.KongService, error) {
	var res entity.KongService
	err := s.table.WithContext(ctx).Where("Id = ?", id).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func (s *kongServiceRepo) CreateService(ctx context.Context, service *entity.KongService) error {
	if service == nil {
		return errors.New("kong service is nil")
	}
	return s.table.WithContext(ctx).Create(service).Error
}

func (s *kongServiceRepo) ReplaceService(ctx context.Context, service entity.KongService) error {
	return s.table.WithContext(ctx).Save(&service).Error
}

func (s *kongServiceRepo) DeleteService(ctx context.Context, id entity.KongServiceId) error {
	return s.table.WithContext(ctx).Where("Id = ?", id).Delete(&entity.KongService{}).Error
}
