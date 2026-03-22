package api

import (
	"encoding/json"
	"net/http"

	"github.com/ThroughTheThornsToTheStarss/todo/internal/domain"
)

func (api *apiConfig) HandlercreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithJSON(w, http.StatusBadRequest, "invalid json")
		return
	}
	if todo.Body == "" {
		respondWithError(w, http.StatusBadRequest, "body is required")
		return
	}
	if err := api.todoUC.CreateTodo(todo.Body); err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot create todo")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "todo created",
	})
}
