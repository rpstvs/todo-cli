package main

import "log"

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		log.Fatal(err)
	}
	processArgs(&todos)

	err = storage.Save(todos)

	if err != nil {
		log.Fatal(err)
	}

}
