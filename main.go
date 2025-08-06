package main

import (
	"fmt"
)

func main() {
	var todo string
	var option int
	cont := "y"
	var todos []string

	fmt.Println("##### Hello There! What's your plan for today #####")
	fmt.Println()
	for cont == "y" {
		fmt.Println("{1 for Adding to todo, 2 for view all todo}")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("input todo")
			fmt.Scanln(&todo)
			todos = addTask(todos, todo)
			viewTodos(todos)

		case 2:
			fmt.Println(todos)
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
	for i, e := range todos {
		fmt.Println(i+1, ".", e)
	}
}
