// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package rdb

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Feature struct {
	FeatureID   int64
	FeatureName string
	Operation   string
	Description pgtype.Text
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Organization struct {
	OrganizationID   int64
	OrganizationName string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

type Project struct {
	ProjectID      int64
	OrganizationID int64
	ProjectName    string
	Description    pgtype.Text
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type ProjectMember struct {
	ProjectMemberID int64
	ProjectID       int64
	StaffID         int64
	RoleID          int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

type Role struct {
	RoleID    int64
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RoleFeature struct {
	RoleID    int64
	FeatureID int64
	CreatedAt time.Time
}

type Staff struct {
	StaffID        int64
	OrganizationID int64
	Email          string
	StaffName      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type StaffSession struct {
	SessionID   int64
	StaffID     int64
	SessionData pgtype.Text
	ExpiresAt   time.Time
	CreatedAt   time.Time
}

type Task struct {
	TaskID      int64
	ProjectID   int64
	Title       string
	Description pgtype.Text
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
