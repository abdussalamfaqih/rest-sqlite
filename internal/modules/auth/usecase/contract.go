package usecase

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
)

type Auth interface {
	Login(ctx context.Context, user domain.LoginRequest) (token domain.TokenData, err error)
}
