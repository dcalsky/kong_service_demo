package dto

type OrganizationForDetail struct {
	Id        uint
	Name      string
	Members   []AccountForDetail
	CreatedAt int64
	UpdatedAt int64
}

type JoinOrganizationRequest struct {
	AccountId      uint `json:",required"`
	OrganizationId uint `json:",required"`
}

type QuitOrganizationRequest struct {
	AccountId      uint `json:",required"`
	OrganizationId uint `json:",required"`
}

type CreateOrganizationRequest struct {
	Name string `json:",required"`
}

type CreateOrganizationResponse struct {
	Organization OrganizationForDetail
}

type DescribeOrganizationRequest struct {
	Id uint `json:",required"`
}

type DescribeOrganizationResponse struct {
	Organization OrganizationForDetail
}
