package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	filterStatus   string
	filterPriority string
)

func init() {
	listCmd.Flags().StringVarP(&filterStatus, "status", "s", "", "Filtrar tarefas por status (pending, done)")
	listCmd.Flags().StringVarP(&filterPriority, "priority", "p", "", "Filtrar tarefas por prioridade (low, medium, high)")

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar tarefas, tambem tem opção de filtro por status e prioridade",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		tasks, err := service.List(strings.ToLower(filterStatus), strings.ToLower(filterPriority))

		if err != nil {
			return err
		}

		return RunListInteractive(service, tasks)
	},
}
