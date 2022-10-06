package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/tasks/requests"
)

func (h handler) AddTask(ctx *gin.Context) {
	body := requests.AddTaskRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task Task

	task.Name = body.Name
	task.Description = body.Description
	task.DueDate = body.DueDate
	task.Status = body.Status

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
	var book Task

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Name = body.Name
	book.Description = body.Description
	book.DueDate = body.DueDate
	book.Status = body.Status

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
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

	if result := h.DB.Find(&tasks); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &tasks)
}
