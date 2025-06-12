package supplier

import (
	"context"

	supplierDTO "github.com/davigomesdev/reconfile/internal/application/dtos/supplier"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type GetSupplierUseCase struct {
	supplierRepository repositories.SupplierRepositoryInterface
}

func NewGetSupplierUseCase(supplierRepository repositories.SupplierRepositoryInterface) *GetSupplierUseCase {
	return &GetSupplierUseCase{supplierRepository: supplierRepository}
}

func (uc *GetSupplierUseCase) Execute(ctx context.Context, input supplierDTO.GetSupplierDTO) (*entities.SupplierEntity, error) {
	return uc.supplierRepository.Get(ctx, input.ID)
}
