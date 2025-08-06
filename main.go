package main

import (
	"fmt"
)

func main() {

	shortGolang := "Watch Go crash course"
	fullGolang := "Watch Nana Golang Full Course"
	rewardDesert := "Reward mysefl ith a donut"

	taskItems := []string{
		shortGolang,
		fullGolang,
		rewardDesert,
	}

	fmt.Println("##### Welcome to our Todolist App! ######")
	fmt.Println()
	fmt.Println("List of my Todos")
	fmt.Println()

	for i,e := range taskItems {
		fmt.Println(i+1,". ", e)
	}

	

}
