package deezer

import (
	"fmt"
)

type ArtistService service

type Artist struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Link          string `json:"link,omitempty"`
	Share         string `json:"share,omitempty"`
	Picture       string `json:"picture,omitempty"`
	PictureSmall  string `json:"picture_small,omitempty"`
	PictureMedium string `json:"picture_medium,omitempty"`
	PictureBig    string `json:"picture_big,omitempty"`
	PictureXl     string `json:"picture_xl,omitempty"`
	NbAlbum       int    `json:"nb_album,omitempty"`
	NbFan         int    `json:"nb_fan,omitempty"`
	Radio         bool   `json:"radio,omitempty"`
	TrackList     string `json:"tracklist,omitempty"`
	Type          string `json:"type,omitempty"`
}

type Top struct {
	Data  []Track `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

type Albums struct {
	Data  []Album `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

type Playlists struct {
	Data  []Playlist `json:"data,omitempty"`
	Total int        `json:"total,omitempty"`
	Next  string     `json:"next,omitempty"`
}

type Artists struct {
	Data  []Artist `json:"data,omitempty"`
	Total int      `json:"total,omitempty"`
	Next  string   `json:"next,omitempty"`
}

type Fans struct {
	Data  []User `json:"data,omitempty"`
	Total int    `json:"total,omitempty"`
	Next  string `json:"next,omitempty"`
}

type Related struct {
	Data  []Artist `json:"data,omitempty"`
	Total int      `json:"total,omitempty"`
	Next  string   `json:"next,omitempty"`
}

type Radio struct {
	Data  []Track `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

type Comments struct {
	Data  []Comment `json:"data,omitempty"`
	Total int       `json:"total,omitempty"`
	Next  string    `json:"next,omitempty"`
}

func (s *ArtistService) Get(ID string) (*Artist, error) {
	var err error
	artist := new(Artist)
	s.client.base.Path("artist/").Get(ID).Receive(artist, err)
	return artist, err
}

func (s *ArtistService) GetTopFive(ID string) (*Top, error) {
	var err error
	top := new(Top)
	path := fmt.Sprintf("artist/%s/top", ID)
	s.client.base.Get(path).Receive(top, err)
	return top, err
}

func (s *ArtistService) GetAlbums(ID string) (*Albums, error) {
	var err error
	albums := new(Albums)
	path := fmt.Sprintf("artist/%s/albums", ID)
	s.client.base.Get(path).Receive(albums, err)
	return albums, err
}

func (s *ArtistService) GetPlaylists(ID string) (*Playlists, error) {
	var err error
	playlists := new(Playlists)
	path := fmt.Sprintf("artist/%s/playlists", ID)
	s.client.base.Get(path).Receive(playlists, err)
	return playlists, err
}

func (s *ArtistService) GetFans(ID string) (*Fans, error) {
	var err error
	fans := new(Fans)
	path := fmt.Sprintf("artist/%s/fans", ID)
	s.client.base.Get(path).Receive(fans, err)
	return fans, err
}

func (s *ArtistService) GetRelated(ID string) (*Related, error) {
	var err error
	related := new(Related)
	path := fmt.Sprintf("artist/%s/related", ID)
	s.client.base.Get(path).Receive(related, err)
	return related, err
}

func (s *ArtistService) GetRadio(ID string) (*Radio, error) {
	var err error
	radio := new(Radio)
	path := fmt.Sprintf("artist/%s/radio", ID)
	s.client.base.Get(path).Receive(radio, err)
	return radio, err
}

func (s *ArtistService) GetComments(ID string) (*Comments, error) {
	var err error
	comments := new(Comments)
	path := fmt.Sprintf("artist/%s/comments", ID)
	s.client.base.Get(path).Receive(comments, err)
	return comments, err
}
