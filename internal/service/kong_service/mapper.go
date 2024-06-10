package kong_service

import (
	"time"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type mapper struct{}

func (s mapper) UpdateKongServiceByUpdateReq(ks *entity.KongService, req dto.UpdateKongServiceRequest) {
	// TODO: do more update operations
	ks.UpdatedAt = time.Now()
}
