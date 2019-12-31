package deezer

import "fmt"

type AlbumService service

type Album struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	UPC         string `json:"upc,omitempty"`
	Link        string `json:"link,omitempty"`
	Share       string `json:"share,omitempty"`
	Cover       string `json:"cover,omitempty"`
	CoverSmall  string `json:"cover_small,omitempty"`
	CoverMedium string `json:"cover_medium,omitempty"`
	CoverBig    string `json:"cover_big,omitempty"`
	CoverXL     string `json:"cover_xl,omitempty"`
	GenreID     int    `json:"genre_id,omitempty"`
	Genres      *struct {
		Data []Genre `json:"data,omitempty"`
	} `json:"genres,omitempty"`
	Label                 string   `json:"label,omitempty"`
	NbTracks              int      `json:"nb_tracks,omitempty"`
	Duration              int      `json:"duration,omitempty"`
	Fans                  int      `json:"fans,omitempty"`
	Rating                int      `json:"rating,omitempty"`
	ReleaseDate           string   `json:"release_date,omitempty"`
	RecordType            string   `json:"record_type,omitempty"`
	Available             bool     `json:"available,omitempty"`
	TrackList             string   `json:"tracklist,omitempty"`
	ExplicitLyrics        bool     `json:"explicit_lyrics,omitempty"`
	ExplicitContentLyrics int      `json:"explicit_content_lyrics,omitempty"`
	ExplicitContentCover  int      `json:"explicit_content_cover,omitempty"`
	Contributors          []Artist `json:"contributors,omitempty"`
	Artist                *Artist  `json:"artist,omitempty"`
	Type                  string   `json:"type,omitempty"`
	Tracks                *struct {
		Data []Track `json:"data,omitempty"`
	} `json:"tracks,omitempty"`
}

type Tracks struct {
	Data  []Track `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

func (s *AlbumService) Get(ID string) (*Album, error) {
	var err error
	album := new(Album)
	s.client.base.Path("album/").Get(ID).Receive(album, err)
	return album, err
}

func (s *AlbumService) GetComments(ID string) (*Comments, error) {
	var err error
	comments := new(Comments)
	path := fmt.Sprintf("album/%s/comments", ID)
	s.client.base.Get(path).Receive(comments, err)
	return comments, err
}

func (s *AlbumService) GetFans(ID string) (*Fans, error) {
	var err error
	fans := new(Fans)
	path := fmt.Sprintf("album/%s/fans", ID)
	s.client.base.Get(path).Receive(fans, err)
	return fans, err
}

func (s *AlbumService) GetTracks(ID string) (*Tracks, error) {
	var err error
	tracks := new(Tracks)
	path := fmt.Sprintf("album/%s/tracks", ID)
	s.client.base.Get(path).Receive(tracks, err)
	return tracks, err
}
