package repository

import "tarefeiro/internal/task/model"

type TaskRepositoryInterface interface {
	Create(*model.Task) error
	GetByID(string) (*model.Task, error)
	GetAll() ([]model.Task, error)
	Update(*model.Task) error
	Delete(string) error
}
