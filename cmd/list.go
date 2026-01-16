// cmd/list.go
package cmd

import (
	"fmt"

	"tarefeiro/internal/task"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar tarefas",
	RunE: func(cmd *cobra.Command, args []string) error {
		service := task.NewService("data/tasks.json")
		tasks, err := service.List()
		if err != nil {
			return err
		}

		for _, t := range tasks {
			fmt.Printf("[%d] %s | %s | %s\n",
				t.ID, t.Title, t.Status, t.Priority)
		}
		return nil
	},
}

// test
