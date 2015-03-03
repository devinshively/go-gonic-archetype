package model

type Document struct {
	Id    int    `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
