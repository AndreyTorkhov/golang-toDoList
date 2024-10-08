package main

import (
    "log"
    "net/http"

    "todo-golang/internal/http-server/handlers"
    "todo-golang/storage"
    _ "todo-golang/docs"

    httpSwagger "github.com/swaggo/http-swagger"
    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// @title ToDo API
// @version 1.0
// @description API для управления списком задач.
// @host localhost:8080
// @BasePath /

func main() {
    dsn := "postgres://postgres:Aat8912000!@my_db:5432/tasks"

    db, err := storage.NewPostgresDB(dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()


    repo := storage.NewPostgresTaskRepository(db)

    h := handlers.NewTaskHandler(repo)

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/docs/*", httpSwagger.WrapHandler)
    h.SetupRoutes(r)

    log.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", r)
}