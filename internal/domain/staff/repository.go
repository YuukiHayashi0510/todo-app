package staff

import (
	"context"
)

type Repository interface {
	Count(ctx context.Context, input *SearchInput) (int64, error)
	FindByID(ctx context.Context, id int64) (*Staff, error)
	Search(ctx context.Context, input *SearchInput) ([]Staff, error)
	Create(ctx context.Context, input *CreateInput) (*Staff, error)
	Update(ctx context.Context, input *UpdateInput) error
	Delete(ctx context.Context, id int64) error
	Restore(ctx context.Context, id int64) error
}