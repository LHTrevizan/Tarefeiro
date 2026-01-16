package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"tarefeiro/internal/task"

	"github.com/spf13/cobra"
)

var (
	editTitle    string
	editPriority string
	editTags     string
)

func init() {
	editCmd.Flags().StringVar(&editTitle, "title", "", "Novo título")
	editCmd.Flags().StringVar(
		&editPriority,
		"priority",
		"",
		"Prioridade: low | medium | high",
	)
	editCmd.Flags().StringVar(&editTags, "tags", "", "Tags separadas por vírgula")

	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Args:  cobra.ExactArgs(1),
	Short: "Editar tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id inválido")
		}

		service := task.NewService("data/tasks.json")

		var tags []string
		if editTags != "" {
			tags = strings.Split(editTags, ",")
		}
		editPriority, err := parsePriority(editPriority)
		if err != nil {
			return err
		}

		return service.Edit(
			id,
			editTitle,
			editPriority,
			tags,
		)
	},
}
