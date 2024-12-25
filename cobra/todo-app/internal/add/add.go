package add

import (
	"cli-todoapp/internal/config"
	fileio "cli-todoapp/internal/fileIO"
	"cli-todoapp/internal/todo"
)

func Add(task string) {
	var todos []todo.Todo
	todos = fileio.GetJson(config.C.Path, todos)
	newTodo := todo.New(len(todos), task, false)
	todos = append(todos, newTodo)
	todos = setTodoIds(todos)

	fileio.SaveJson(config.C.Path, todos)
}
func setTodoIds(todos []todo.Todo) []todo.Todo {
	for i := range todos {
		todos[i].Id = i + 1
	}
	return todos
}
