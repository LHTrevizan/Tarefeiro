package main

import (
	"encoding/json"
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
	Short: "Listar tarefas",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		tasks, err := service.List(strings.ToLower(filterStatus), strings.ToLower(filterPriority))

		if err != nil {
			return err
		}
		output := make([]struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Status   string `json:"status"`
			Priority string `json:"priority"`
		}, len(tasks))

		for i, t := range tasks {
			output[i] = struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				Status   string `json:"status"`
				Priority string `json:"priority"`
			}{
				ID:       t.ID,
				Title:    t.Title,
				Status:   string(t.Status),
				Priority: string(t.Priority),
			}
		}
		data, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	},
}
