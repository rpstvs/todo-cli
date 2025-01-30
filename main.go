package main

import (
	"fmt"
)

func main() {

	todos := Todos{}

	todos.add("add cenas")
	todos.add("mais cenas")
	todos.add("por fim mais")
	fmt.Println(todos)
	todos.markDone(0)
	todos.remove(1)
	fmt.Println(todos)
}
