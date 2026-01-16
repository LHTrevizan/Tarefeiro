// internal/task/service.go
package task

import (
	"errors"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(file string) *Service {
	return &Service{
		repo: &Repository{File: file},
	}
}

func (s *Service) Add(title string, priority Priority, tags []string) error {
	if title == "" {
		return errors.New("título não pode ser vazio")
	}

	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:        id,
		Title:     title,
		Status:    "pending",
		Priority:  priority,
		Tags:      tags,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, task)
	return s.repo.Save(tasks)
}

func (s *Service) Edit(id int, title string, priority Priority, tags []string) error {
	if title == "" {
		return errors.New("título não pode ser vazio")
	}

	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if title != "" {
				tasks[i].Title = title
			}

			if priority != "" {
				tasks[i].Priority = Priority(priority)
			}

			if tags != nil {
				tasks[i].Tags = tags
			}
			return s.repo.Save(tasks)
		}
	}

	return errors.New("tarefa não encontrada")
}

func (s *Service) List() ([]Task, error) {
	return s.repo.Load()
}

func (s *Service) Complete(id int) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "done"
			now := time.Now()
			tasks[i].CompletedAt = &now
			return s.repo.Save(tasks)
		}
	}

	return errors.New("tarefa não encontrada")
}

func (s *Service) Delete(id int) error {
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	var result []Task
	found := false

	for _, t := range tasks {
		if t.ID != id {
			result = append(result, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("tarefa não encontrada")
	}

	return s.repo.Save(result)
}
