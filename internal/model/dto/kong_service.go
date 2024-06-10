package dto

type KongServiceForDetail struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

type CreateKongServiceRequest struct {
	Name        string `json:",required"`
	Description string `json:",required"`
}

type CreateKongServiceResponse struct {
	Service KongServiceForDetail
}

type ListKongServicesRequest struct {
	Pagination *PagingOption
	SortBy     []string
}

type ListKongServicesResponse struct {
	Services []KongServiceForDetail
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
