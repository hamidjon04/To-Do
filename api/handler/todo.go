package handler

import (
	"net/http"
	"strconv"
	"todo/models"

	"github.com/gin-gonic/gin"
)

// @Description Create a new ToDo with title, description, and user ID
// @Tags ToDo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body models.CreateToDoReq true "Create ToDo"
// @Success 201 {object} models.Todo
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /todos [post]
func (h *Handler) CreateTodo(c *gin.Context) {
	var req models.CreateToDoReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo, err := h.Storage.CreateTodo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// @Description Get all ToDos for a specific user by user ID
// @Tags ToDo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query int true "User ID"
// @Success 200 {array} models.Todo
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /todos [get]
func (h *Handler) GetTodos(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	todos, err := h.Storage.GetTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// @Description Update the title, description, and completion status of a ToDo
// @Tags ToDo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body models.UpdateToDoReq true "Update ToDo"
// @Success 200 {object} string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /todos [put]
func (h *Handler) UpdateTodo(c *gin.Context) {
	var req models.UpdateToDoReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Storage.UpdateTodo(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

// @Description Delete a ToDo by ID and User ID
// @Tags ToDo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body models.DeleteToDoReq true "Delete ToDo"
// @Success 200 {object} string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /todos [delete]
func (h *Handler) DeleteTodo(c *gin.Context) {
	var req models.DeleteToDoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Storage.DeleteTodo(req.ID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
