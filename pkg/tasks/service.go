package tasks

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/tasks/requests"
)

func (h handler) AddTask(ctx *gin.Context) {
	body := requests.AddTaskRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	current := time.Now()
	status := "pending"

	if body.DueDate.After(current) {
		if body.DueDate.Before(current.AddDate(0, 0, 7)) {
			status = TaskStatus.String(DUE_SOON)
		} else {
			status = TaskStatus.String(NOT_URGENT)
		}
	}

	if body.DueDate.Before(current) {
		status = TaskStatus.String(OVERDUE)
	}

	task := Task{
		Name:        body.Name,
		Description: body.Description,
		DueDate:     body.DueDate,
		Status:      status,
	}

	if result := h.DB.Create(&task); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &body)
}

func (h handler) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &task)
}

func (h handler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	body := requests.UpdateTasksRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Name = body.Name
	task.Description = body.Description
	task.DueDate = body.DueDate
	task.Status = body.Status

	h.DB.Save(&task)

	ctx.JSON(http.StatusOK, &task)
}

func (h handler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&task)

	ctx.Status(http.StatusOK)
}

func (h handler) ListTasks(ctx *gin.Context) {
	var tasks []Task

	if result := h.DB.Order("created_at").Find(&tasks); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &tasks)
}
