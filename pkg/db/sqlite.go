package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

func CreateSession(cfg *Config) (*sql.DB, error) {
	session, err := sql.Open("sqlite3", fmt.Sprintf("%s/%s.db", cfg.InternalPath, cfg.Name))
	return session, err
}

func NewSqlite(cfg *Config) (Adapter, error) {
	session, err := CreateSession(cfg)

	if err != nil {
		log.Fatalf("error creating db session, %v", err)
		return nil, err
	}
	return &sqlite{
		db: session,
	}, nil
}

func (d *sqlite) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *sqlite) QueryRows(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *sqlite) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}
