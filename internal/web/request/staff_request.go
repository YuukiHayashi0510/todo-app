package request

type StaffRequest struct {
	OrganizationID int64  `json:"organization_id" form:"organization_id" binding:"omitempty,number"`
	StaffID        int64  `json:"staff_id" form:"staff_id" binding:"omitempty,number"`
	StaffName      string `json:"staff_name" form:"staff_name" binding:"omitempty,min=1,max=256"`
	Email          string `json:"email" form:"email" binding:"omitempty,min=1,max=256"`
}
