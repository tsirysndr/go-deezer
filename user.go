package deezer

type UserService service

type User struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Link          string `json:"link,omitempty"`
	Picture       string `json:"picture,omitempty"`
	PictureSmall  string `json:"picture_small,omitempty"`
	PictureMedium string `json:"picture_medium,omitempty"`
	PictureBig    string `json:"picture_big,omitempty"`
	PictureXL     string `json:"picture_xl,omitempty"`
	Tracklist     string `json:"tracklist,omitempty"`
	Type          string `json:"type,omitempty"`
}
