package main

type Todo struct {
	task string
	done bool
}

type Todos []Todo

var tasks map[int]Todo

func (todos *Todos) add(message string) {
	todo := Todo{
		task: message,
		done: false,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) remove(id int) {

	t := *todos

	*todos = append(t[:id], t[id+1:]...)
}

func (todos *Todos) markDone(id int) {
	t := *todos

	isCompleted := t[id].done

	if !isCompleted {
		t[id].done = true

	}

}

func (todos *Todos) edit(id int, task string) {
	t := *todos

	t[id].task = task

}
