/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-todoapp/cmd"
	"cli-todoapp/internal/config"
)

func main() {
	cmd.Execute()

	config.LoadConfig()
}
