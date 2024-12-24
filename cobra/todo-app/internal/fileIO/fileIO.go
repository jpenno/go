package fileio

import (
	"cli-todoapp/internal/todo"
	"encoding/json"
	"os"
)

func GetTodos(path string) []todo.Todo {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var todos []todo.Todo
	json.Unmarshal(fileData, &todos)

	return todos
}

func SaveTodos(path string, todos []todo.Todo) {
	jsonString, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.WriteString(string(jsonString))

	// os.WriteFile(path, jsonString, os.ModePerm)
}
