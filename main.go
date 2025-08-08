package main

import (
	"fmt"
)

func main() {
	var todo string
	var option int
	cont := "y"
	var todos []string
	var mark int
	fmt.Println("##### Hello There! What's your plan for today #####")
	fmt.Println()
	for cont == "y" {
		fmt.Println("{1 for Adding to todo, 2 for view all todo, 3 for marking as complete, 4 for deleteing todo}")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("input todo")
			fmt.Scanln(&todo)
			todos = addTask(todos, todo)
			viewTodos(todos)

		case 2:
			fmt.Println("Here are your todos")
			fmt.Println()
			fmt.Println(todos)
			viewTodos(todos)

		case 3:
			fmt.Println("Choose the todo you want to mark as complete")
			fmt.Scanln(&mark)
			todos = completed(mark, todos)
			viewTodos(todos)
		case 4:
			fmt.Println("Choose the todo you want to delete")
			fmt.Scanln(&mark)
			todos = deleted(mark, todos)
			viewTodos(todos)

		default:
			fmt.Println("Invalid Option, try again")

		}
		fmt.Print("Do you want to perform another opertaion")
		fmt.Println()
		fmt.Println("{y/n}")
		fmt.Scanln(&cont)

	}

}

func addTask(todos []string, todo string) []string {
	todos = append(todos, todo)
	return todos
}

func viewTodos(todos []string) {
	if len(todos) == 0 {
		fmt.Println("No todos yet!")
		return
	}

	for i, e := range todos {
		fmt.Println(i+1, ".", e)
	}

}

func completed(index int, todos []string) []string {
	index = index - 1
	if index >= 0 && index < len(todos) {
		todos[index] = "[x] " + todos[index]
		fmt.Println("Marked as completed:", todos[index])
	} else {
		fmt.Println("Invalid task number.")
	}
	return todos
}

func deleted(index int, todos []string) []string {
	index = index - 1
	if index >= 0 && index < len(todos) {
		todos = append(todos[:index], todos[index+1:]...)
		fmt.Println("Successfully deleted:", todos[index])
		

	} else {
		fmt.Println("Invalid task number.")
	}
	return todos
}
