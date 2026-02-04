package handlers

import (
	"StillCode/server/internal/models"
	"StillCode/server/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TaskHandler handles task-related HTTP requests
type TaskHandler struct {
	taskService *services.TaskService
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// GetTasks handles GET /api/tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	var params models.TasksQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		params = models.TasksQueryParams{Page: 1, PageSize: 20}
	}

	// Ensure valid pagination
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 20
	}

	tasks, err := h.taskService.GetTasks(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles GET /api/tasks/:id
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := h.taskService.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch task"})
		return
	}

	c.JSON(http.StatusOK, task)
}
