package jwt_auth

import (
	"fmt"
	"time"

	env_config "github.com/davigomesdev/reconfile/internal/infrastructure/env-config"
	"github.com/golang-jwt/jwt/v5"
)

type AuthTokens struct {
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	AccessExpiresIn  int64  `json:"accessExpiresIn"`
	RefreshExpiresIn int64  `json:"refreshExpiresIn"`
}

type JWTClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type JWTAuthService struct {
	env *env_config.EnvConfig
}

func NewJWTAuthService(env *env_config.EnvConfig) *JWTAuthService {
	return &JWTAuthService{env: env}
}

func (ja *JWTAuthService) GenerateTokens(id string, name string) (*AuthTokens, error) {
	now := time.Now()

	accessExp := now.Add(ja.env.JWTExpiresAccessIn)
	refreshExp := now.Add(ja.env.JWTExpiresRefreshIn)

	accessToken, err := ja.signToken(id, name, accessExp)
	if err != nil {
		return nil, err
	}

	refreshToken, err := ja.signToken(id, name, refreshExp)
	if err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  int64(ja.env.JWTExpiresAccessIn.Seconds()),
		RefreshExpiresIn: int64(ja.env.JWTExpiresRefreshIn.Seconds()),
	}, nil
}

func (s *JWTAuthService) signToken(id, name string, exp time.Time) (string, error) {
	claims := JWTClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "reconfile-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.env.JWTSecretKey))
}

func (ja *JWTAuthService) VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.env.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
