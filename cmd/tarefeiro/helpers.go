package main

import (
	"fmt"
	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/repository"
	"tarefeiro/internal/task/service"
)

func InitService() (*service.Service, error) {
	repo, err := repository.NewRepository("../../data/tasks.json")
	if err != nil {
		return nil, err
	}
	service := service.NewService(repo)
	return service, nil
}

func parsePriority(p string) (model.Priority, error) {
	switch p {
	case string(model.PriorityLow):
		return model.PriorityLow, nil
	case string(model.PriorityMedium):
		return model.PriorityMedium, nil
	case string(model.PriorityHigh):
		return model.PriorityHigh, nil
	default:
		return "", fmt.Errorf("prioridade inválida: %s (use low, medium ou high)", p)
	}
}
