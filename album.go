package deezer

import "fmt"

type AlbumService service

type Album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	UPC         string `json:"upc"`
	Link        string `json:"link"`
	Share       string `json:"share"`
	Cover       string `json:"cover"`
	CoverSmall  string `json:"cover_small"`
	CoverMedium string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXL     string `json:"cover_xl"`
	GenreID     int    `json:"genre_id"`
	Genres      *struct {
		Data []Genre `json:"data"`
	} `json:"genres"`
	Label                 string   `json:"label"`
	NbTracks              int      `json:"nb_tracks"`
	Duration              int      `json:"duration"`
	Fans                  int      `json:"fans"`
	Rating                int      `json:"rating"`
	ReleaseDate           string   `json:"release_date"`
	RecordType            string   `json:"record_type"`
	Available             bool     `json:"available"`
	TrackList             string   `json:"tracklist"`
	ExplicitLyrics        bool     `json:"explicit_lyrics"`
	ExplicitContentLyrics int      `json:"explicit_content_lyrics"`
	ExplicitContentCover  int      `json:"explicit_content_cover"`
	Contributors          []Artist `json:"contributors"`
	Artist                *Artist  `json:"artist"`
	Type                  string   `json:"type"`
	Tracks                *struct {
		Data []Track `json:"data"`
	} `json:"tracks"`
}

type Genre struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Type    string `json:"type"`
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
