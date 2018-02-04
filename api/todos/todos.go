package todos

import (
	"../../model/todo"
)

// Create controller.
func Create(task string) error {
	return todo.Create(task)
}

// List controller.
func List() ([]todo.Todo, error) {
	return todo.List()
}

// Get controller.
func Get(id int) (todo.Todo, error) {
	return todo.Get(id)
}

// Update controller.
func Update(id int, task string) error {
	return todo.Update(id, task)
}

// Delete controller.
func Delete(id int) error {
	return todo.Delete(id)
}
