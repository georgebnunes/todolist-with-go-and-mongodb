package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (h *TodoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /todos", h.Create)
	mux.HandleFunc("GET /todos", h.FindAll)
	mux.HandleFunc("GET /todos/{id}", h.FindByID)
	mux.HandleFunc("PUT /todos/{id}", h.Update)
	mux.HandleFunc("DELETE /todos/{id}", h.Delete)
}

// -- Helpers --
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func mapServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrTodoNotFound):
		errorResponse(w, http.StatusNotFound, err.Error())
	case errors.Is(err, service.ErrTitleRequired):
		errorResponse(w, http.StatusBadRequest, err.Error())
	default:
		errorResponse(w, http.StatusInternalServerError, "internal server error")
	}
}

// -- handlers --
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := newContext()
	defer cancel()

	created, err := h.service.Create(ctx, todo)
	if err != nil {
		mapServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *TodoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := newContext()
	defer cancel()

	todos, err := h.service.FindAll(ctx)
	if err != nil {
		mapServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	ctx, cancel := newContext()
	defer cancel()

	todo, err := h.service.FindByID(ctx, id)
	if err != nil {
		mapServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := newContext()
	defer cancel()

	updated, err := h.service.Update(ctx, id, todo)
	if err != nil {
		mapServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	ctx, cancel := newContext()
	defer cancel()

	if err := h.service.Delete(ctx, id); err != nil {
		mapServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}
