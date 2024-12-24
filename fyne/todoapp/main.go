package main

import (
	"fmt"
	"todoapp/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("todo-app")

	w.Resize(fyne.NewSize(300, 400))

	data := []models.Todo{
		models.NewTodo("Some stuff"),
		models.NewTodo("Some more stuff"),
		models.NewTodo("Some other things"),
	}

	data[0].Done = true

	todos := binding.NewUntypedList()
	for _, t := range data {
		todos.Append(t)
	}

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."

	addBtn := widget.NewButton("Add", func() {
		fmt.Printf("Add: %q\n", newtodoDescTxt.Text)
		if newtodoDescTxt.Text != "" {
			todos.Append(models.NewTodo(newtodoDescTxt.Text))
		}
		newtodoDescTxt.Text = ""
		newtodoDescTxt.Refresh()
	})
	addBtn.Disable()

	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // LEFT
				// RIGHT ↓
				addBtn,
				// take the rest of the space
				newtodoDescTxt,
			),
			nil, // Right
			nil, // Left

			// the rest will take all the rest of the space
			widget.NewListWithData(
				// the bniding.List type
				todos,
				// func that returns the component structure of the List Item
				// exactly the same as the Simple List
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil, nil,
						// Left of the border
						widget.NewCheck("", func(b bool) {
							fmt.Println("------------------------------------------")
							tmp, _ := todos.Get()
							for _, t := range tmp {
								fmt.Println(t)
							}
							fmt.Println("------------------------------------------")
						}),
						// takes the rest of the spece
						widget.NewLabel(""),
					)
				},
				// func that is called for each item in the list and allows
				// you to show the content on the previously defined ui structure
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)
					// ideally we should check `ok` for each one of those casting
					// but we know that they are those types for sure
					l := ctr.Objects[0].(*widget.Label)
					c := ctr.Objects[1].(*widget.Check)
					todo := models.NewTodoFromDataItem(di)

					l.SetText(todo.Description)
					c.SetChecked(todo.Done)
				}),
			// container.NewCenter(),
		),
	)
	w.ShowAndRun()
}
