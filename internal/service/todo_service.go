package service

import (
	"context"
	"errors"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	if todo.Title == "" {
		return domain.Todo{}, nil
	}
	return s.repo.Create(ctx, todo)

}

func (s *todoService) FindAll(ctx context.Context) ([]domain.Todo, error) {
	return s.repo.FindAll(ctx)
}

func (s *todoService) FindByID(ctx context.Context, id string) (domain.Todo, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Todo{}, ErrTodoNotFound
	}

	todo, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Todo{}, ErrTodoNotFound
		}
		return domain.Todo{}, err
	}

	return todo, nil

}

func (s *todoService) Update(ctx context.Context, id string, todo domain.Todo) (domain.Todo, error) {
	if todo.Title == "" {
		return domain.Todo{}, ErrTitleRequired
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Todo{}, ErrTodoNotFound
	}

	_, err = s.repo.FindByID(ctx, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Todo{}, ErrTodoNotFound
		}
		return domain.Todo{}, err
	}

	return s.repo.Update(ctx, objID, todo)
}

func (s *todoService) Delete(ctx context.Context, id string) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrTodoNotFound
	}

	_, err = s.repo.FindByID(ctx, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrTodoNotFound
		}
		return err
	}

	return s.repo.Delete(ctx, objID)
}
