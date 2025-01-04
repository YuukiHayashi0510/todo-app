package staff

import (
	"time"

	"github.com/YuukiHayashi0510/todo-app/internal/app/common"
	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
)

type Staff struct {
	common.BaseModel
	StaffID        int64
	OrganizationID int64
	Email          string
	StaffName      string
	Organization   organization.Organization
}

func New(
	staffID, organizationID int64, email, staffName string,
	createdAt, updatedAt time.Time, deletedAt *time.Time,
	organization organization.Organization,
) *Staff {
	return &Staff{
		StaffID:        staffID,
		OrganizationID: organizationID,
		Email:          email,
		StaffName:      staffName,
		BaseModel: common.BaseModel{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		},
		Organization: organization,
	}
}
