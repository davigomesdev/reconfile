package supplier

import (
	"context"

	supplierDTO "github.com/davigomesdev/reconfile/internal/application/dtos/supplier"
	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type SearchSupplierUseCase struct {
	supplierRepository repositories.SupplierRepositoryInterface
}

func NewSearchSupplierUseCase(supplierRepository repositories.SupplierRepositoryInterface) *SearchSupplierUseCase {
	return &SearchSupplierUseCase{supplierRepository: supplierRepository}
}

func (uc *SearchSupplierUseCase) Execute(ctx context.Context, input supplierDTO.SearchSupplierDTO) (*contracts.SearchResult[*entities.SupplierEntity], error) {
	searchInput := &contracts.SearchInput{
		Page:    input.Page,
		PerPage: input.PerPage,
		Sort:    input.Sort,
		SortDir: input.SortDir,
		Filter:  input.Filter,
	}

	return uc.supplierRepository.Search(ctx, searchInput)
}
