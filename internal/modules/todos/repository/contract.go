package repository

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/domain"
)

type Todos interface {
	FetchByUserID(ctx context.Context, userID int) ([]domain.Todo, error)
	CreateTodo(ctx context.Context, data domain.CreateData) (domain.Todo, error)
	FetchByID(ctx context.Context, userID, todoID int) (domain.Todo, error)
	UpdateByID(ctx context.Context, todoID int, data domain.UpdateData) (domain.Todo, error)
	DeleteByID(ctx context.Context, userID, todoID int) error
}
