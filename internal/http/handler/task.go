package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	tasksentity "todolist/internal/entity/tasks"
	taskusecase "todolist/internal/usecase/task"
)

type Task struct {
	task taskusecase.TaskUseCaseIml
}

func NewTask(task taskusecase.TaskUseCaseIml) *Task {
	return &Task{task: task}
}

// CreateTask godoc
// @Summary CreateTask
// @Description CreateTask  users
// @Tags task
// @Accept json
// @Produce json
// @Param body body tasksentity.CreateTaskReq true " "
// @Security ApiKeyAuth
// @Success 201 {object} tasksentity.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/create [post]
func (t *Task) CreateTask(c *gin.Context) {
	var req tasksentity.CreateTaskReq // Инициализируем структуру

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	req.UserId = id.(string)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.task.CreateTask(c.Request.Context(), &req) // Передаём указатель на структуру
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary UpdateTask
// @Description update task
// @Tags task
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Param body body tasksentity.UpdateTaskReq true " "
// @Security ApiKeyAuth
// @Success 201 {object} tasksentity.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/update/{task_id} [patch]
func (t *Task) UpdateTask(c *gin.Context) {
	taskID := c.Param("task_id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}

	var req tasksentity.UpdateTaskReq
	req.UserId = id.(string)
	req.Id = taskID

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.task.UpdateTask(c.Request.Context(), &req) // Передаём указатель на структуру
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetTask godoc
// @Summary GetTask
// @Description GetTask task
// @Tags task
// @Accept json
// @Produce json
// @Param field path string true "field"
// @Param value path string true "value"
// @Security ApiKeyAuth
// @Success 201 {object} tasksentity.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/{field}/{value} [get]
func (t *Task) GetTask(c *gin.Context) {
	field := c.Param("field")
	if field == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}
	value := c.Param("value")
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task value required"})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	var req tasksentity.GetTaskReq
	req.UserId = id.(string)
	req.Field = field
	req.Value = value
	task, err := t.task.GetTask(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// GetTasks godoc
// @Summary GetTasks
// @Description GetTasks tasks
// @Tags task
// @Accept json
// @Produce json
// @Param field header string false "Field"
// @Param value header string false "Value"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Security ApiKeyAuth
// @Success 201 {object} []tasksentity.Task
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/tasks   [get]
func (t *Task) GetTasks(c *gin.Context) {
	var req tasksentity.GetAllTaskReq

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	req.UserId = id.(string)
	req.Field = c.GetHeader("Field")
	req.Value = c.GetHeader("Value")

	offsetStr := c.Query("offset")
	if offsetStr == "" {
		req.Offset = 0
	} else {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		req.Offset = offset
	}
	limitStr := c.Query("limit")
	if limitStr == "" {
		req.Limit = 0
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		req.Limit = limit
	}
	tasks, err := t.task.GetALlTask(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// DeleteTask godoc
// @Summary DeleteTask
// @Description DeleteTask tasks
// @Tags task
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /task/delete/{task_id}   [delete]
func (t *Task) DeleteTask(c *gin.Context) {
	id := c.Param("task_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	err := t.task.DeleteTask(c.Request.Context(), userID.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
