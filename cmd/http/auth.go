package http

import (
	"github.com/abdussalamfaqih/rest-sqlite/internal/appconfig"
	"github.com/abdussalamfaqih/rest-sqlite/internal/bootstrap"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/authenticator"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/delivery/http"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/repository"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/usecase"
	"github.com/gorilla/mux"
)

func RegisterAuth(r *mux.Router, cfg appconfig.Config) {
	db := bootstrap.NewSqliteDB(cfg.Database)
	repo := repository.NewSqliteUserRepo(db)

	// authenticator
	jwtAuth := authenticator.NewJWTAuthenticator(cfg.App)

	// usecases
	ucase := usecase.NewAuthUsecase(repo, jwtAuth)

	http.NewAuthHandler(r, ucase)
}
