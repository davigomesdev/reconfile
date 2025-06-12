package user

import (
	"context"

	userDTO "github.com/davigomesdev/reconfile/internal/application/dtos/user"
	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
)

type SearchUserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewSearchUserUseCase(userRepository repositories.UserRepositoryInterface) *SearchUserUseCase {
	return &SearchUserUseCase{userRepository: userRepository}
}

func (uc *SearchUserUseCase) Execute(ctx context.Context, input userDTO.SearchUserDTO) (*contracts.SearchResult[*entities.UserEntity], error) {
	searchInput := &contracts.SearchInput{
		Page:    input.Page,
		PerPage: input.PerPage,
		Sort:    input.Sort,
		SortDir: input.SortDir,
		Filter:  input.Filter,
	}

	return uc.userRepository.Search(ctx, searchInput)
}
