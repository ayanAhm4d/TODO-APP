package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo-app/storage"
	"todo-app/todo"
)

func main() {
	store := storage.NewInMemoryStorage()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo App ---")
		fmt.Println("1. Add todo")
		fmt.Println("2. List todos")
		fmt.Println("3. Mark todo as complete")
		fmt.Println("4. Delete todo")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTodo(scanner, store)
		case "2":
			listTodos(store)
		case "3":
			markTodoComplete(scanner, store)
		case "4":
			deleteTodo(scanner, store)
		case "5":
			fmt.Println("Good Bye!")
			return
		default:
			fmt.Println("invalid, inter correct choice.")
		}

	}
}

func addTodo(scanner *bufio.Scanner, store storage.Storage) {
	fmt.Print("inter todo description: ")
	scanner.Scan()

	newTodo := todo.NewTodo(scanner.Text())
	err := store.Add(newTodo)
	if err != nil {
		fmt.Println("Error adding todo", err)
	} else {
		fmt.Println("Todo added succesfully.")
	}
}

func listTodos(store storage.Storage) {
	todos, err := store.List()
	if err != nil {
		fmt.Println("Error in listing todos:", err)
		return
	}

	if len(todos) == 0 {
		fmt.Println("no todo found")
		return
	}

	status := " "
	for _, todo := range todos {
		if todo.IsComplelted {
			status = "X"
		}
		fmt.Printf("[%s] %d. %s\n", status, todo.ID, todo.Description)
	}
}

func markTodoComplete(scanner *bufio.Scanner, store storage.Storage) {
	fmt.Print("Inter todo ID to mark complete: ")
	scanner.Scan()

	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID, try again.")
		return
	}
	todo, err := store.Get(id)
	if err != nil {
		fmt.Println("Error getting todo: ", err)
		return
	}

	todo.MarkComplete()
	err = store.Update(todo)
	if err != nil {
		fmt.Println("Error updating todo:", err)
	} else {
		fmt.Println("Todo marked as complete!")
	}
}

func deleteTodo(scanner *bufio.Scanner, store storage.Storage) {
	fmt.Print("Enter todo ID to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	err = store.Delete(id)
	if err != nil {
		fmt.Println("Error deleting todo:", err)
	} else {
		fmt.Println("Todo deleted successfully!")
	}
}
