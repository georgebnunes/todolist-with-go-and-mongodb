# todolist-with-go-and-mongodb
Go workshop, todo list with MongoDB aiming to get familiar with Crud Operations


## Folder Structure
This is how the files in this project is structured. Useful for further review and for creating new projects based on this basic structure.

```
todolist-with-go-and-mongodb/
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   ├── domain/
│   │   └── todo.go          # Todo struct (the domain model)
│   ├── repository/
│   │   └── todo_repository.go  # MongoDB queries
│   ├── service/
│   │   └── todo_service.go  # Business logic
│   └── handler/
│       └── todo_handler.go  # Gin HTTP handlers
├── config/
│   └── config.go            # Env vars / app config
├── .env
├── docker-compose.yml
└── go.mod
```