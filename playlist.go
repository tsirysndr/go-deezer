package deezer

type PlaylistService service

type Playlist struct {
	ID            int    `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Public        bool   `json:"public,omitempty"`
	Link          string `json:"link,omitempty"`
	Picture       string `json:"picture,omitempty"`
	PictureSmall  string `json:"picture_small,omitempty"`
	PictureMedium string `json:"picture_medium,omitempty"`
	PictureBig    string `json:"picture_big,omitempty"`
	PictureXL     string `json:"picture_xl,omitempty"`
	Checksum      string `json:"checksum,omitempty"`
	Tracklist     string `json:"tracklist,omitempty"`
	CreationDate  string `json:"creation_date,omitempty"`
	User          *User  `json:"user,omitempty"`
	Type          string `json:"type,omitempty"`
}
