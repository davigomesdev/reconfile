package auth

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email" label:"E-mail"`
	Password string `json:"password" binding:"required" label:"Senha"`
}
