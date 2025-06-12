package repositories

import (
	"context"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
)

type UserRepositoryInterface interface {
	EmailExists(ctx context.Context, email string) error
	Get(ctx context.Context, id string) (*entities.UserEntity, error)
	GetByEmail(ctx context.Context, email string) (*entities.UserEntity, error)
	GetAll(ctx context.Context) ([]*entities.UserEntity, error)
	Search(ctx context.Context, input *contracts.SearchInput) (*contracts.SearchResult[*entities.UserEntity], error)
	Create(ctx context.Context, entity *entities.UserEntity) error
	Update(ctx context.Context, entity *entities.UserEntity) error
	Delete(ctx context.Context, id string) error
}
