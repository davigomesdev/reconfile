package auth

import (
	"context"

	authDTO "github.com/davigomesdev/reconfile/internal/application/dtos/auth"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
	jwt_auth "github.com/davigomesdev/reconfile/internal/infrastructure/jwt-auth"
)

type RefreshTokensUseCase struct {
	userRepository repositories.UserRepositoryInterface
	jwtAuthService *jwt_auth.JWTAuthService
}

func NewRefreshTokensUseCase(
	userRepository repositories.UserRepositoryInterface,
	jwtAuthService *jwt_auth.JWTAuthService,
) *RefreshTokensUseCase {
	return &RefreshTokensUseCase{
		userRepository: userRepository,
		jwtAuthService: jwtAuthService,
	}
}

func (uc *RefreshTokensUseCase) Execute(ctx context.Context, input authDTO.RefreshTokensDTO) (*jwt_auth.AuthTokens, error) {
	claims, err := uc.jwtAuthService.VerifyToken(input.RefreshToken)
	if err != nil {
		return nil, err
	}

	entity, err := uc.userRepository.Get(ctx, claims.ID)
	if err != nil {
		return nil, err
	}

	tokens, err := uc.jwtAuthService.GenerateTokens(entity.ID, entity.Name)
	if err != nil {
		return nil, err
	}

	entity.UpdateRefreshToken(&tokens.RefreshToken)

	if err = uc.userRepository.Update(ctx, entity); err != nil {
		return nil, err
	}

	return tokens, nil
}
