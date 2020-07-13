package pojo

type Article struct {
	Base
	Author_id   uint
	Category_id uint   `json:"category"`
	Title       string `json:"title;not null"`
	Text        string `json:"text;not null"`
	Views       int
}
