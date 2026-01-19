package service_test

import (
	"errors"
	"testing"

	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/service"

	"github.com/stretchr/testify/assert"
)

type FakeRepo struct {
	tasks map[string]*model.Task
}

func NewFakeRepo() *FakeRepo {
	return &FakeRepo{tasks: make(map[string]*model.Task)}
}

func (r *FakeRepo) Create(t *model.Task) error {
	if t.Priority != model.PriorityLow && t.Priority != model.PriorityMedium && t.Priority != model.PriorityHigh {
		return errors.New("priority inválida")
	}
	r.tasks[t.ID] = t
	return nil
}

func (r *FakeRepo) GetAll() ([]model.Task, error) {
	tasks := []model.Task{}
	for _, t := range r.tasks {
		tasks = append(tasks, *t)
	}
	return tasks, nil
}

func (r *FakeRepo) GetByID(id string) (*model.Task, error) {
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return t, nil
}

func (r *FakeRepo) Update(t *model.Task) error {
	_, ok := r.tasks[t.ID]
	if !ok {
		return errors.New("task not found")
	}
	r.tasks[t.ID] = t
	return nil
}

func (r *FakeRepo) Delete(id string) error {
	_, ok := r.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

func TestService_Add_Success(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	err := svc.Add("task 1", model.PriorityMedium, nil)
	assert.NoError(t, err)

	tasks, _ := svc.List("", "")
	assert.Len(t, tasks, 1)
	assert.Equal(t, "task 1", tasks[0].Title)
	assert.Equal(t, model.StatusPending, tasks[0].Status)
}

func TestService_Add_Failed(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	err := svc.Add("task 2", "invalid", nil)
	assert.Error(t, err)
}

func TestService_Edit_Success(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	_ = svc.Add("old title", model.PriorityMedium, nil)
	var id string
	for _, t := range repo.tasks {
		id = t.ID
	}
	newTitle := "new title"
	priority := model.PriorityHigh
	err := svc.Edit(id, &newTitle, &priority, &[]string{"tag1"})
	assert.NoError(t, err)

	task, _ := svc.Show(id)
	assert.Equal(t, "new title", task.Title)
	assert.Equal(t, model.PriorityHigh, task.Priority)
	assert.Equal(t, []string{"tag1"}, task.Tags)
}

func TestService_Edit_Failed(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)
	newTitle := "new title"
	priority := model.PriorityLow

	err := svc.Edit("notfound", &newTitle, &priority, nil)
	assert.Error(t, err)
}

func TestService_Complete_Success(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	_ = svc.Add("task", model.PriorityMedium, nil)
	var id string
	for _, t := range repo.tasks {
		id = t.ID
	}

	err := svc.Complete(id)
	assert.NoError(t, err)

	task, _ := svc.Show(id)
	assert.Equal(t, model.StatusCompleted, task.Status)
	assert.NotNil(t, task.CompletedAt)
}

func TestService_Complete_Failed(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	err := svc.Complete("notfound")
	assert.Error(t, err)
}

func TestService_Delete_Success(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	_ = svc.Add("task", model.PriorityMedium, nil)
	var id string
	for _, t := range repo.tasks {
		id = t.ID
	}

	err := svc.Delete(id)
	assert.NoError(t, err)
	assert.Empty(t, repo.tasks)
}

func TestService_Delete_Failed(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	err := svc.Delete("notfound")
	assert.Error(t, err)
}

func TestService_Show_Success(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	_ = svc.Add("task", model.PriorityMedium, nil)
	var id string
	for _, t := range repo.tasks {
		id = t.ID
	}

	task, err := svc.Show(id)
	assert.NoError(t, err)
	assert.Equal(t, "task", task.Title)
}

func TestService_Show_Failed(t *testing.T) {
	repo := NewFakeRepo()
	svc := service.NewService(repo)

	_, err := svc.Show("notfound")
	assert.Error(t, err)
}
