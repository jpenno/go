package todo

import (
	"cli-todoapp/internal/config"
	fileio "cli-todoapp/internal/fileIO"
	"errors"
	"fmt"
)

var (
	todos []Todo
)

func Init() {
	todos = fileio.GetJson(config.C.Path, todos)
	syncIds()
}

func AddTask(task string) {
	todo := Todo{len(todos) + 1, task, false}
	todos = append(todos, todo)
	fileio.SaveJson(config.C.Path, todos)
}

func DeleteTodo(id int) {
	todos = remove(todos, id-1)
	fileio.SaveJson(config.C.Path, todos)
}

func GetTodos() []Todo {
	return todos
}

func GetById(id int) (*Todo, error) {
	for i := range todos {
		if todos[i].Id == id {
			return &todos[i], nil
		}
	}
	msg := fmt.Sprintf("Could not find id: %v", id)
	return nil, errors.New(msg)
}

func Complete(id int) {
	completeTodo, _ := GetById(id)
	completeTodo.Done = !completeTodo.Done
	fileio.SaveJson(config.C.Path, todos)
}

func syncIds() {
	for i := range todos {
		todos[i].Id = i + 1
	}
}

func remove[t any](slice []t, s int) []t {
	return append(slice[:s], slice[s+1:]...)
}
