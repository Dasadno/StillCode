package handlers

import (
	"StillCode/server/internal/models"
	"StillCode/server/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SubmissionHandler handles code execution HTTP requests
type SubmissionHandler struct {
	submissionService *services.SubmissionService
}

// NewSubmissionHandler creates a new SubmissionHandler
func NewSubmissionHandler(submissionService *services.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

// RunCode handles POST /api/run
func (h *SubmissionHandler) RunCode(c *gin.Context) {
	var req models.RunCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request format"})
		return
	}

	result, err := h.submissionService.RunCode(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// SubmitSolution handles POST /api/submit/:id
func (h *SubmissionHandler) SubmitSolution(c *gin.Context) {
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var req models.SubmitSolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request format"})
		return
	}

	result, err := h.submissionService.SubmitSolution(c.Request.Context(), taskID, &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrTaskNotFound):
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Task not found"})
		case errors.Is(err, services.ErrNoTestCases):
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "No test cases found for this task"})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
