package deezer

type TrackService service

type Track struct {
	ID                    int           `json:"id"`
	Readable              bool          `json:"readable"`
	Title                 string        `json:"title"`
	TitleShort            string        `json:"title_short"`
	TitleVersion          string        `json:"title_version"`
	Link                  string        `json:"link"`
	Duration              int           `json:"duration"`
	Rank                  int           `json:"rank"`
	ExplicitLyrics        bool          `json:"explicit_lyrics"`
	ExplicitContentLyrics int           `json:"explicit_content_lyrics"`
	ExplicitContentCover  int           `json:"explicit_content_cover"`
	Preview               string        `json:"preview"`
	Contributors          []Contributor `json:"contributors"`
	Artist                *Artist       `json:"artist"`
	Album                 *Album        `json:"album"`
	Type                  string        `json:"type"`
}

type Contributor struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXL     string `json:"picture_xl"`
	Radio         bool   `json:"radio"`
	TrackList     string `json:"tracklist"`
	Type          string `json:"type"`
	Role          string `json:"role"`
}
