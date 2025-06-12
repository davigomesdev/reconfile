package user

type GetUserDTO struct {
	ID string `uri:"id" binding:"required,uuid" label:"ID"`
}
