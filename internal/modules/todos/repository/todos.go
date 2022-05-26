package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/domain"
	"github.com/abdussalamfaqih/rest-sqlite/pkg/db"
)

type sqliteTodoRepo struct {
	db db.Adapter
}

func NewSqliteTodoRepo(db db.Adapter) Todos {
	return &sqliteTodoRepo{
		db: db,
	}
}

func (repo *sqliteTodoRepo) FetchByUserID(ctx context.Context, userID int) ([]domain.Todo, error) {
	var result []domain.Todo

	query := `
			SELECT 
				id,
				name, 
				description
			FROM
				todos
			WHERE 
				user_id = ?
			`

	rows, err := repo.db.QueryRows(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		var data domain.Todo

		err = rows.Scan(
			&data.ID,
			&data.Name,
			&data.Description,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (repo *sqliteTodoRepo) FetchByID(ctx context.Context, userID, todoID int) (domain.Todo, error) {
	var result domain.Todo

	query := `
			SELECT 
				id,
				name, 
				description
			FROM
				todos
			WHERE 
				user_id = ?
				AND id = ?
			`

	row := repo.db.QueryRow(ctx, query, userID, todoID)

	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Description,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Todo{}, nil
		}
		return domain.Todo{}, err
	}

	return result, nil
}

func (repo *sqliteTodoRepo) CreateTodo(ctx context.Context, data domain.CreateData) (domain.Todo, error) {

	query := `INSERT INTO todos(user_id, name, description) VALUES (?, ?, ?)`
	res, err := repo.db.Exec(ctx, query, data.UserID, data.Name, data.Description)

	if err != nil {
		return domain.Todo{}, err
	}

	id, _ := res.LastInsertId()

	return domain.Todo{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
	}, nil

}

func (repo *sqliteTodoRepo) UpdateByID(ctx context.Context, todoID int, data domain.UpdateData) (domain.Todo, error) {

	query := `UPDATE todos SET name = ?, description = ? WHERE user_id = ?`
	res, err := repo.db.Exec(ctx, query, data.Name, data.Description, data.UserID)

	if err != nil {
		return domain.Todo{}, err
	}

	id, _ := res.LastInsertId()

	return domain.Todo{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
	}, nil

}

func (repo *sqliteTodoRepo) DeleteByID(ctx context.Context, userID, todoID int) error {

	query := `DELTE todos WHERE user_id = ? AND id = ?`
	_, err := repo.db.Exec(ctx, query, userID, todoID)

	if err != nil {
		return err
	}
	return nil
}
