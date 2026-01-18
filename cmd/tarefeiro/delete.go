package main

import (
	"tarefeiro/internal/task/repository"
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
		repo, _ := repository.NewRepository("data/tasks.json")
		service := service.NewService(repo)
		return service.Delete(args[0])
	},
}
