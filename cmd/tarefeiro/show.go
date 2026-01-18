package main

import (
	"fmt"
	"tarefeiro/internal/task/model"
	"tarefeiro/internal/task/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Exibir detalhes de uma tarefa",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := service.NewService("data/tasks.json")
		if err != nil {
			return err
		}
		t, err := service.Show(args[0])
		if err != nil {
			return err
		}

		printTask(t)
		return nil
	},
}

func printTask(t *model.Task) {
	fmt.Println("──────────────")
	fmt.Printf("ID: %s\n", t.ID)
	fmt.Printf("Título: %s\n", t.Title)
	fmt.Printf("Status: %s\n", t.Status)
	fmt.Printf("Prioridade: %s\n", t.Priority)

	if len(t.Tags) > 0 {
		fmt.Printf("Tags: %v\n", t.Tags)
	}

	fmt.Printf("Criada em: %s\n", t.CreatedAt.Format("02/01/2006 15:04"))

	if t.CompletedAt != nil {
		fmt.Printf("Concluída em: %s\n",
			t.CompletedAt.Format("02/01/2006 15:04"))
	}

	fmt.Println("──────────────")
}
