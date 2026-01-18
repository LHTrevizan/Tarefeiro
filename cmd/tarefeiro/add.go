package main

import (
	"fmt"
	"strings"

	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/repository"
	"tarefeiro/internal/task/service"

	"github.com/spf13/cobra"
)

var (
	priority string
	tags     string
)

func init() {
	addCmd.Flags().StringVar(&priority, "priority", string(model.PriorityMedium), "Prioridade: low | medium | high")
	addCmd.Flags().StringVar(&tags, "tags", "", "Tags separadas por vírgula")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add \"título\"",
	Args:  cobra.ExactArgs(1),
	Short: "Adicionar tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, _ := repository.NewRepository("data/tasks.json")
		service := service.NewService(repo)

		var tagList []string
		if tags != "" {
			tagList = strings.Split(tags, ",")
		}

		priority, err := parsePriority(priority)
		if err != nil {
			return err
		}

		return service.Add(args[0], priority, tagList)
	},
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
