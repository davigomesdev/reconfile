package auth

type RefreshTokensDTO struct {
	RefreshToken string `json:"refreshToken" binding:"required" label:"Refresh token"`
}
