package organization

import (
	"context"
	"errors"

	"github.com/YuukiHayashi0510/todo-app/internal/domain/common"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Search(ctx context.Context, input *SearchInput) (*SearchOutput, error) {
	totalCount, err := s.repository.Count(ctx, input)
	if err != nil {
		return nil, err
	}

	dbOrgs, err := s.repository.Search(ctx, input)
	if err != nil {
		return nil, err
	}

	orgs := make([]Organization, 0, len(dbOrgs))
	for _, v := range dbOrgs {
		orgs = append(orgs, *New(
			v.OrganizationID,
			v.OrganizationName,
			v.CreatedAt,
			v.UpdatedAt,
			v.DeletedAt,
		))
	}

	return &SearchOutput{
		Organizations: orgs,
		PageInfo:      common.NewPageInfoWith(input.Page, input.PerPage, totalCount),
	}, nil
}

func (s *Service) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	org, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &CreateOutput{
		Organization: *New(
			org.OrganizationID,
			org.OrganizationName,
			org.CreatedAt,
			org.UpdatedAt,
			org.DeletedAt,
		),
	}, nil
}

func (s *Service) Update(ctx context.Context, input *UpdateInput) (*UpdateOutput, error) {
	// 存在確認
	_, err := s.repository.FindByID(ctx, input.OrganizationID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrOrganizationNotFound
		}
		return nil, err
	}

	err = s.repository.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	// 更新後の値の再取得
	org, err := s.repository.FindByID(ctx, input.OrganizationID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrOrganizationNotFound
		}
		return nil, err
	}

	return &UpdateOutput{
		Organization: *New(
			org.OrganizationID,
			org.OrganizationName,
			org.CreatedAt,
			org.UpdatedAt,
			org.DeletedAt,
		),
	}, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	// 存在確認
	org, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrOrganizationNotFound
		}
		return err
	}

	// 既に削除されている場合
	if org.DeletedAt != nil {
		return ErrOrganizationHasAlreadyDeleted
	}

	// 削除
	err = s.repository.Delete(ctx, org.OrganizationID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Restore(ctx context.Context, id int64) error {
	// 存在確認
	org, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrOrganizationNotFound
		}
		return err
	}

	// 削除されていない場合
	if org.DeletedAt == nil {
		return ErrOrganizationIsNotDeleted
	}

	// 削除
	err = s.repository.Restore(ctx, org.OrganizationID)
	if err != nil {
		return err
	}

	return nil
}
