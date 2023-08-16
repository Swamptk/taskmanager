# Task Manager CLI

-------------

Task is a simple task manager for your console, made in Go following the 7th [gophercise](https://courses.calhoun.io/courses) by [Jon Calhoun](https://twitter.com/joncalhoun), mainly using [Cobra](https://github.com/spf13/cobra) and [BoltDB](https://github.com/boltdb/bolt). This app implements the features that Calhoun adds as a bonus to this exercise, so it might serve as help for the ones who don't know how to do it.

## Message from the author

>The way I implemented those features is probably not the best, so I encourage any kind of constructive criticism.
>
>At the same time, it is an exercise for getting used to Github and Git, since this is the first repository I make public. As a result, this is more of an exercise for myself than a proper submit, so again, I would appreciate any kind of tip or advice.
>
>Finally, I would like to apologize for my English, which obviously is not my main language :D

## Status of the Project

The project should be fully functioning and every feature listed below should work properly. That being said, there may be cases that could crash the program since I have not tested it completely.

## Getting Started

### Installing

Download the files and navigate into the folder. From there you can install it using `go install` so the app gets added to the path:

```bash
go install .
```

This way you will be able to access the app at anytime using the command `task`.

In case you do not want to install the CLI, you can achieve the same behavior with `go run`:

```bash
go run .
```

However, the `go install` method is preferred, since it makes it easier to work with the commands of the CLI.

### Database

This project makes use of [BoltDB](https://github.com/boltdb/bolt) for storing your tasks.

The app will create a `tasks.db` file inside your home directory. This is a cross-platform feature thanks to the [go-homedir](https://github.com/boltdb/bolt) package.

In case you want to place your database anywhere else, you can specify your desired path inside the `home` variable in `main.go`:

```go
// main.go
func main() {
    home, _ := "/Your/Path"
    dbPath := filepath.Join(home, "tasks.db")
    ...
}
```

## Usage

`task` is fairly easy to use and includes documentation and examples for every command available.

To explain its usage, we will assume that the `go install` method was used.

### Help

When calling the `task` command the program will show the general command help menu, showing every command available as well as a short description for all of them. This feature is included by `Cobra`, which is the CLI framework that we choose for this project.

```bash
$ task
Task is fast and easy way of organizing your TODOs inside your terminal.

Usage:
  task [command]

Available Commands:
  add         Add a task to your task list.
  ...
```

You can get specific documentation for each command using the `-h` or `--help` flags:

```bash
task [command] --help
```

### Adding Tasks

You can add a new task to your list using the `add` command as follows:

```bash
$ task add Your new task
Added "Your new task" to your task list.
```

It will also update your database to add the task.

### List your Tasks

To see your pending tasks, use the `list` command:

```bash
$ task list
1 - Walk the dog
2 - Your new task
```

The indexes shown with this command will be the ones used for setting your tasks as done or to remove them.

### Task Completed

When you have completed your task, you can mark it as done using the `do` command:

```bash
$ task do 1
Task 1 has been marked as done.
```

You can mark more than one task as done at the same time by separating the indexes with spaces.

```bash
$ task do 1 2 4
Task 1 has been marked as done.
Task 2 has been marked as done.
Task 4 has been marked as done.
```

When you mark a task as done, it will remain in the database for one day, so you can keep track of what you have accomplished today!

To see what tasks you have completed, run the `completed` command:

```bash
$ task completed
✅ Do washes
✅ Go shopping
✅ Walk the dog
✅ New task!
```

The tasks marked as completed will be removed completely from the database a day after completion, once the program is launched.

### Remove a Task

If you want to delete a task without marking it as done, you can remove it directly using `rm`:

```bash
$ task rm 1
Task 1 has been removed.
```

Like the `do` command, it supports more that one index at a time
