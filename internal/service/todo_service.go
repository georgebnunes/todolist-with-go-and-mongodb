package service

import (
	"context"
	"errors"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/repository"
)

var (
	ErrTodoNotFound  = errors.New("todo not found")
	ErrTitleRequired = errors.New("title is required")
)

type TodoService interface {
	Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	FindAll(ctx context.Context) ([]domain.Todo, error)
	FindByID(ctx context.Context, id string) (domain.Todo, error)
	Update(ctx context.Context, id string, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, id string) error
}

type todoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (s *todoService) FindAll(ctx context.Context) ([]domain.Todo, error) {
	return []domain.Todo{}, nil
}

func (s *todoService) FindByID(ctx context.Context, id string) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (s *todoService) Update(ctx context.Context, id string, todo domain.Todo) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (s *todoService) Delete(ctx context.Context, id string) error {
	return nil
}
