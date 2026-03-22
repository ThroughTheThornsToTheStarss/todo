package usecase

import (
	"github.com/ThroughTheThornsToTheStarss/todo/internal/domain"
	"github.com/ThroughTheThornsToTheStarss/todo/internal/repo"
)

type TodoUscase interface {
	CreateTodo(todoBody string) error
	DeleteTodo(todoID int) error
	GetAllTodo() ([]*domain.Todo, error)
	UpdateTodo(todoID int) error
}

type todoUscase struct {
	repo repo.Repository
}

func NewTodoUsecase(r repo.Repository) TodoUscase {
	return &todoUscase{repo: r}
}
