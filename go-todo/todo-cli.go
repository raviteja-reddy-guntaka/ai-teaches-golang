package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo

func main() {
	for {
		printMenu()
		choice := getUserInput("Enter your choice: ")

		switch choice {
		case "1":
			addTodo()
		case "2":
			listTodos()
		case "3":
			markTodoAsDone()
		case "4":
			deleteTodo()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\n--- Todo List ---")
	fmt.Println("1. Add Todo")
	fmt.Println("2. List Todos")
	fmt.Println("3. Mark Todo as Done")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func addTodo() {
	task := getUserInput("Enter todo task: ")
	todo := Todo{
		ID:   len(todos) + 1,
		Task: task,
		Done: false,
	}
	todos = append(todos, todo)
	fmt.Println("Todo added successfully!")
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	for _, todo := range todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Task)
	}
}

func markTodoAsDone() {
	id := getUserInput("Enter todo ID to mark as done: ")
	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos[i].Done = true
			fmt.Println("Todo marked as done!")
			return
		}
	}
	fmt.Println("Todo not found.")
}

func deleteTodo() {
	id := getUserInput("Enter todo ID to delete: ")
	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted successfully!")
			return
		}
	}
	fmt.Println("Todo not found.")
}
