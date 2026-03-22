package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (api *apiConfig) HandlerDeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoID := r.PathValue("id")
	if todoID == "" {
		respondWithError(w, http.StatusBadRequest, "todo_id is required")
		return
	}

	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil || todoIDInt == 0 {
		respondWithError(w, http.StatusBadRequest, "todo_id must be a positive integer")
		return
	}

	err = api.todoUC.DeleteTodo(int(todoIDInt))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot delete todo")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "todo deleted",
	})
}
