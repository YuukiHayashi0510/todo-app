package organization

import "github.com/YuukiHayashi0510/todo-app/internal/domain/common"

type SearchInput struct {
	Organization
	common.PaginationParams
	SearchStatus common.SearchStatus
}

type SearchOutput struct {
	Organizations []Organization  `json:"organizations"`
	PageInfo      common.PageInfo `json:"page_info"`
}

type CreateInput struct {
	OrganizationName string
}

type CreateOutput struct {
	Organization Organization `json:"organization"`
}

type UpdateInput struct {
	OrganizationID   int64
	OrganizationName string
}

type UpdateOutput struct {
	Organization Organization `json:"organization"`
}
