package bootstrap

import (
	"github.com/abdussalamfaqih/rest-sqlite/internal/appconfig"
	"github.com/abdussalamfaqih/rest-sqlite/pkg/db"
)

func NewSqliteDB(cfg appconfig.Database) db.Adapter {
	db, _ := db.NewSqlite(&db.Config{
		Name:         cfg.Name,
		InternalPath: "database/local_db",
	})
	return db
}
