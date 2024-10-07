package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

	"github.com/go-chi/chi/v5"

    "todo-golang/internal/config"
    "todo-golang/storage"
)

type TaskHandler struct {
    repo storage.TaskRepository
}

func NewTaskHandler(repo storage.TaskRepository) *TaskHandler {
    return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.repo.Add(task); err != nil {
        http.Error(w, "Failed to add task", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    if err := h.repo.Delete(id); err != nil {
        http.Error(w, "Failed to delete task", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) MarkTaskDone(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    if err := h.repo.MarkDone(id); err != nil {
        http.Error(w, "Failed to mark task as done", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) SetupRoutes(r *chi.Mux) {
    r.Get("/tasks", h.GetTasks)
    r.Post("/tasks", h.CreateTask)
    r.Delete("/tasks/{id}", h.DeleteTask)
    r.Patch("/tasks/{id}/done", h.MarkTaskDone)
}