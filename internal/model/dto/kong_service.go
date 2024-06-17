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
	Name             string                        // the name of the service
	Description      string                        // what is the service about
	CurrentVersionId *uint                         // the id of the current version, optional
	Versions         []KongServiceVersionForDetail // the versions of the service
	CreatedAt        int64
	UpdatedAt        int64
}

type CreateKongServiceRequest struct {
	Name           string `json:",required"` // the name of the service
	Description    string `json:",required"` // what is the service about
	OrganizationId uint   `json:",required"` // the organization the service wants to join
}

type CreateKongServiceResponse struct {
	Service KongServiceForDetail
}

type ListKongServicesRequest struct {
	Name        *string       // fuzzy search by service name
	Description *string       // fuzzy search by service description
	Fuzzy       *string       // fuzzy search by both service name and description
	Pagination  *PagingOption // the pagination option, default, page size: 10, page number: 1
	SortBy      []string      // the sort option, default, sort by id. Supporting fields: ID, UpdatedAt, CreatedAt, Name. Add "-" prefix for descending order, Example: ["-CreatedAt", "Name"]
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
