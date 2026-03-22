package usecase

import "github.com/ThroughTheThornsToTheStarss/todo/internal/domain"

func (u *todoUscase) CreateTodo(todoBody string) error {
	return u.repo.CreateTodo(todoBody)
}

func (u *todoUscase) DeleteTodo(todoID int) error {
	return u.repo.DeleteTodo(todoID)
}
func (u *todoUscase) GetAllTodo() ([]*domain.Todo, error) {
	return u.repo.GetAllTodo()
}
func (u *todoUscase) UpdateTodo(todoID int) error {
	return u.repo.UpdateTodo(todoID)
}
