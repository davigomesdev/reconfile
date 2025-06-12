package supplier

import "mime/multipart"

type ImportSuppliersDTO struct {
	File *multipart.FileHeader `form:"file" binding:"required" label:"Arquivo de Fornecedores"`
}
