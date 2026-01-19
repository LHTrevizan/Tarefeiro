package main

import (
	"fmt"
	"strings"

	"tarefeiro/internal/task/model"

	"github.com/spf13/cobra"
)

var (
	priority string
	tags     string
)

func init() {
	addCmd.Flags().StringVar(&priority, "priority", string(model.PriorityMedium), "Prioridade: low | medium | high")
	addCmd.Flags().StringVar(&tags, "tags", "", "Tags separadas por vírgula")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add \"título\"",
	Args:  cobra.ExactArgs(1),
	Short: "Adicionar tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		var tagList []string
		if tags != "" {
			tagList = strings.Split(tags, ",")
		}

		priority, err := parsePriority(priority)
		if err != nil {
			return err
		}

		return service.Add(args[0], priority, tagList)
	},
}
