package staff

import "github.com/YuukiHayashi0510/todo-app/internal/domain/common"

type SearchInput struct {
	Staff
	common.PaginationParams
	SearchStatus common.SearchStatus
}

type SearchOutput struct {
	Staffs   []Staff         `json:"staffs"`
	PageInfo common.PageInfo `json:"page_info"`
}

type CreateInput struct {
	OrganizationID int64
	Email          string
	StaffName      string
}

type CreateOutput struct {
	Staff Staff `json:"staff"`
}

type UpdateInput struct {
	OrganizationID int64
	StaffID        int64
	Email          string
	StaffName      string
}

type UpdateOutput struct {
	Staff Staff `json:"staff"`
}
