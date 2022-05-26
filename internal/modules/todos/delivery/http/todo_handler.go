package http

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/domain"
	"github.com/abdussalamfaqih/rest-sqlite/internal/modules/todos/usecase"
	"github.com/gorilla/mux"
)

type (
	ResponsePayload struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

type TodoHandler struct {
	todos usecase.TodoUseCase
}

func NewTodoHandler(r *mux.Router, todos usecase.TodoUseCase) {
	handler := &TodoHandler{
		todos: todos,
	}

	r.HandleFunc("/todos", handler.GetTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", handler.CreateTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", handler.GetTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", handler.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", handler.DeleteTodoHandler).Methods(http.MethodDelete)
}

func (handler *TodoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.Header.Get("X-User-Id")

	uID, _ := strconv.Atoi(userID)

	result, err := handler.todos.FetchTodos(ctx, int64(uID))
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    result,
	})
}

func (handler *TodoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.Header.Get("X-User-Id")

	uID, _ := strconv.Atoi(userID)

	todoID := mux.Vars(r)["id"]

	dcID, _ := base64.RawStdEncoding.DecodeString(todoID)
	tdID, _ := strconv.Atoi(string(dcID))

	result, err := handler.todos.FetchByID(ctx, int64(uID), int64(tdID))
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	if result.ID == "" {
		w.WriteHeader(http.StatusBadRequest)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "DATA_NOT_FOUND",
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Data:    result,
	})
}

func (handler *TodoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var reqData domain.CreateRequest

	ctx := r.Context()

	userID := r.Header.Get("X-User-Id")

	uID, _ := strconv.Atoi(userID)
	json.NewDecoder(r.Body).Decode(&reqData)

	result, err := handler.todos.CreateTodo(ctx, int64(uID), reqData)
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload{
		Code:    http.StatusCreated,
		Message: "SUCCESS",
		Data:    result,
	})
}

func (handler *TodoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var reqData domain.UpdateRequest

	ctx := r.Context()

	userID := r.Header.Get("X-User-Id")

	uID, _ := strconv.Atoi(userID)

	todoID := mux.Vars(r)["id"]

	dcID, _ := base64.RawStdEncoding.DecodeString(todoID)
	tdID, _ := strconv.Atoi(string(dcID))

	json.NewDecoder(r.Body).Decode(&reqData)

	resultGet, err := handler.todos.FetchByID(ctx, int64(uID), int64(tdID))
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	if resultGet.ID == "" {
		w.WriteHeader(http.StatusBadRequest)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "DATA_NOT_FOUND",
		})

		return
	}

	result, err := handler.todos.UpdateTodo(ctx, int64(uID), int64(tdID), reqData)
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
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

func (handler *TodoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.Header.Get("X-User-Id")

	uID, _ := strconv.Atoi(userID)

	todoID := mux.Vars(r)["id"]

	dcID, _ := base64.RawStdEncoding.DecodeString(todoID)

	tdID, _ := strconv.Atoi(string(dcID))

	resultGet, err := handler.todos.FetchByID(ctx, int64(uID), int64(tdID))
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	if resultGet.ID == "" {
		w.WriteHeader(http.StatusBadRequest)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "DATA_NOT_FOUND",
		})

		return
	}

	err = handler.todos.DeleteTodo(ctx, int64(uID), int64(tdID))
	if err != nil {
		log.Printf("error occured, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponsePayload{
			Code:    http.StatusBadRequest,
			Message: "BAD_REQUEST",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponsePayload{
		Code:    http.StatusOK,
		Message: "SUCCESS",
	})
}
