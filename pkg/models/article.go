package models

type Article struct {
	Id 		string 	 `json:"id"`
	Title 	string 	 `json:"title"`
	Desc 	string 	 `json:"desc"`
	Content string 	 `json:"content"`
	Tags 	[]string `json:"tags,omitempty"`
}