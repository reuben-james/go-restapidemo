package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/reuben-james/restapidemo/pkg/mocks"
)

const QueryExistsTableArticle string = "SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'articles' );"
const QueryCreateTableArticle string = "CREATE TABLE articles (id VARCHAR(36) PRIMARY KEY, title VARCHAR(25) NOT NULL, description VARCHAR(50) NOT NULL, content VARCHAR(512) NOT NULL, tags TEXT[]);"
const QueryInitArticles string = "INSERT INTO articles (id,title,description,content,tags) VALUES ($1, $2, $3, $4, $5) RETURNING id;"

func Connect(host string, port uint64, user string, password string, dbname string) *sql.DB {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfull connected to %s.\n", dbname)
	return db
}

func CreateTable(db *sql.DB) {
	var exists bool
	if err := db.QueryRow(QueryExistsTableArticle).Scan(&exists); err != nil {
        fmt.Println("Failed to execute query:", err)
        return
    }
	if !exists {
		_, err := db.Exec(QueryCreateTableArticle)

		if err != nil {
			fmt.Println("Failed to execute query:", err)
			return
		}

		fmt.Println("Table created successfully.")

		fmt.Println("Adding mock articles into 'articles' table...")
		for _, article := range mocks.Articles {
			// fmt.Printf("Executing Query: %s\nWith Parameters: %v, %v, %v, %v, %v\n", 
    		// 	QueryInitArticles, article.Id, &article.Title, &article.Desc, &article.Content, article.Tags)
			
			err := db.QueryRow(
				QueryInitArticles, 
				article.Id, 
				&article.Title, 
				&article.Desc,
				&article.Content,
				pq.Array(article.Tags),
			).Scan(&article.Id)
			
			if err != nil {
				log.Printf("Failed to insert article (ID: %s): %v\n", article.Id, err)
				continue
			}

			log.Printf("Inserted article with ID: %s\n", article.Id)
		}
		fmt.Println("Completed adding mock articles to Table 'articles'")
	} else {
		fmt.Println("Table 'articles' already exists.")
	}
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}