# Go CLI Todo Manager

A simple command-line Todo Manager built in **Go** that allows users to create, view, update, and delete tasks directly from the terminal.

This project demonstrates core Go concepts such as **structs, slices, file persistence using JSON, and CLI argument handling**.

---

## Features

* Add new todos
* List all todos
* Mark todos as completed
* Delete todos
* Persistent storage using **JSON**
* Clean CLI feedback messages

---

## Project Structure

```
todo-cli
│
├── main.go
├── todo
│   ├── manager.go
│   └── todo.go
│
└── todos.json
```

---

## Installation

Clone the repository:

```
git clone https://github.com/Amodinee09/todo-cli.git
cd todo-cli
```

Make sure Go is installed:

```
go version
```

Run the CLI:

```
go run main.go
```

---

## Usage

### Add a Todo

```
go run main.go add "Learn Go"
```

Output:

```
✔ Todo 'Learn Go' added successfully (ID: 1)
```

---

### List Todos

```
go run main.go list
```

Example output:

```
ID    TASK            STATUS     CREATED
1     Learn Go        pending    2026-03-05
```

---

### Mark Todo as Done

```
go run main.go done 1
```

Output:

```
✔ Todo 1 marked as done
```

---

### Delete a Todo

```
go run main.go delete 1
```

Output:

```
✔ Todo 1 deleted successfully
```

---

## Data Persistence

Todos are stored in a local JSON file:

```
todos.json
```

Example:

```
[
  {
    "ID": 1,
    "Task": "Learn Go",
    "Status": "pending",
    "CreatedAt": "2026-03-05T18:30:00Z"
  }
]
```

---

## Concepts Used

* Go Structs
* Slices
* File I/O
* JSON Encoding / Decoding
* CLI Argument Parsing
* Basic Project Structuring in Go

---

## Future Improvements

* Add task priority
* Add due dates
* Add colored CLI output
* Add unit tests
* Build binary CLI tool

---

## Author

**Amodinee Nagrale**

---

If you found this project helpful, feel free to ⭐ the repository.
