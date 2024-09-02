package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/reuben-james/restapidemo/pkg/models"
)

const QueryUpdateArticle = "UPDATE articles SET title = $2, description = $3, content = $4, tags = $5 WHERE id = $1 RETURNING id;"

func (h handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedArticle models.Article
	json.Unmarshal(body, &updatedArticle)

	err = h.DB.QueryRow(
		QueryUpdateArticle, 
		&id, 
		&updatedArticle.Title, 
		&updatedArticle.Desc, 
		&updatedArticle.Content,
		pq.Array(&updatedArticle.Tags),
	).Scan(&id)
	if err != nil {
        log.Printf("Failed to execute update on article (ID: %s): %v", id, err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Updated article (ID: %s)", id))
}