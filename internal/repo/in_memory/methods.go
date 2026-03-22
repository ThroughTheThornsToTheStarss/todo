package inmemory

import (
	"errors"
	"github.com/ThroughTheThornsToTheStarss/todo/internal/domain"
)

func (mr *MemoryRepository) CreateTodo(todoBody string) error {
	if todoBody == "" {
		return errors.New("Todo body cannot be empty")
	}
	newTodo := domain.Todo{
		ID:        len(mr.todos) + 1,
		Body:      todoBody,
		Completed: false,
	}
	mr.todos = append(mr.todos, &newTodo)
	return nil
}

func (mr *MemoryRepository) DeleteTodo(todoID int) error {
	for i, todo := range mr.todos {
		if todo.ID == todoID {
			mr.todos = append(mr.todos[:i], mr.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

func (mr *MemoryRepository) GetAllTodo() ([]*domain.Todo, error) {
	return mr.todos, nil
}

func (mr *MemoryRepository) UpdateTodo(todoID int) error {
	for _, todo := range mr.todos {
		if todo.ID == todoID {
			todo.Completed = true
			return nil
		}
	}
	return errors.New("Todo not found")
}
