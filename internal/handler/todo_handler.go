package handler

import (
	"net/http"

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

}
