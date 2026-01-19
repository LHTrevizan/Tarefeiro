package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Gerenciador de tarefas simples",
	Long: `Tarefeiro é um gerenciador de tarefas simples de linha de comando com funcionalidades básicas
como adicionar, listar, concluir, exibir e remover tarefas.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
