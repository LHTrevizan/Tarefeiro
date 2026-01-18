package main

import (
	"tarefeiro/internal/task/repository"
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
		repo, _ := repository.NewRepository("data/tasks.json")
		service := service.NewService(repo)
		return service.Complete(args[0])
	},
}
