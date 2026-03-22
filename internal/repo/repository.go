package repo

import "github.com/ThroughTheThornsToTheStarss/todo/internal/domain"

type Repository interface {
	CreateTodo(todoBody string) error
	DeleteTodo(todoID int) error
	GetAllTodo() ([]*domain.Todo, error)
	UpdateTodo(todoID int) error
}
