package main

// Note
type Note struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}

// Notes
type Notes []Note
