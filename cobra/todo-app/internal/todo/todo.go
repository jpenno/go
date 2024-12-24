package todo

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func New(id int, task string, done bool) Todo {
	return Todo{
		Id:   id,
		Task: task,
		Done: done,
	}
}

func PrintTodos(todos []Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tDone")
	for _, t := range todos {
		t.Print(w)
	}
	w.Flush()
}

func (t Todo) Print(w io.Writer) {

	fmt.Fprintf(w, "%v\t", t.Id)
	fmt.Fprintf(w, "%v\t", t.Task)
	if t.Done {
		fmt.Fprintf(w, "Done\t")
	} else {
		fmt.Fprintf(w, "Not done\t")
	}
	fmt.Fprintln(w, "")
}
