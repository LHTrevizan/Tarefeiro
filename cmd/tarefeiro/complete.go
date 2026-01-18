package main

import (
	"tarefeiro/internal/task/service"

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
		service, err := service.NewService("data/tasks.json")
		if err != nil {
			return err
		}
		return service.Complete(args[0])
	},
}
