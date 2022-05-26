package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/domain"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/auth/usecase"
	"github.com/gorilla/mux"
)

type (
	ResponsePayload struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)
type AuthHandler struct {
	auth usecase.Auth
}

func NewAuthHandler(r *mux.Router, auth usecase.Auth) {
	handler := &AuthHandler{
		auth: auth,
	}

	r.HandleFunc("/auth/login", handler.Login).Methods(http.MethodPost)

}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	ctx, reqData := r.Context(), domain.LoginRequest{}
	json.NewDecoder(r.Body).Decode(&reqData)

	result, err := handler.auth.Login(ctx, reqData)

	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	if result.Token == "" {
		w.WriteHeader(http.StatusBadRequest)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "DATA_NOT_FOUND",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    result,
	})
}
