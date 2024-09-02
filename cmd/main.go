package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reuben-james/restapidemo/pkg/db"
	"github.com/reuben-james/restapidemo/pkg/handlers"
	"github.com/spf13/viper"
)

func GetEnv(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env file '.env'")
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Failed to assert %s is a string", value)
	}

	return value
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Articles API!\n\nYou can browse all of the articles at: /articles")
}

func HandleRequests(DB *sql.DB) {
    h := handlers.New(DB)
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", HomePage)
    myRouter.HandleFunc("/articles", h.GetAllArticles).Methods(http.MethodGet)
    myRouter.HandleFunc("/articles/{id}", h.GetArticle).Methods(http.MethodGet)
    myRouter.HandleFunc("/articles", h.AddArticle).Methods(http.MethodPost)
    myRouter.HandleFunc("/articles/{id}", h.UpdateArticle).Methods(http.MethodPut)
    myRouter.HandleFunc("/articles/{id}", h.DeleteArticle).Methods(http.MethodDelete)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    var (
		DBHost     = GetEnv("DB_HOST")
		DBPort, _  = strconv.ParseUint(GetEnv("DB_PORT"), 10, 64)
		DBUser     = GetEnv("DB_USER")
		DBPassword = GetEnv("DB_PASSWORD")
		DBName     = GetEnv("DB_NAME")
	)

    DB := db.Connect(DBHost, DBPort, DBUser, DBPassword, DBName)
    db.CreateTable(DB)
    HandleRequests(DB)
    db.CloseConnection(DB)
}