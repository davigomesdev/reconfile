package user

type UpdatePasswordUserDTO struct {
	ID          string `json:"id"`
	OldPassword string `json:"oldpassword" binding:"required" label:"Senha atual"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=255" label:"Nova senha"`
}
