package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lib/pq"
	"github.com/reuben-james/restapidemo/pkg/models"

	"github.com/google/uuid"
)

const QueryInsertArticle string = "INSERT INTO articles (id,title,description,content,tags) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

func (h handler) AddArticle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	article := models.NewArticle()
	
	if err := json.Unmarshal(body, article); err != nil {
		log.Println("Failed to unmarshal request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article.Id = (uuid.New()).String()

	err = h.DB.QueryRow(
		QueryInsertArticle, 
		article.Id, 
		article.Title, 
		article.Desc, 
		article.Content,
		pq.Array(article.Tags),
	).Scan(&article.Id)

	if err != nil {
		log.Println("Failed to execute database query:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("Created article (ID: %s, Title: %s)", article.Id, article.Title))
}