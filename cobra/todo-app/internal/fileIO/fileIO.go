package fileio

import (
	"cli-todoapp/internal/todo"
	"encoding/json"
	"os"
)

func GetJson[T any](path string, t T) T {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(fileData, &t)

	return t
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
