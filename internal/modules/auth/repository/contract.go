package repository

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
)

type User interface {
	GetUserByUserPassword(ctx context.Context, user domain.LoginData) (domain.User, error)
	GetUserByID(ctx context.Context, userID int) (domain.User, error)
}
