package presenters

import (
	"github.com/davigomesdev/reconfile/internal/adapters/collections"
	"github.com/davigomesdev/reconfile/internal/domain/contracts"
)

func NewSupplierOverviewPresenter(so *contracts.SupplierOverview) *collections.Presenter[*contracts.SupplierOverview] {
	return &collections.Presenter[*contracts.SupplierOverview]{
		Data: so,
	}

}
