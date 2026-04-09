package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/georgebnunes/todolist-with-go-and-mongodb/config"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/handler"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/repository"
	"github.com/georgebnunes/todolist-with-go-and-mongodb/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB is not reachable:", err)
	}

	log.Println("Connected to MongoDB successfully!")

	db := client.Database(cfg.MongoDatabase)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	mux := http.NewServeMux()
	todoHandler.RegisterRoutes(mux)

	log.Printf("Server running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, mux))
}
