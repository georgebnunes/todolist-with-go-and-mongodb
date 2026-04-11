package repository

import (
	"context"
	"time"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	FindAll(ctx context.Context) ([]domain.Todo, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (domain.Todo, error)
	Update(ctx context.Context, id primitive.ObjectID, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

// Implementation for mongodb
type mongoTodoRepo struct {
	collection *mongo.Collection
}

func NewTodoRepository(db *mongo.Database) TodoRepository {
	return &mongoTodoRepo{
		collection: db.Collection("todos"),
	}
}

func (r *mongoTodoRepo) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {

	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, todo)
	if err != nil {
		return domain.Todo{}, nil
	}

	return todo, nil
}

func (r *mongoTodoRepo) FindAll(ctx context.Context) ([]domain.Todo, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []domain.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	if todos == nil {
		return []domain.Todo{}, nil
	}

	return todos, nil
}

func (r *mongoTodoRepo) FindByID(ctx context.Context, id primitive.ObjectID) (domain.Todo, error) {
	var todo domain.Todo

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return domain.Todo{}, nil
	}

	return todo, nil
}

func (r *mongoTodoRepo) Update(ctx context.Context, id primitive.ObjectID, todo domain.Todo) (domain.Todo, error) {

	todo.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":       todo.Title,
			"description": todo.Description,
			"done":        todo.Done,
			"updated_at":  todo.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return domain.Todo{}, err
	}

	return r.FindByID(ctx, id)
}

func (r *mongoTodoRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
