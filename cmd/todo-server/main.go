package main

import (
    "log"
    "net/http"
    "todo-golang/internal/http-server/handlers"
    "todo-golang/storage"

    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    dsn := "postgres://postgres:Aat8912000!@localhost:5432"

    db, err := storage.NewPostgresDB(dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()


    repo := storage.NewPostgresTaskRepository(db)

    h := handlers.NewTaskHandler(repo)

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    h.SetupRoutes(r)

    log.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", r)
}