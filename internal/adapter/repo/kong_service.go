package repo

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

//go:generate mockgen -package repo --build_flags=--mod=mod  --destination kong_service_mock.go . IKongServiceRepo
type IKongServiceRepo interface {
	ListServices(ctx context.Context, req ListServicesRequest) ([]entity.KongService, dto.PagingResult, error)
	DescribeService(ctx context.Context, id entity.KongServiceId) (*entity.KongService, error)
	CreateService(ctx context.Context, service *entity.KongService) error
	ReplaceService(ctx context.Context, service entity.KongService) error
	DeleteService(ctx context.Context, id entity.KongServiceId) error
	CountServicesVersionAmount(ctx context.Context, ids []entity.KongServiceId) (map[entity.KongServiceId]int, error)
	ListVersionsByServiceId(ctx context.Context, id entity.KongServiceId) ([]entity.KongServiceVersion, error)
	CreateServiceVersion(ctx context.Context, version *entity.KongServiceVersion) error
	DescribeServiceVersion(ctx context.Context, versionId entity.KongServiceVersionId) (*entity.KongServiceVersion, error)
}

type kongServiceRepo struct {
	db IRepoHelper
}

type ListServicesRequest struct {
	OrganizationIds []entity.OrganizationId
	CreatorId       *entity.AccountId
	Name            *string
	Description     *string
	Fuzzy           *string
	Pagination      *dto.PagingOption
	SortBy          []string
	All             bool
}

func NewKongServiceRepo(db IRepoHelper) IKongServiceRepo {
	s := &kongServiceRepo{
		db: db,
	}
	return s
}

func (s *kongServiceRepo) ListServices(ctx context.Context, req ListServicesRequest) ([]entity.KongService, dto.PagingResult, error) {
	tx := s.db.WithContext(ctx).Model(&entity.KongService{})
	if req.CreatorId != nil {
		tx.Where("CreatorId = ?", *req.CreatorId)
	}
	if len(req.OrganizationIds) > 0 {
		tx.Where("OrganizationId in ?", req.OrganizationIds)
	}
	if req.Fuzzy != nil {
		tx.Where("Name like ? or Description like ?", *req.Name, *req.Name)
	} else {
		if req.Name != nil {
			tx.Where("Name like ?", *req.Name)
		}
		if req.Description != nil {
			tx.Where("Description like ?", *req.Description)
		}
	}
	pageResult := BuildPageResult(tx, req.Pagination)
	BuildGormSortBy(tx, req.SortBy)
	BuildGormPagination(tx, req.Pagination)
	var res []entity.KongService
	err := tx.Find(&res).Error
	if err != nil {
		return nil, pageResult, err
	}
	return res, pageResult, nil
}

func (s *kongServiceRepo) DescribeService(ctx context.Context, id entity.KongServiceId) (*entity.KongService, error) {
	var res entity.KongService
	err := s.db.WithContext(ctx).Where("ID = ?", id).First(&res).Error
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
	return s.db.WithContext(ctx).Create(service).Error
}

func (s *kongServiceRepo) ReplaceService(ctx context.Context, service entity.KongService) error {
	return s.db.WithContext(ctx).Save(&service).Error
}

func (s *kongServiceRepo) DeleteService(ctx context.Context, id entity.KongServiceId) error {
	return s.db.WithContext(ctx).Where("ID = ?", id).Delete(&entity.KongService{}).Error
}

func (s *kongServiceRepo) CountServicesVersionAmount(ctx context.Context, ids []entity.KongServiceId) (map[entity.KongServiceId]int, error) {
	type countResultItem struct {
		KongServiceId entity.KongServiceId
		Amount        int
	}
	res := make(map[entity.KongServiceId]int)
	var countResult []countResultItem
	err := s.db.WithContext(ctx).Raw(`select KongServiceId, count(*) as Amount from kong_service_version where KongServiceId in (?) group by KongServiceId`, ids).Scan(&countResult).Error
	if err != nil {
		return res, err
	}
	for _, item := range countResult {
		res[item.KongServiceId] = item.Amount
	}
	return res, nil
}

func (s *kongServiceRepo) ListVersionsByServiceId(ctx context.Context, id entity.KongServiceId) ([]entity.KongServiceVersion, error) {
	var versions []entity.KongServiceVersion
	err := s.db.WithContext(ctx).Where("KongServiceId = ?", id).Order("ID desc").Find(&versions).Error
	return versions, err
}

func (s *kongServiceRepo) CreateServiceVersion(ctx context.Context, version *entity.KongServiceVersion) error {
	return s.db.WithContext(ctx).Create(version).Error
}

func (s *kongServiceRepo) DescribeServiceVersion(ctx context.Context, versionId entity.KongServiceVersionId) (*entity.KongServiceVersion, error) {
	var version entity.KongServiceVersion
	err := s.db.WithContext(ctx).Where("ID = ?", versionId).First(&version).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &version, err
}
