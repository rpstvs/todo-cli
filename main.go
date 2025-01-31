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
	todos.edit(1, "decidi mudar")
	//todos.remove(1)

	//todos.clear()
	todos.print()
}
