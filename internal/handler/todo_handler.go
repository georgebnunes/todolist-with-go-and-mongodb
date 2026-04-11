package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/helper"
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

// -- handlers --
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := helper.NewContext()
	defer cancel()

	created, err := h.service.Create(ctx, todo)
	if err != nil {
		helper.MapServiceError(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusCreated, map[string]any{"data": created})
}

func (h *TodoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := helper.NewContext()
	defer cancel()

	todos, err := h.service.FindAll(ctx)
	if err != nil {
		helper.MapServiceError(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]any{"data": todos})
}

func (h *TodoHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	ctx, cancel := helper.NewContext()
	defer cancel()

	todo, err := h.service.FindByID(ctx, id)
	if err != nil {
		helper.MapServiceError(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]any{"data": todo})
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	ctx, cancel := helper.NewContext()
	defer cancel()

	updated, err := h.service.Update(ctx, id, todo)
	if err != nil {
		helper.MapServiceError(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]any{"data": updated})
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.PathValue("id"), "/")

	ctx, cancel := helper.NewContext()
	defer cancel()

	if err := h.service.Delete(ctx, id); err != nil {
		helper.MapServiceError(w, err)
		return
	}

	helper.WriteJSON(w, http.StatusNoContent, nil)
}
