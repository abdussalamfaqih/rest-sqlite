package usecase

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/domain"
)

type TodoUseCase interface {
	FetchTodos(ctx context.Context, userID int64) ([]domain.TodoData, error)
	CreateTodo(ctx context.Context, userID int64, data domain.CreateRequest) (domain.TodoData, error)
	FetchByID(ctx context.Context, userID, todoID int64) (domain.TodoData, error)
	UpdateTodo(ctx context.Context, userID, todoID int64, data domain.UpdateRequest) (domain.TodoData, error)
	DeleteTodo(ctx context.Context, userID, todoID int64) error
}
