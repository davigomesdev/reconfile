package supplier

import "github.com/davigomesdev/reconfile/internal/domain/contracts"

type SearchSupplierDTO struct {
	Page    *int                     `form:"page" binding:"omitempty,gte=1" label:"Página"`
	PerPage *int                     `form:"perPage" binding:"omitempty,gte=1,lte=100" label:"Itens por página"`
	Sort    *string                  `form:"sort" binding:"omitempty,oneof=partnerId" label:"Campo de ordenação"`
	SortDir *contracts.SortDirection `form:"sortDir" binding:"omitempty,oneof=asc desc" label:"Direção de ordenação"`
	Filter  *string                  `form:"filter" binding:"omitempty,max=255" label:"Filtro"`
}
