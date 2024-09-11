# Go CLI Todo App

A simple command-line todo list application built with Go.

### Features

- Add todos: Add new tasks to your list with a title.
- List todos: View all your current tasks, including their completion status.
- Toggle completion: Mark tasks as complete or incomplete.
- Edit todos: Update the title of existing tasks.
- Delete todos: Remove tasks from your list.

### Installation

1. Make sure you have Go installed on your system. You can download it from the official website: [golang.org](https://golang.org)

1. Clone this repository:

```sh
git clone https://github.com/ibrahimraimi/go-cli-todo.git
```

1. Navigate into the project directory:

```sh
cd go-cli-todo
```

1. Build the application:

```sh
go build
```

### Usage

```sh
./go-cli-todo [command] [options]
```

#### Available Commands

- `--add` "Task Title": Add a new todo with the specified title.
- `--list`: List all todos.
- `--toggle` [index]: Toggle the completion status of the todo at the given index.
- `--edit` [index] "New Title": Edit the title of the todo at the given index.
- `--del` [index]: Delete the todo at the given index.
- `--help` Display help information about the available commands.

#### Examples

- Add a new todo:


```sh
./go-cli-todo --add "Buy groceries"
```

- List all todos:

```sh
./go-cli-todo --list
```

- Mark the first todo as complete:

```sh
./go-cli-todo --toggle 0
```

- Edit the second todo:

```sh
./go-cli-todo --edit 1 "Finish report"
```

- Delete the third todo:

```sh
./go-cli-todo --del 2
```
