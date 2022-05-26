package usecase

import (
	"context"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/domain"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/repository"
)

type todoUseCase struct {
	repo repository.Todos
}

func NewTodoCase(todoRepo repository.Todos) TodoUseCase {
	return &todoUseCase{
		repo: todoRepo,
	}
}

func (ucase *todoUseCase) FetchTodos(ctx context.Context, userID int64) ([]domain.TodoData, error) {

	var result []domain.TodoData

	data, err := ucase.repo.FetchByUserID(ctx, int(userID))
	if err != nil {
		return nil, err
	}

	for _, d := range data {
		result = append(result, d.TransformToPresentation())
	}

	return result, nil
}

func (ucase *todoUseCase) CreateTodo(ctx context.Context, userID int64, data domain.CreateRequest) (domain.TodoData, error) {
	result, err := ucase.repo.CreateTodo(ctx, domain.CreateData{
		UserID:      userID,
		Name:        data.Name,
		Description: data.Description,
	})

	if err != nil {
		return domain.TodoData{}, err
	}

	return result.TransformToPresentation(), nil
}

func (ucase *todoUseCase) FetchByID(ctx context.Context, userID, todoID int64) (domain.TodoData, error) {

	data, err := ucase.repo.FetchByID(ctx, int(userID), int(todoID))
	if err != nil {
		return domain.TodoData{}, err
	}

	if data.ID == 0 {
		return domain.TodoData{
			ID: "",
		}, nil
	}

	return data.TransformToPresentation(), nil
}

func (ucase *todoUseCase) UpdateTodo(ctx context.Context, userID, todoID int64, data domain.UpdateRequest) (domain.TodoData, error) {
	result, err := ucase.repo.UpdateByID(ctx, int(todoID), domain.UpdateData{
		UserID:      userID,
		Name:        data.Name,
		Description: data.Description,
	})

	if err != nil {
		return domain.TodoData{}, err
	}

	return result.TransformToPresentation(), nil
}

func (ucase *todoUseCase) DeleteTodo(ctx context.Context, userID, todoID int64) error {

	err := ucase.repo.DeleteByID(ctx, int(userID), int(todoID))
	if err != nil {
		return err
	}

	return nil
}
