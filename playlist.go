package deezer

type PlaylistService service

type Playlist struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Public        bool   `json:"public"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXL     string `json:"picture_xl"`
	Checksum      string `json:"checksum"`
	Tracklist     string `json:"tracklist"`
	CreationDate  string `json:"creation_date"`
	User          *User  `json:"user"`
	Type          string `json:"type"`
}
