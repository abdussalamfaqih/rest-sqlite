package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/abdussalamfaqih/rest-sqlite/internal/appconfig"
	"github.com/gorilla/mux"
)

func Start(ctx context.Context) {
	router := mux.NewRouter()

	cfg := appconfig.LoadConfig()

	RegisterHandlers(router, cfg)
	RegisterAuth(router, cfg)

	// Start the server
	log.Printf("Starting Server on port %s\n", cfg.App.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.App.Port), router))
}
