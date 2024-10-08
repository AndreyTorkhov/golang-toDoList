package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"

    "todo-golang/internal/config"
)

// GetTasks
// @Summary Получить список задач
// @Description Возвращает список всех задач
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} model.Task "Список задач"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.repo.GetAll()
    if err != nil {
        http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID
// @Summary Получить задачу по идентификатору
// @Description Возвращает задачу по указанному идентификатору
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} model.Task "Задача"
// @Failure 404 {object} map[string]string "Задача не найдена"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    task, err := h.repo.GetByID(id)
    if err != nil {
        if err.Error() == "task not found" {
            http.Error(w, "Task not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

// CreateTask
// @Summary Создать новую задачу
// @Description Добавляет новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Создание задачи"
// @Success 201 {object} model.Task "Созданная задача"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks [post]
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

// DeleteTask
// @Summary Удалить задачу
// @Description Удаляет задачу по идентификатору
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 204 "Задача успешно удалена"
// @Failure 400 {object} map[string]string "Некорректный идентификатор задачи"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks/{id} [delete]
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

// MarkTaskDone
// @Summary Пометить задачу как выполненную
// @Description Помечает задачу как выполненную по идентификатору
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} model.Task "Задача помечена как выполненная"
// @Failure 400 {object} map[string]string "Некорректный идентификатор задачи"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks/{id}/done [patch]
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
