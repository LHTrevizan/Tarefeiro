package repository

import (
	"errors"
	"tarefeiro/internal/task/infra"
	"tarefeiro/internal/task/model"
)

type TaskRepository struct {
	storage *infra.JSONStorage
}

func NewRepository(file string) (*TaskRepository, error) {
	storage, err := infra.NewJSONStorage(file)
	if err != nil {
		return nil, err
	}
	return &TaskRepository{storage: storage}, nil
}
func (r *TaskRepository) Create(task *model.Task) error {
	tasks, err := r.storage.Load()
	if err != nil {
		return err
	}
	tasks = append(tasks, *task)
	return r.storage.Save(tasks)
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	return r.storage.Load()
}

func (r *TaskRepository) GetByID(id string) (*model.Task, error) {
	tasks, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepository) Update(updatedTask *model.Task) error {
	tasks, err := r.GetAll()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == updatedTask.ID {
			tasks[i] = *updatedTask
			return r.storage.Save(tasks)
		}
	}
	return errors.New("task not found")
}

func (r *TaskRepository) Delete(id string) error {
	tasks, err := r.GetAll()
	if err != nil {
		return err
	}
	found := false
	newTasks := make([]model.Task, 0)
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("task not found")
	}
	return r.storage.Save(newTasks)
}
