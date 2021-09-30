package domain

type Article struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Articles []Article
