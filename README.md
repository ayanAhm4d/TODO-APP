
## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)](https://linkedin.com/in/ayanahmad15) [![X](https://img.shields.io/badge/X-black.svg?logo=X&logoColor=white)](https://x.com/ayanAhm4d) 

# Go Todo Application

This is a simple command-line todo application written in Go. It demonstrates the use of interfaces, package organization, and basic CRUD operations with both in-memory and MySQL storage options.

## Features

- Add new todos
- List all todos
- Mark todos as complete
- Delete todos
- Flexible storage backend (in-memory)

## Project Structure
odo-app/ </br>
├── main.go</br>
├── todo/</br>
│   └── todo.go</br>
└── storage/</br>
└── storage.go</br>

- `main.go`: Contains the main application logic and user interface
- `todo/todo.go`: Defines the Todo struct and related methods
- `storage/storage.go`: Defines the Storage interface and implementations

## Prerequisites

- Go 1.22 


## Installation

1. Clone the repository:
```
https://github.com/ayanAhm4d/TODO-APP.git
```
2. Install dependencies:
```
go mod tidy
```
## Usage

1. Run the application:
```
go run main.go
```
2. Follow the on-screen prompts to manage your todos.
