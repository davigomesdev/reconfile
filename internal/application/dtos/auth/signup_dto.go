package auth

type SignUpDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=255" label:"Nome"`
	Email    string `json:"email" binding:"required,email" label:"E-mail"`
	Password string `json:"password" binding:"required,min=6,max=255" label:"Senha"`
}
