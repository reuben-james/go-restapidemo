package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lib/pq"
	"github.com/reuben-james/restapidemo/pkg/models"
)

const QueryGetAllArticles string = "SELECT * FROM articles;"

func (h handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query(QueryGetAllArticles)
	if err != nil {
		log.Println("Failed to execute query:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var articles = make([]models.Article, 0)
	for results.Next() {
		var article models.Article
		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content, pq.Array(&article.Tags))
		if err != nil {
			log.Println("Failed to scan results:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}

		articles = append(articles, article)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}