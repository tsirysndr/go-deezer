package deezer

type TrackService service

type Track struct {
	ID                    int           `json:"id,omitempty"`
	Readable              bool          `json:"readable,omitempty"`
	Title                 string        `json:"title,omitempty"`
	TitleShort            string        `json:"title_short,omitempty"`
	TitleVersion          string        `json:"title_version,omitempty"`
	ISRC                  string        `json:"isrc,omitempty"`
	Link                  string        `json:"link,omitempty"`
	Duration              int           `json:"duration,omitempty"`
	Rank                  int           `json:"rank,omitempty"`
	ExplicitLyrics        bool          `json:"explicit_lyrics,omitempty"`
	ExplicitContentLyrics int           `json:"explicit_content_lyrics,omitempty"`
	ExplicitContentCover  int           `json:"explicit_content_cover,omitempty"`
	Preview               string        `json:"preview,omitempty"`
	Contributors          []Contributor `json:"contributors,omitempty"`
	Artist                *Artist       `json:"artist,omitempty"`
	Album                 *Album        `json:"album,omitempty"`
	Type                  string        `json:"type,omitempty"`
}

type Contributor struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Link          string `json:"link,omitempty"`
	Share         string `json:"share,omitempty"`
	Picture       string `json:"picture,omitempty"`
	PictureSmall  string `json:"picture_small,omitempty"`
	PictureMedium string `json:"picture_medium,omitempty"`
	PictureBig    string `json:"picture_big,omitempty"`
	PictureXL     string `json:"picture_xl,omitempty"`
	Radio         bool   `json:"radio,omitempty"`
	TrackList     string `json:"tracklist,omitempty"`
	Type          string `json:"type,omitempty"`
	Role          string `json:"role,omitempty"`
}

func (s *TrackService) Get(ID string) (*Track, error) {
	var err error
	track := new(Track)
	s.client.base.Path("track/").Get(ID).Receive(track, err)
	return track, err
}
