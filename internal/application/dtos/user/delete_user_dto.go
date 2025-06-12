package user

type DeleteUserDTO struct {
	ID string `uri:"id" binding:"required,uuid" label:"ID"`
}
