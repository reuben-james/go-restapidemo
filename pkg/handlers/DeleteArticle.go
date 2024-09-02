package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const QueryDelArticle = "DELETE FROM articles WHERE id = $1;"

func (h handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := h.DB.Query(QueryDelArticle, &id)
	if err != nil {
		log.Printf("Failed to delete article (ID: %s): %v", id, err)
        w.WriteHeader(http.StatusInternalServerError)
        return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Deleted article (ID: %s)", id))
}