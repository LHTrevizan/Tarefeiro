package main

import (
	"tarefeiro/internal/task/service"

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
		service, err := service.NewService("data/tasks.json")
		if err != nil {
			return err
		}
		return service.Delete(args[0])
	},
}
