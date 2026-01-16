// cmd/complete.go
package cmd

import (
	"strconv"

	"tarefeiro/internal/task"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Args:  cobra.ExactArgs(1),
	Short: "Concluir tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := strconv.Atoi(args[0])
		service := task.NewService("data/tasks.json")
		return service.Complete(id)
	},
}
