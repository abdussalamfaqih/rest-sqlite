package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/authenticator"
	httpAuth "github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/delivery/http"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/repository"
)

type GoMiddleware struct {
	auth     authenticator.JWTAuthenticator
	userRepo repository.User
}

func InitMiddleware(auth authenticator.JWTAuthenticator, userRepo repository.User) *GoMiddleware {
	return &GoMiddleware{
		auth:     auth,
		userRepo: userRepo,
	}
}

// JWTAuthorization will handle the jwt middleware
func (m *GoMiddleware) JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/auth") {
			next.ServeHTTP(w, r)
		}
		authHeader := r.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(httpAuth.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: "INVALID_TOKEN",
			})
			return
		}
		authToken := strings.Split(authHeader, " ")[1]

		ctx := r.Context()

		tokenData, err := m.auth.ValidateAccessToken(ctx, authToken)
		if err != nil {
			log.Printf("middleware error: %v\n", err)

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(httpAuth.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: "INVALID_TOKEN",
			})
			return
		}

		user, err := m.userRepo.GetUserByID(ctx, tokenData.UserID)
		if err != nil {
			log.Printf("middleware error: %v\n", err)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(httpAuth.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: "INVALID_TOKEN",
			})

			return
		}

		if user.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(httpAuth.ResponsePayload{
				Code:    http.StatusBadRequest,
				Message: "DATA_NOT_FOUND",
			})

			return
		}

		r.Header.Set("X-User-Id", strconv.Itoa(int(user.ID)))

		next.ServeHTTP(w, r)
	})
}
