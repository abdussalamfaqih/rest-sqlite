package usecase

import (
	"context"
	"fmt"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/authenticator"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/repository"
)

type authUsecase struct {
	repo          repository.User
	authenticator authenticator.JWTAuthenticator
}

func NewAuthUsecase(repo repository.User, authenticator authenticator.JWTAuthenticator) Auth {
	return &authUsecase{repo: repo, authenticator: authenticator}
}

func (ucase *authUsecase) Login(ctx context.Context, user domain.LoginRequest) (token domain.TokenData, err error) {
	userData, err := ucase.repo.GetUserByUserPassword(ctx, domain.LoginData{
		Name:     user.Name,
		Password: user.Password,
	})

	if err != nil {
		return
	}

	if userData.ID == 0 {
		return
	}

	token, err = ucase.authenticator.GenerateAccessToken(ctx, userData)
	if err != nil {
		if err != nil {
			fmt.Printf("error authenticator, %v\n", err)
		}
		return
	}

	return
}
