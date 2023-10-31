package todo

import (
	"context"
	"errors"
	"test-mongo/models"
)

type TodoUsecase struct {
	repo *TodoMemoryDataLayer
}

func NewTodoUsecase(repo *TodoMemoryDataLayer) *TodoUsecase {
	return &TodoUsecase{repo}
}

func (uc *TodoUsecase) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return uc.repo.GetTodos(ctx)
}

func (uc *TodoUsecase) InsertTodo(ctx context.Context, todo models.Todo) error {
	if todo.UUID == "" {
		return errors.New("UUID is required")
	}
	if todo.Title == "" {
		return errors.New("Title is required")
	}
	return uc.repo.InsertTodo(ctx, todo)
}
