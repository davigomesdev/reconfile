package supplier

type GetSupplierDTO struct {
	ID string `uri:"id" binding:"required,uuid" label:"ID"`
}
