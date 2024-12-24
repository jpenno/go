package models

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{Description: description, Done: false}
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}

func NewTodoFromDataItem(item binding.DataItem) Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(Todo)
}
