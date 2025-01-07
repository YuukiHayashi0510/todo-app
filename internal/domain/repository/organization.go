package repository

import (
	"context"

	"github.com/YuukiHayashi0510/todo-app/internal/domain/organization"
	"github.com/YuukiHayashi0510/todo-app/internal/infrastructure/postgres"
	"github.com/YuukiHayashi0510/todo-app/internal/persistence/rdb"
)

type OrganizationRepository struct {
	db *postgres.DB
}

func NewOrganizationRepository(db *postgres.DB) *OrganizationRepository {
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

func (or *OrganizationRepository) FindByID(ctx context.Context, id int64) (*organization.Organization, error) {
	queries := rdb.New(or.db)

	org, err := queries.GetOrganizationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &organization.Organization{
		OrganizationID:   org.OrganizationID,
		OrganizationName: org.OrganizationName,
	}, nil
}

func (or *OrganizationRepository) Search(ctx context.Context, input *organization.SearchInput) ([]organization.Organization, error) {
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

	dbOrgs, err := queries.SearchOrganizations(ctx, params)
	if err != nil {
		return nil, err
	}

	orgs := make([]organization.Organization, 0, len(dbOrgs))
	for _, v := range dbOrgs {
		orgs = append(orgs,
			*organization.New(
				v.OrganizationID,
				v.OrganizationName,
				v.CreatedAt,
				v.UpdatedAt,
				v.DeletedAt,
			))
	}

	return orgs, nil
}

func (or *OrganizationRepository) Create(ctx context.Context, input *organization.CreateInput) (*organization.Organization, error) {
	queries := rdb.New(or.db)

	dbOrg, err := queries.CreateOrganization(ctx, input.OrganizationName)
	if err != nil {
		return nil, err
	}

	org := organization.New(
		dbOrg.OrganizationID,
		dbOrg.OrganizationName,
		dbOrg.CreatedAt,
		dbOrg.UpdatedAt,
		dbOrg.DeletedAt,
	)

	return org, nil
}

func (or *OrganizationRepository) Update(ctx context.Context, input *organization.UpdateInput) error {
	queries := rdb.New(or.db)

	return queries.UpdateOrganization(ctx, rdb.UpdateOrganizationParams{
		OrganizationID:   input.OrganizationID,
		OrganizationName: input.OrganizationName,
	})
}

func (or *OrganizationRepository) Delete(ctx context.Context, id int64) error {
	queries := rdb.New(or.db)

	return queries.SoftDeleteOrganization(ctx, id)
}

func (or *OrganizationRepository) Restore(ctx context.Context, id int64) error {
	queries := rdb.New(or.db)

	return queries.RestoreOrganization(ctx, id)
}
