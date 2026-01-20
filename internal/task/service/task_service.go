package service

import (
	"errors"
	"strings"
	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/repository"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo repository.TaskRepositoryInterface
}

func NewService(repo repository.TaskRepositoryInterface) *Service {
	return &Service{
		repo: repo,
	}
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

func (s *Service) Edit(id string, title *string, priority *model.Priority, tags *[]string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if title != nil {
		task.Title = *title
	}

	if priority != nil {
		task.Priority = *priority
	}

	if tags != nil {
		task.Tags = *tags
	}
	now := time.Now()
	task.UpdatedAt = &now

	if err := model.ValidateTask(task); err != nil {
		return err
	}

	return s.repo.Update(task)
}

func (s *Service) List(status string, priority string, text string) ([]model.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	filtered := make([]model.Task, 0)
	for _, t := range tasks {
		if status != "" && string(t.Status) != status {
			continue
		}
		if priority != "" && string(t.Priority) != priority {
			continue
		}
		if text != "" && !strings.Contains(strings.ToLower(t.Title), text) {
			continue
		}
		for _, tag := range t.Tags {
			if strings.Contains(strings.ToLower(tag), text) {
				filtered = append(filtered, t)
				break
			}
		}
		filtered = append(filtered, t)
	}
	return filtered, nil
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
