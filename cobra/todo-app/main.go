/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-todoapp/cmd"
	"cli-todoapp/internal/config"
	"cli-todoapp/internal/todo"
)

func main() {
	config.Init()
	todo.Init()
	cmd.Execute()
}
