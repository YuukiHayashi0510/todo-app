package common

const (
	defaultPage    = 1
	defaultPerPage = 100
)

type PageInfo struct {
	Page        int   `json:"page"`
	PerPage     int   `json:"per_page"`
	TotalCount  int64 `json:"total_count"`
	TotalPages  int   `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
}

func NewPageInfoWith(page, perPage int, totalCount int64) PageInfo {
	if page == 0 {
		page = defaultPage
	}
	if perPage == 0 {
		perPage = defaultPerPage
	}

	totalPages := (int(totalCount) + perPage - 1) / perPage

	return PageInfo{
		Page:        page,
		PerPage:     perPage,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}
}
