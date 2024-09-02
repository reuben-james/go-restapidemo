package mocks

import "github.com/reuben-james/restapidemo/pkg/models"

var Articles = []models.Article{
	{
		Id: "a77a0b50-68a1-11ef-8b78-6fe1c538af5b", 
		Title: "Article 1", 
		Desc: "This is Article 1", 
		Content: "This is the content for Article 1",
		Tags: []string{"tag-1","tag-2"},
	},
	{
		Id: "add23356-68a1-11ef-8cea-1302a92b16a5", 
		Title: "Article 2", 
		Desc: "This is Article 2", 
		Content: "This is the content for Article 2",
		Tags: nil,
	},
}