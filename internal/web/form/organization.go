package form

// TODO: bindingタグ
type OrganizationForm struct {
	OrganizationID   int64  `json:"organization_id" form:"organization_id" binding:"omitempty"`
	OrganizationName string `json:"organization_name" form:"organization_name" binding:"omitempty"`
}
