package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/reuben-james/restapidemo/pkg/models"
)

const QueryGetArticle string = "SELECT * FROM articles WHERE id = $1 ;"

func (h handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	results, err := h.DB.Query(QueryGetArticle, id)
	if err != nil {
        log.Println("Failed to query database:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	var article models.Article
	for results.Next() {
		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content, pq.Array(&article.Tags))
		if err != nil {
            log.Printf("Failed to retrieve article (ID: %s): %v\n", id, err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}