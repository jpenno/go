/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todoapp/internal/todo"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task by id",
	Long: `delete task by id 
	eg:
	delete 1`,
	Run: deleteTodoCmd,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTodoCmd(cmd *cobra.Command, args []string) {
	id, _ := strconv.Atoi(args[0])
	td, _ := todo.GetById(id)
	fmt.Printf("Delete: %q\n", td.Task)
	todo.DeleteTodo(id)
}
