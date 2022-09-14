package models

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{ID: 1, Title: "Into The Woods", Author: "Jack"},
	{ID: 2, Title: "Breaking Back", Author: "Simon"},
}
