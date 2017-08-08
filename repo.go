package main

import "fmt"

var currentID int
var todos Todos

//Get some data
func init() {
	RepoCreateTodo(Todo{Name: "Todo 1"})

	RepoCreateTodo(Todo{Name: "Todo 2"})
}

func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.Id = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo some doc
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
