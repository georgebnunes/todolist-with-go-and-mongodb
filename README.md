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

## Endpoints
| Method | Route | Description |
| :--- | :--- | :--- |
| **POST** | `/todos` | Create a todo |
| **GET** | `/todos` | List all todos |
| **GET** | `/todos/:id` | Get one todo |
| **PUT** | `/todos/:id` | Update a todo |
| **DELETE** | `/todos/:id` | Delete a todo |

## Libs used for this project

| Name | Description | Download command |
| :--- | :--- | :--- |
| **mongo-driver** | official MongoDB driver for Go | `go get go.mongodb.org/mongo-driver/mongo` |
| **godotenv** |  loads your `.env` file into environment variables | `go get github.com/joho/godotenv` |
| **uuid** |  generates uuid using Go | `go get github.com/google/uuid` |