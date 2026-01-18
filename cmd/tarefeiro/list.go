package main

import (
	"fmt"

	"tarefeiro/internal/task/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar tarefas",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := service.NewService("data/tasks.json")
		if err != nil {
			return err
		}
		tasks, err := service.List()
		if err != nil {
			return err
		}
		for _, t := range tasks {
			fmt.Printf("[%d] %s | %s | %s | %v\n",
				t.ID, t.Title, t.Status, t.Priority, tags)
		}
		return nil
	},
}
