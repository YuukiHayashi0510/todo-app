package organization

import (
	"context"

	"github.com/YuukiHayashi0510/todo-app/internal/persistence/rdb"
)

type Repository interface {
	Count(ctx context.Context, input *SearchInput) (int64, error)
	FindByID(ctx context.Context, id int64) (*rdb.Organization, error)
	Search(ctx context.Context, input *SearchInput) ([]rdb.Organization, error)
	Create(ctx context.Context, name string) (*rdb.Organization, error)
	Update(ctx context.Context, input *UpdateInput) error
	Delete(ctx context.Context, id int64) error
	Restore(ctx context.Context, id int64) error
}
