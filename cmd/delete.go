// cmd/delete.go
package cmd

import (
	"strconv"

	"tarefeiro/internal/task"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Args:  cobra.ExactArgs(1),
	Short: "Remover tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		service := task.NewService("data/tasks.json")
		return service.Delete(id)
	},
}
