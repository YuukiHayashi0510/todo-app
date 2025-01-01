package organization

import "github.com/YuukiHayashi0510/todo-app/internal/app/common"

type SearchInput struct {
	Organization
	common.PaginationParams
	SearchStatus common.SearchStatus
}
