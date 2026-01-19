package main

import (
	"fmt"
	"strings"
	"tarefeiro/internal/task/model"

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
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}

		var (
			titlePtr    *string
			priorityPtr *model.Priority
			tagsPtr     *[]string
		)

		if editTitle != "" {
			titlePtr = &editTitle
		}

		if editPriority != "" {
			p, err := parsePriority(editPriority)
			if err != nil {
				return err
			}
			priorityPtr = &p
		}

		if editTags != "" {
			t := strings.Split(editTags, ",")
			tagsPtr = &t
		}

		return service.Edit(
			args[0],
			titlePtr,
			priorityPtr,
			tagsPtr,
		)
	}}
