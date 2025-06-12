package handlers

import (
	"net/http"

	"github.com/davigomesdev/reconfile/internal/adapters/httpx"
	"github.com/davigomesdev/reconfile/internal/adapters/presenters"
	supplierDTO "github.com/davigomesdev/reconfile/internal/application/dtos/supplier"
	supplierUC "github.com/davigomesdev/reconfile/internal/application/usecases/supplier"
	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	getSupplierUseCase      *supplierUC.GetSupplierUseCase
	searchSupplierUseCase   *supplierUC.SearchSupplierUseCase
	overviewSupplierUseCase *supplierUC.OverviewSupplierUseCase
	importSuppliersUseCase  *supplierUC.ImportSuppliersUseCase
}

func NewSupplierHandler(
	getSupplierUseCase *supplierUC.GetSupplierUseCase,
	searchSupplierUseCase *supplierUC.SearchSupplierUseCase,
	overviewSupplierUseCase *supplierUC.OverviewSupplierUseCase,
	importSuppliersUseCase *supplierUC.ImportSuppliersUseCase,
) *SupplierHandler {
	return &SupplierHandler{
		getSupplierUseCase:      getSupplierUseCase,
		searchSupplierUseCase:   searchSupplierUseCase,
		overviewSupplierUseCase: overviewSupplierUseCase,
		importSuppliersUseCase:  importSuppliersUseCase,
	}
}

func (h *SupplierHandler) Get(c *gin.Context) {
	input, ok := httpx.ParamValidate[supplierDTO.GetSupplierDTO](c)
	if !ok {
		return
	}

	customer, err := h.getSupplierUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewSupplierPresenter(customer))
}

func (h *SupplierHandler) Overview(c *gin.Context) {
	supplierOverview, err := h.overviewSupplierUseCase.Execute(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewSupplierOverviewPresenter(supplierOverview))
}

func (h *SupplierHandler) Search(c *gin.Context) {
	input, ok := httpx.QueryValidate[supplierDTO.SearchSupplierDTO](c)
	if !ok {
		return
	}

	searchResult, err := h.searchSupplierUseCase.Execute(c.Request.Context(), *input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, presenters.NewSupplierCollectionPresenter(searchResult))
}

func (h *SupplierHandler) Import(c *gin.Context) {
	input, ok := httpx.FormValidate[supplierDTO.ImportSuppliersDTO](c)
	if !ok {
		return
	}

	if err := h.importSuppliersUseCase.Execute(c.Request.Context(), *input); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
