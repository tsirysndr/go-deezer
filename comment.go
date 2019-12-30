package deezer

type CommentService service

type Comment struct {
	ID     int    `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Date   int    `json:"date,omitempty"`
	Author *User  `json:"author,omitempty"`
	Type   string `json:"type,omitempty"`
}
