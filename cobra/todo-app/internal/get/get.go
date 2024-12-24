package get

import (
	"cli-todoapp/internal/config"
	"cli-todoapp/internal/fileIO"
	"cli-todoapp/internal/todo"
)

func GetTodos() []todo.Todo {
	var todos []todo.Todo
	todos = fileio.GetJson(config.C.Path, todos)
	return todos
}

func GetTodoById(id int) todo.Todo {
	todos := GetTodos()
	return todos[id-1]
}
