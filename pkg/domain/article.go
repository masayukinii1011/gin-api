package domain

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Articles []Article
