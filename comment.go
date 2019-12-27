package deezer

type CommentService service

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Date   int    `json:"date"`
	Author *User  `json:"author"`
	Type   string `json:"type"`
}
