package repo

import (
	"github.com/ThroughTheThornsToTheStarss/todo/internal/domain"
)

func (r *PostgresRepository) CreateTodo(todoBody string) error {
	todo := &Todo{Body: todoBody, Completed: false}
	return r.db.Create(todo).Error
}

func (r *PostgresRepository) GetAllTodo() ([]*domain.Todo, error) {
	var models []Todo
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	
	todos := make([]*domain.Todo, 0, len(models))
	for i := range models {
		todos = append(todos, &domain.Todo{
			ID:        models[i].ID,
			Body:      models[i].Body,
			Completed: models[i].Completed,
		})
	}
	return todos, nil
}

func (r *PostgresRepository) DeleteTodo(todoID int) error {
	return r.db.Delete(&Todo{}, todoID).Error
}

func (r *PostgresRepository) UpdateTodo(todoID int) error {
	return r.db.Model(&Todo{}).Where("id = ?", todoID).Update("completed", true).Error
}