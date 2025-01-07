package organization

import (
	"time"

	"github.com/YuukiHayashi0510/todo-app/internal/domain/common"
)

type Organization struct {
	OrganizationID   int64  `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	common.BaseModel
}

func New(id int64, name string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Organization {
	return &Organization{
		OrganizationID:   id,
		OrganizationName: name,
		BaseModel: common.BaseModel{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		},
	}
}
