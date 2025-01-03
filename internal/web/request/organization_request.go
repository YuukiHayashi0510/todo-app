package request

type OrganizationRequest struct {
	OrganizationID   int64  `json:"organization_id" form:"organization_id" binding:"omitempty,number"`
	OrganizationName string `json:"organization_name" form:"organization_name" binding:"omitempty,min=1,max=256"`
}
