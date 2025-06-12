package user

type UpdateUserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required,min=3,max=255" label:"Nome"`
	Email string `json:"email" binding:"required,email" label:"E-mail"`
}
