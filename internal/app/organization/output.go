package organization

import "github.com/YuukiHayashi0510/todo-app/internal/app/common"

type SearchOutput struct {
	Organizations []Organization  `json:"organizations"`
	PageInfo      common.PageInfo `json:"page_info"`
}

type CreateOutput struct {
	Organization Organization `json:"organization"`
}

type UpdateOutput struct {
	Organization Organization `json:"organization"`
}
