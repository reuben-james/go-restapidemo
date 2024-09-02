package models

type Article struct {
	Id 		string 	 `json:"Id"`
	Title 	string 	 `json:"Title"`
	Desc 	string 	 `json:"Desc"`
	Content string 	 `json:"Content"`
	Tags 	[]string `json:"Tags"`
}

// Set global defaults to ensure DB data consistency
func NewArticle() *Article {
	return &Article{
		Tags: []string{},
	}
}