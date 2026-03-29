package models

//go:generate easyjson -all note.go

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Notes = []Note{
	{ID: 1, Title: "Buy milk", Content: "2 liters"},
	{ID: 2, Title: "Read book", Content: "10 pages"},
	{ID: 3, Title: "movie", Content: "horror"},
}
