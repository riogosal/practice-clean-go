package todo

import (
	"context"
	"test-mongo/models"
)

type TodoMemoryDataLayer struct {
	todos []models.Todo
}

func NewTodoMemoryDataLayer() *TodoMemoryDataLayer {
	return &TodoMemoryDataLayer{
		todos: []models.Todo{
			{
				UUID:      "63057e5f-bc8a-4990-9966-c1b13fefde43",
				Title:     "Buy milk",
				Completed: false,
			},
			{
				UUID:      "8b84ecd0-ddd8-4109-90fe-8d5ad5b5aade",
				Title:     "Buy eggs",
				Completed: true,
			},
		},
	}
}

func (data *TodoMemoryDataLayer) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return data.todos, nil
}

func (data *TodoMemoryDataLayer) InsertTodo(ctx context.Context, todo models.Todo) error {
	data.todos = append(data.todos, todo)
	return nil
}
