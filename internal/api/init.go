package api

import (
	"net/http"

	"github.com/ThroughTheThornsToTheStarss/todo/internal/usecase"
)

type apiConfig struct {
	todoUC usecase.TodoUscase
}

func New(todoUC usecase.TodoUscase) http.Handler {

	apiCfg := &apiConfig{
		todoUC: todoUC,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/todo", apiCfg.HandlercreateTodo)
	mux.HandleFunc("GET /api/todo", apiCfg.HandlerGetAllTodos)
	mux.HandleFunc("PATCH /api/todo/{id}", apiCfg.HandlerUpdateTodo)
	mux.HandleFunc("DELETE /api/todo/{id}", apiCfg.HandlerDeleteTodo)

	return mux
}
