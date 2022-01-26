package proto

import "time"

type RspFindArticle struct {
	ID uint `json:"id"`
	Title   string `json:"title"`
	Name string  `json:"name"`
	CreatedAt time.Time `json:"CreatedAt"`
	Desc    string ` json:"desc"`
	Content string `json:"content"`
	Img     string ` json:"img"`
}