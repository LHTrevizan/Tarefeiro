package main

import (
	"encoding/json"
	"fmt"

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
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		t, err := service.Show(args[0])
		if err != nil {
			return err
		}
		data, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	},
}
