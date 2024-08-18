package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/domain"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEchoHandler(t *testing.T) {
	e := NewRouter()

	req := httptest.NewRequest("GET", "/", nil) // Request
	rec := httptest.NewRecorder()               // ResponseWriter

	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}

func TestTasksShowHandler(t *testing.T) {
	e := NewRouter()

	req := httptest.NewRequest("GET", "/tasks", nil) // Request
	rec := httptest.NewRecorder()                    // ResponseWriter

	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusOK, rec.Code)

	// Response bodyをparseして検証する
	var tasks []domain.Task
	err := json.Unmarshal(rec.Body.Bytes(), &tasks)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(tasks))
	assert.Equal(t, "task1", tasks[0].Title)
	assert.Equal(t, "task2", tasks[1].Title)
	assert.NotEmpty(t, tasks[0].ID)
	assert.NotEmpty(t, tasks[1].ID)
	assert.NotEmpty(t, tasks[0].CreatedAt)
	assert.NotEmpty(t, tasks[1].CreatedAt)
	assert.Empty(t, tasks[0].UpdatedAt)
	assert.Empty(t, tasks[1].UpdatedAt)
}

func TestTaskCreateHander(t *testing.T) {
	e := NewRouter()

	task := `{"title":"task3"}`

	req := httptest.NewRequest("POST", "/tasks", strings.NewReader(task)) // Request
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)      // Content-Type: application/json

	rec := httptest.NewRecorder() // ResponseWriter

	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Response bodyをparseして検証する
	var createdTask domain.Task
	err := json.Unmarshal(rec.Body.Bytes(), &createdTask)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "task3", createdTask.Title)
	assert.NotEmpty(t, createdTask.ID)
	assert.NotEmpty(t, createdTask.CreatedAt)
	assert.Empty(t, createdTask.UpdatedAt)
}
