package supplier

import (
	"context"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type OverviewSupplierUseCase struct {
	supplierRepository repositories.SupplierRepositoryInterface
}

func NewOverviewSupplierUseCase(supplierRepository repositories.SupplierRepositoryInterface) *OverviewSupplierUseCase {
	return &OverviewSupplierUseCase{supplierRepository: supplierRepository}
}

func (uc *OverviewSupplierUseCase) Execute(ctx context.Context) (*contracts.SupplierOverview, error) {
	return uc.supplierRepository.GetOverview(ctx)
}
