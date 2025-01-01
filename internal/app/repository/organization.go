package repository

import (
	"context"
	"database/sql"

	"github.com/YuukiHayashi0510/todo-app/internal/app/organization"
	"github.com/YuukiHayashi0510/todo-app/internal/persistence/rdb"
)

type OrganizationRepository struct {
	db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
	return &OrganizationRepository{db}
}

func (or *OrganizationRepository) Count(ctx context.Context, input *organization.SearchInput) (int64, error) {
	queries := rdb.New(or.db)

	params := rdb.CountSearchOrganizationsParams{
		OrganizationID:   input.OrganizationID,
		OrganizationName: input.OrganizationName,
		SearchStatus:     string(input.SearchStatus),
	}

	return queries.CountSearchOrganizations(ctx, params)
}

func (or OrganizationRepository) FindByID(ctx context.Context, id int64) (*rdb.Organization, error) {
	queries := rdb.New(or.db)

	org, err := queries.GetOrganizationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

func (or *OrganizationRepository) Search(ctx context.Context, input *organization.SearchInput) ([]rdb.Organization, error) {
	queries := rdb.New(or.db)

	params := rdb.SearchOrganizationsParams{
		OrganizationID:   input.OrganizationID,
		OrganizationName: input.OrganizationName,
		SearchStatus:     string(input.SearchStatus),
		Limit:            int32(input.PerPage),
	}

	// offsetの計算
	if input.Page > 0 {
		params.Offset = int32((input.Page - 1) * input.PerPage)
	} else {
		params.Offset = 0
	}

	return queries.SearchOrganizations(ctx, params)
}

func (or *OrganizationRepository) Create(ctx context.Context, name string) (*rdb.Organization, error) {
	queries := rdb.New(or.db)

	org, err := queries.CreateOrganization(ctx, name)
	if err != nil {
		return nil, err
	}

	return &org, err
}

func (or *OrganizationRepository) Update(ctx context.Context, input *organization.UpdateInput) error {
	queries := rdb.New(or.db)

	return queries.UpdateOrganization(ctx, rdb.UpdateOrganizationParams{
		OrganizationID:   input.OrganizationID,
		OrganizationName: input.OrganizationName,
	})
}

func (or OrganizationRepository) Delete(ctx context.Context, id int64) error {
	queries := rdb.New(or.db)

	return queries.SoftDeleteOrganization(ctx, id)
}

func (or OrganizationRepository) Restore(ctx context.Context, id int64) error {
	queries := rdb.New(or.db)

	return queries.SoftDeleteOrganization(ctx, id)
}
