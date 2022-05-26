package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
	"github.com/abdussalamfaqih/rest-sqlite/pkg/db"
)

type sqliteUserRepo struct {
	db db.Adapter
}

func NewSqliteUserRepo(db db.Adapter) User {
	return &sqliteUserRepo{
		db: db,
	}
}

func (repo *sqliteUserRepo) GetUserByUserPassword(ctx context.Context, user domain.LoginData) (domain.User, error) {
	var result domain.User

	query := `
			SELECT 
				id,
				name
			FROM
				users
			WHERE 
				name = ?
				AND password = ?
			`

	row := repo.db.QueryRow(ctx, query, user.Name, user.Password)

	err := row.Scan(
		&result.ID,
		&result.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return result, nil
}

func (repo *sqliteUserRepo) GetUserByID(ctx context.Context, userID int) (domain.User, error) {
	var result domain.User

	query := `
			SELECT 
				id,
				name
			FROM
				users
			WHERE 
				id = ?
			`

	row := repo.db.QueryRow(ctx, query, userID)

	err := row.Scan(
		&result.ID,
		&result.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return result, nil
}
