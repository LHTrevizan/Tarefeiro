package main

import (
	"fmt"

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
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		return service.Delete(args[0])
	},
}
