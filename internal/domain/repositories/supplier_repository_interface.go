package repositories

import (
	"context"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
)

type SupplierRepositoryInterface interface {
	Get(ctx context.Context, id string) (*entities.SupplierEntity, error)
	GetAll(ctx context.Context) ([]*entities.SupplierEntity, error)
	GetOverview(ctx context.Context) (*contracts.SupplierOverview, error)
	Search(ctx context.Context, input *contracts.SearchInput) (*contracts.SearchResult[*entities.SupplierEntity], error)
	CreateMany(ctx context.Context, entities []*entities.SupplierEntity) error
}
