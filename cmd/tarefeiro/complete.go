package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Args:  cobra.ExactArgs(1),
	Short: "Concluir tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		service, err := InitService()
		if err != nil {
			return fmt.Errorf("Erro ao inicializar service %s\n", err)
		}
		return service.Complete(args[0])
	},
}
