package tasks_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/config/db"
	"github.com/ianneo97/checkbox/pkg/tasks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var id uint

func TestListTasks(t *testing.T) {
	r := setupInitialBase()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/all", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddTask(t *testing.T) {
	r := setupInitialBase()
	w := httptest.NewRecorder()

	// This is to signify that it is a due soon task
	ct := time.Now().AddDate(0, 0, 1)

	task := tasks.Task{
		Name:        "Test Task",
		Description: "This is a test task",
		DueDate:     &ct,
	}

	jsonValue, _ := json.Marshal(task)

	req, _ := http.NewRequest(http.MethodPost, "/tasks/", bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)

	res := tasks.Task{}
	json.Unmarshal([]byte(w.Body.Bytes()), &res)

	id = res.ID

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, task.Name, res.Name)
	assert.Equal(t, task.Description, res.Description)
	assert.Equal(t, tasks.TaskStatus.String(tasks.DUE_SOON), res.Status)
}

func TestUpdateTask(t *testing.T) {
	r := setupInitialBase()
	w := httptest.NewRecorder()

	// This is to signify the task is not urgent
	ct := time.Now().AddDate(0, 0, 8)
	task := tasks.Task{
		Name:        "Test Task",
		Description: "This is a test task",
		DueDate:     &ct,
	}

	jsonValue, _ := json.Marshal(task)

	req, _ := http.NewRequest(http.MethodPatch, "/tasks/"+strconv.FormatUint(uint64(id), 10), bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)

	res := tasks.Task{}
	json.Unmarshal([]byte(w.Body.Bytes()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, task.Name, res.Name)
	assert.Equal(t, task.Description, res.Description)
	assert.Equal(t, tasks.TaskStatus.String(tasks.NOT_URGENT), res.Status)
}

func TestDeleteTask(t *testing.T) {
	r := setupInitialBase()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+strconv.FormatUint(uint64(id), 10), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func setupInitialBase() *gin.Engine {
	viper.SetConfigFile("../pkg/config/envs/.env")
	viper.ReadInConfig()

	db := db.Init()
	r := gin.New()

	tasks.RegisterRoutes(r, db)

	return r
}
