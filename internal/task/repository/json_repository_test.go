package repository

import (
	"path/filepath"
	"tarefeiro/internal/task/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestRepo(t *testing.T) *TaskRepository {
	t.Helper()

	dir := t.TempDir()
	file := filepath.Join(dir, "tasks.json")

	repo, err := NewRepository(file)
	assert.NoError(t, err)

	return repo
}

func sampleTask(id string) *model.Task {
	return &model.Task{
		ID:       id,
		Title:    "Test task",
		Status:   model.StatusPending,
		Priority: model.PriorityMedium,
	}
}

func TestRepository_Create_Success(t *testing.T) {
	repo := newTestRepo(t)

	err := repo.Create(sampleTask("1"))
	assert.NoError(t, err)

	tasks, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "1", tasks[0].ID)
}

func TestRepository_GetAll_Empty(t *testing.T) {
	repo := newTestRepo(t)

	tasks, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Empty(t, tasks)
}

func TestRepository_GetByID_Success(t *testing.T) {
	repo := newTestRepo(t)
	_ = repo.Create(sampleTask("abc"))

	task, err := repo.GetByID("abc")
	assert.NoError(t, err)
	assert.Equal(t, "abc", task.ID)
}

func TestRepository_GetByID_NotFound(t *testing.T) {
	repo := newTestRepo(t)

	task, err := repo.GetByID("404")
	assert.Error(t, err)
	assert.Nil(t, task)
}

func TestRepository_Update_Success(t *testing.T) {
	repo := newTestRepo(t)

	task := sampleTask("1")
	_ = repo.Create(task)

	task.Title = "Updated title"
	err := repo.Update(task)
	assert.NoError(t, err)

	updated, _ := repo.GetByID("1")
	assert.Equal(t, "Updated title", updated.Title)
}

func TestRepository_Update_NotFound(t *testing.T) {
	repo := newTestRepo(t)

	err := repo.Update(sampleTask("404"))
	assert.Error(t, err)
}

func TestRepository_Delete_Success(t *testing.T) {
	repo := newTestRepo(t)

	_ = repo.Create(sampleTask("1"))

	err := repo.Delete("1")
	assert.NoError(t, err)

	tasks, _ := repo.GetAll()
	assert.Len(t, tasks, 0)
}

func TestRepository_Delete_NotFound(t *testing.T) {
	repo := newTestRepo(t)

	err := repo.Delete("404")
	assert.Error(t, err)
}
