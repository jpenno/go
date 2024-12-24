/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todoapp/internal/get"
	"cli-todoapp/internal/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all todos",
	Long:  `List all todos`,
	Run:   listTodoCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTodoCmd(cmd *cobra.Command, args []string) {
	todos := get.GetTodos()
	todo.PrintTodos(todos)
}
