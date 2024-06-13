package kong_service

import (
	"time"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type mapper struct{}

func (s mapper) UpdateKongServiceByUpdateReq(ks *entity.KongService, req dto.UpdateKongServiceRequest) {
	ks.UpdatedAt = time.Now()
	if req.Name != nil {
		ks.Name = *req.Name
	}
	if req.Description != nil {
		ks.Description = *req.Description
	}
}

func (s mapper) CreateKongServiceVersionByCreateReq(creatorId entity.AccountId, req dto.CreateKongServiceVersionRequest) entity.KongServiceVersion {
	return entity.NewKongServiceVersion(entity.KongServiceId(req.KongServiceId), req.Changelog, creatorId)
}

func (s mapper) UpdateKongServiceVersion(ks *entity.KongService, versionId entity.KongServiceVersionId) {
	ks.VersionId = &versionId
	ks.UpdatedAt = time.Now()
}
