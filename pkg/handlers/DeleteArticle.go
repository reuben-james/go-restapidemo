package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reuben-james/restapidemo/pkg/mocks"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range mocks.Articles {
		if article.Id == id {
			mocks.Articles = append(mocks.Articles[:index], mocks.Articles[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(fmt.Sprintf("Deleted article %s", id))
			break
		}
	}
}