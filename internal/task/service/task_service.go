package service

import (
	"errors"
	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/repository"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.TaskRepository
}

func NewService(file string) (*Service, error) {
	repo, err := repository.NewRepository(file)
	if err != nil {
		return nil, err
	}
	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) Add(title string, priority model.Priority, tags []string) error {
	task := &model.Task{
		ID:        uuid.NewString(),
		Title:     title,
		Status:    model.StatusPending,
		Priority:  priority,
		Tags:      tags,
		CreatedAt: time.Now(),
	}

	if err := model.ValidateTask(task); err != nil {
		return err
	}
	return s.repo.Create(task)
}

func (s *Service) Edit(id string, title string, priority model.Priority, tags []string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if title != "" {
		task.Title = title
	}
	if priority != "" {
		task.Priority = priority
	}
	if tags != nil {
		task.Tags = tags
	}
	now := time.Now()
	task.UpdatedAt = &now

	if err := model.ValidateTask(task); err != nil {
		return err
	}

	return s.repo.Update(task)
}

func (s *Service) List() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *Service) Complete(id string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if task.Status == model.StatusCompleted {
		return errors.New("task já está concluída")
	}

	now := time.Now()
	task.UpdatedAt = &now
	task.Status = model.StatusCompleted
	task.CompletedAt = &now

	return s.repo.Update(task)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) Show(id string) (*model.Task, error) {
	return s.repo.GetByID(id)
}
