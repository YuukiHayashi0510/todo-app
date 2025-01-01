package organization

import (
	"time"
)

type Organization struct {
	OrganizationID   int64      `json:"organization_id"`
	OrganizationName string     `json:"organization_name"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}
