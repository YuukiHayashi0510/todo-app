package staff

import (
	"time"

	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
)

type Staff struct {
	StaffID        int64
	OrganizationID int64
	Email          string
	StaffName      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Organization   organization.Organization
}
