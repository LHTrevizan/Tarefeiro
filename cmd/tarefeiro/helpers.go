package main

import (
	"encoding/json"
	"fmt"
	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/repository"
	"tarefeiro/internal/task/service"

	"github.com/manifoldco/promptui"
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

func RunListInteractive(service *service.Service, tasks []model.Task) error {
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada")
		return nil
	}

	items := make([]string, len(tasks))
	for i, t := range tasks {
		items[i] = fmt.Sprintf("[%s] %-15s | %-8s | %-6s", t.ID, t.Title, t.Status, t.Priority)
	}

	prompt := promptui.Select{
		Label: "Selecione uma tarefa",
		Items: items,
		Size:  10,
	}

	idx, _, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			fmt.Println("Saindo...")
			return nil
		}
		return err
	}

	selected := tasks[idx]

	data, err := json.MarshalIndent(selected, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}
