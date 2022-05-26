package http

import (
	"github.com/abdussalamfaqih/rest-sqlite/internal/appconfig"
	"github.com/abdussalamfaqih/rest-sqlite/internal/bootstrap"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/authenticator"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/delivery/http/middleware"
	userRepo "github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/repository"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/delivery/http"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/repository"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/usecase"
	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router, cfg appconfig.Config) {

	// repositories
	db := bootstrap.NewSqliteDB(cfg.Database)
	repo := repository.NewSqliteTodoRepo(db)

	// usecases
	ucase := usecase.NewTodoCase(repo)

	uRepo := userRepo.NewSqliteUserRepo(db)

	// authenticator
	jwtAuth := authenticator.NewJWTAuthenticator(cfg.App)

	mf := middleware.InitMiddleware(jwtAuth, uRepo)

	r.Use(mf.JWTAuthorization)

	// handlers
	http.NewTodoHandler(r, ucase)
}
