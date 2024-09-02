package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/reuben-james/restapidemo/pkg/mocks"
	"github.com/reuben-james/restapidemo/pkg/models"

	"github.com/google/uuid"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var article models.Article
	json.Unmarshal(body, &article)

	article.Id = (uuid.New()).String()
	// Allow setting all fields on creation? 
	mocks.Articles = append(mocks.Articles, article)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("Created article %s", article.Title))
}