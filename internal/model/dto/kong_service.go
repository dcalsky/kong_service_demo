package dto

type CreatorInKongServiceVersion struct {
	Id       uint
	Email    string
	NickName string
}

type KongServiceVersionForDetail struct {
	Id            uint
	KongServiceId uint
	Creator       CreatorInKongServiceVersion
	Changelog     string
	CreatedAt     int64
	UpdatedAt     int64
}

type KongServiceForList struct {
	Id               uint
	Name             string // the name of the service
	Description      string // what is the service about
	VersionAmount    int    // how many versions does this service have
	CurrentVersionId *uint  // the id of the current version
	CreatedAt        int64
	UpdatedAt        int64
}

type KongServiceForDetail struct {
	Id               uint
	Name             string
	Description      string
	CurrentVersionId *uint
	Versions         []KongServiceVersionForDetail
	CreatedAt        int64
	UpdatedAt        int64
}

type CreateKongServiceRequest struct {
	Name           string `json:",required"`
	Description    string `json:",required"`
	OrganizationId uint   `json:",required"`
}

type CreateKongServiceResponse struct {
	Service KongServiceForDetail
}

type ListKongServicesRequest struct {
	Name        *string
	Description *string
	Fuzzy       *string
	Pagination  *PagingOption
	SortBy      []string
}

type ListKongServicesResponse struct {
	Services   []KongServiceForList
	Pagination PagingResult
}

type UpdateKongServiceRequest struct {
	Id          uint `json:",required"`
	Name        *string
	Description *string
}

type UpdateKongServiceResponse struct {
	Service KongServiceForDetail
}

type DeleteKongServiceRequest struct {
	Id uint `json:",required"`
}

type DescribeKongServiceRequest struct {
	Id uint `json:",required"`
}

type DescribeKongServiceResponse struct {
	Service KongServiceForDetail
}

type CreateKongServiceVersionRequest struct {
	KongServiceId      uint   `json:",required"`
	Changelog          string `json:",required"`
	SwitchToNewVersion *bool
}

type CreateKongServiceVersionResponse struct {
	Version KongServiceVersionForDetail
}

type SwitchKongServiceVersionRequest struct {
	KongServiceId uint `json:",required"`
	VersionId     uint `json:",required"`
}

type SwitchKongServiceVersionResponse struct {
	KongService KongServiceForDetail
}
