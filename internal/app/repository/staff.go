package repository

import (
	"context"

	"github.com/YuukiHayashi0510/todo-app/internal/app/common"
	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
	"github.com/YuukiHayashi0510/todo-app/internal/app/staff"
	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/postgres"
	"github.com/YuukiHayashi0510/todo-app/internal/persistence/rdb"
)

type ServiceStaff = staff.Staff

type StaffRepository struct {
	db *postgres.DB
}

func NewStaffRepository(db *postgres.DB) *StaffRepository {
	return &StaffRepository{db}
}

func (r *StaffRepository) Count(ctx context.Context, input *staff.SearchInput) (int64, error) {
	queries := rdb.New(r.db)

	params := rdb.CountSearchStaffsParams{
		StaffName:    input.StaffName,
		SearchStatus: string(input.SearchStatus),
	}

	return queries.CountSearchStaffs(ctx, params)
}

func (r StaffRepository) FindByID(ctx context.Context, id int64) (*ServiceStaff, error) {
	queries := rdb.New(r.db)

	value, err := queries.GetStaffByID(ctx, id)
	if err != nil {
		return nil, err
	}

	staff := staff.New(
		value.StaffID,
		value.OrganizationID,
		value.Email,
		value.StaffName,
		value.CreatedAt,
		value.UpdatedAt,
		value.DeletedAt,
		organization.Organization(value.Organization),
	)

	return staff, nil
}

func (r *StaffRepository) Search(ctx context.Context, input *staff.SearchInput) ([]ServiceStaff, error) {
	queries := rdb.New(r.db)

	params := rdb.SearchStaffsParams{
		StaffName:    input.StaffName,
		SearchStatus: string(input.SearchStatus),
		Limit:        int32(input.PerPage),
	}

	// offsetの計算
	if input.Page > 0 {
		params.Offset = int32((input.Page - 1) * input.PerPage)
	} else {
		params.Offset = 0
	}

	staffs, err := queries.SearchStaffs(ctx, params)
	if err != nil {
		return nil, err
	}

	retStaffs := make([]ServiceStaff, 0, len(staffs))
	for _, v := range staffs {
		retStaffs = append(retStaffs, *staff.New(
			v.StaffID,
			v.OrganizationID,
			v.Email,
			v.StaffName,
			v.CreatedAt,
			v.UpdatedAt,
			v.DeletedAt,
			organization.Organization(v.Organization),
		))
	}

	return retStaffs, nil
}

func (r *StaffRepository) Create(ctx context.Context, input *staff.CreateInput) (*ServiceStaff, error) {
	queries := rdb.New(r.db)

	v, err := queries.CreateStaff(ctx, rdb.CreateStaffParams{
		OrganizationID: input.OrganizationID,
		Email:          input.Email,
		StaffName:      input.StaffName,
	})
	if err != nil {
		return nil, err
	}

	return &ServiceStaff{
		StaffID:        v.StaffID,
		OrganizationID: v.OrganizationID,
		Email:          v.Email,
		StaffName:      v.StaffName,
		BaseModel: common.BaseModel{
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		},
	}, nil
}

func (r *StaffRepository) Update(ctx context.Context, input *staff.UpdateInput) error {
	queries := rdb.New(r.db)

	return queries.UpdateStaff(ctx, rdb.UpdateStaffParams{
		OrganizationID: input.OrganizationID,
		StaffID:        input.StaffID,
		Email:          input.Email,
		StaffName:      input.StaffName,
	})
}

func (r *StaffRepository) Delete(ctx context.Context, id int64) error {
	queries := rdb.New(r.db)

	return queries.SoftDeleteStaff(ctx, id)
}

func (r StaffRepository) Restore(ctx context.Context, id int64) error {
	queries := rdb.New(r.db)

	return queries.RestoreStaff(ctx, id)
}
