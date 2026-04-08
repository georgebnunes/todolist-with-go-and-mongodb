package repository

import (
	"context"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	FindAll(ctx context.Context) ([]domain.Todo, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (domain.Todo, error)
	Update(ctx context.Context, id primitive.ObjectID, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
