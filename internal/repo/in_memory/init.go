package inmemory

import (
	"github.com/ThroughTheThornsToTheStarss/todo/internal/domain"
)

type MemoryRepository struct {
	todos []*domain.Todo
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		todos: []*domain.Todo{},
	}
}
