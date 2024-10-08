package handlers

import (
    "encoding/json"
    "net/http"
	"strconv"

    "todo-golang/storage"
)

type TaskHandler struct {
    repo storage.TaskRepository
}

func NewTaskHandler(repo storage.TaskRepository) *TaskHandler {
    return &TaskHandler{repo: repo}
}

// GetFilteredTasks
// @Summary Получить отфильтрованный список задач
// @Description Возвращает список задач на основе статуса выполнения
// @Tags tasks
// @Accept json
// @Produce json
// @Param done query bool true "Статус выполнения (true - выполненные, false - не выполненные)"
// @Success 200 {array} model.Task "Список задач"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /tasks/filtered [get]
func (h *TaskHandler) GetFilteredTasks(w http.ResponseWriter, r *http.Request) {
    doneStr := r.URL.Query().Get("done")
    var doneFilter *bool

    if doneStr != "" {
        done, err := strconv.ParseBool(doneStr)
        if err != nil {
            http.Error(w, "Invalid 'done' query parameter", http.StatusBadRequest)
            return
        }
        doneFilter = &done
    }

    tasks, err := h.repo.GetFiltered(doneFilter)
    if err != nil {
        http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}