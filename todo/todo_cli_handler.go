package todo

import (
	"context"
	"fmt"
	"test-mongo/models"
)

// GET - /todos
// return all todos
type TodoCLIHandler struct {
	uc *TodoUsecase
}

func NewTodoHandler(uc *TodoUsecase) *TodoCLIHandler {
	return &TodoCLIHandler{uc}
}

func (h *TodoCLIHandler) GetTodos(ctx context.Context) {
	todos, err := h.uc.GetTodos(ctx)
	if err != nil {
		panic(err)
	}

	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func (h TodoCLIHandler) InsertTodo(ctx context.Context, todo models.Todo) {
	if err := h.uc.InsertTodo(ctx, todo); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Todo inserted")
}
