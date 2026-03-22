package api

import (
	"net/http"
	"strconv"
)

func (api *apiConfig) HandlerUpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := r.PathValue("id")
	if todoID == "" {
		respondWithError(w, http.StatusBadRequest, "todo_id is required")
		return
	}

	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil || todoIDInt == 0 {
		respondWithError(w, http.StatusBadRequest, "account_id must be a positive integer")
		return
	}

	err = api.todoUC.UpdateTodo(int(todoIDInt))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot update account")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
