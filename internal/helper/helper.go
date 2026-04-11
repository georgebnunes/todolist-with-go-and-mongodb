package helper

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/service"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func MapServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrTodoNotFound):
		ErrorResponse(w, http.StatusNotFound, err.Error())
	case errors.Is(err, service.ErrTitleRequired):
		ErrorResponse(w, http.StatusBadRequest, err.Error())
	default:
		ErrorResponse(w, http.StatusInternalServerError, "internal server error")
	}
}
