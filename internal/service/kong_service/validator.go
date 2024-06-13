package kong_service

import (
	"fmt"
	"unicode/utf8"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type validator struct {
}

func (s validator) ValidateCreateKongService(req dto.CreateKongServiceRequest) {
	s.validateKongServiceName(req.Name)
}

func (s validator) ValidateDescribeKongService(req dto.DescribeKongServiceRequest) {
	s.validateKongServiceId(req.Id)
}

func (s validator) ValidateUpdateKongService(req dto.UpdateKongServiceRequest) {
	s.validateKongServiceId(req.Id)
	if req.Name != nil {
		s.validateKongServiceName(*req.Name)
	}
}

func (s validator) ValidateDeleteKongService(req dto.DeleteKongServiceRequest) {
	s.validateKongServiceId(req.Id)
}

func (validator) validateKongServiceName(name string) {
	// let's assume the length of the name is less than 50
	base.PanicIf(len(name) > 50, base.ExceedMaximumKongServiceName.WithRawError(fmt.Errorf("name: %s", name)))
}

func (validator) validateKongServiceId(id uint) {
	base.PanicIf(id == 0, base.InvalidKongServiceId.WithRawError(fmt.Errorf("id: %d", id)))
}

func (validator) ValidateCreateVersion(req dto.CreateKongServiceVersionRequest) {
	changelogLen := utf8.RuneCountInString(req.Changelog)
	base.PanicIf(changelogLen < 6, base.InvalidChangelogLength.WithRawError(fmt.Errorf("changelog: %s", req.Changelog)))
}
