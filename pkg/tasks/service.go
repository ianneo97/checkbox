package tasks

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/tasks/requests"
)

func (h Handler) AddTask(ctx *gin.Context) {
	body := requests.AddTaskRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	task := Task{
		Name:        body.Name,
		Description: body.Description,
		DueDate:     body.DueDate,
		Status:      getStatus(body.DueDate),
	}

	if result := h.DB.Create(&task); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &task)
}

func (h Handler) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &task)
}

func (h Handler) UpdateTask(ctx *gin.Context) {
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
	task.Status = getStatus(body.DueDate)

	h.DB.Save(&task)

	ctx.JSON(http.StatusOK, &task)
}

func (h Handler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task Task

	if result := h.DB.First(&task, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&task)

	ctx.Status(http.StatusOK)
}

func (h Handler) ListTasks(ctx *gin.Context) {
	name := ctx.Query("name")

	var tasks []Task

	query := h.DB.Order("created_at")

	if name != "" {
		query.Where("name LIKE ?", "%"+name+"%")
	}

	if result := query.Find(&tasks); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &tasks)
}

func getStatus(d *time.Time) string {
	status := ""
	current := time.Now()
	oneWeekAfter := current.AddDate(0, 0, 7)

	if (d.After(current) && d.Before(oneWeekAfter)) || d.Equal(oneWeekAfter) {
		status = TaskStatus.String(DUE_SOON)
	}

	if d.After(oneWeekAfter) {
		status = TaskStatus.String(NOT_URGENT)
	}

	if d.Before(current) {
		status = TaskStatus.String(OVERDUE)
	}

	return status
}
