package staff

import (
	"context"
	"errors"

	"github.com/YuukiHayashi0510/todo-app/internal/app/common"
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

	staffs, err := s.repository.Search(ctx, input)
	if err != nil {
		return nil, err
	}

	return &SearchOutput{
		Staffs:   staffs,
		PageInfo: common.NewPageInfoWith(input.Page, input.PerPage, totalCount),
	}, nil
}

func (s *Service) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	// 作成したスタッフ情報の取得
	createStaff, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	// 作成時にJOINできないため、FindByを実行してOrganizations含めて取得
	staff, err := s.repository.FindByID(ctx, createStaff.StaffID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStaffNotFound
		}
		return nil, err
	}

	return &CreateOutput{
		Staff: *staff,
	}, nil
}

func (s *Service) Update(ctx context.Context, input *UpdateInput) (*UpdateOutput, error) {
	// 存在確認
	_, err := s.repository.FindByID(ctx, input.StaffID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStaffNotFound
		}
		return nil, err
	}

	err = s.repository.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	// 更新後の値の再取得
	staff, err := s.repository.FindByID(ctx, input.StaffID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStaffNotFound
		}
		return nil, err
	}

	return &UpdateOutput{
		Staff: *staff,
	}, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	// 存在確認
	staff, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrStaffNotFound
		}
		return err
	}

	// 既に削除されている場合
	if staff.DeletedAt != nil {
		return ErrStaffHasAlreadyDeleted
	}

	// 削除
	err = s.repository.Delete(ctx, staff.StaffID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Restore(ctx context.Context, id int64) error {
	// 存在確認
	staff, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrStaffNotFound
		}
		return err
	}

	// 削除されていない場合
	if staff.DeletedAt == nil {
		return ErrStaffIsNotDeleted
	}

	// 削除
	err = s.repository.Restore(ctx, staff.StaffID)
	if err != nil {
		return err
	}

	return nil
}
