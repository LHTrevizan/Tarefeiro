package cmd

import (
	"fmt"
	"strings"

	"tarefeiro/internal/task"

	"github.com/spf13/cobra"
)

var (
	priority string
	tags     string
)

func init() {
	addCmd.Flags().StringVar(&priority, "priority", string(task.PriorityMedium), "Prioridade: low | medium | high")
	addCmd.Flags().StringVar(&tags, "tags", "", "Tags separadas por vírgula")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add \"título\"",
	Args:  cobra.ExactArgs(1),
	Short: "Adicionar tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		service := task.NewService("data/tasks.json")

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

func parsePriority(p string) (task.Priority, error) {
	switch p {
	case string(task.PriorityLow):
		return task.PriorityLow, nil
	case string(task.PriorityMedium):
		return task.PriorityMedium, nil
	case string(task.PriorityHigh):
		return task.PriorityHigh, nil
	default:
		return "", fmt.Errorf("prioridade inválida: %s (use low, medium ou high)", p)
	}
}
