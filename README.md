## Installation

Clone the project

```bash
go build
```

Optional (LINUX):
Move the exe to the bin folder of your distro for general access.

## Usage

Commands implemented:

add - adds a task to the table

```go
todo-cli add <task>
```

remove - removes a task from the table

```go
todo-cli remove <id>
```

done - marks the task as done
(!) doesn't remove the task

```go
todo-cli done <id>
```

print - lists the tasks

```go
todo-cli print
```

help - lists the available commands

```go
todo-cli help
```
