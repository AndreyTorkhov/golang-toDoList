package handlers

import (
	"github.com/go-chi/chi/v5"
)


func (h *TaskHandler) SetupRoutes(r *chi.Mux) {
    r.Get("/tasks", h.GetTasks)
    r.Get("/tasks/{id}", h.GetTaskByID)
    r.Post("/tasks", h.CreateTask)
    r.Delete("/tasks/{id}", h.DeleteTask)
    r.Patch("/tasks/{id}/done", h.MarkTaskDone)
    r.Get("/tasks/filter", h.GetFilteredTasks)
}