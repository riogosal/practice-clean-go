package main

import (
	"context"
	"test-mongo/models"
	"test-mongo/todo"
	"time"
)

// 3 - layers of a web application or clean architecture
// ----------------------------------------------------
// 1. Handler (deals with http requests) - GIN
// 2. Business Logic / Service / Usecase (deals with business logic) - GO
// 3. Data layer - MONGO

func main() {
	mainCtx := context.Background()

	todoRepo := todo.NewTodoMemoryDataLayer()
	todoUC := todo.NewTodoUsecase(todoRepo)
	todoH := todo.NewTodoHandler(todoUC)

	todoH.GetTodos(mainCtx)
	todoH.InsertTodo(mainCtx, models.Todo{
		UUID:  "63057e5f-bc8a-4990-9966-c1b13fefde43",
		Title: "Buy comic books",
	})

	time.Sleep(2 * time.Second)
	todoH.GetTodos(mainCtx)
}
