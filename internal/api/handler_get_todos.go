package api

import (
	"net/http"
)

func (api *apiConfig) HandlerGetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := api.todoUC.GetAllTodo()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot get todos")
		return
	}

	respondWithJSON(w, http.StatusOK, todos)
}
