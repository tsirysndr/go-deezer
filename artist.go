package deezer

import (
	"fmt"
)

type ArtistService service

type Artist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	NbAlbum       int    `json:"nb_album"`
	NbFan         int    `json:"nb_fan"`
	Radio         bool   `json:"radio"`
	TrackList     string `json:"tracklist"`
	Type          string `json:"type"`
}

type Top struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type Albums struct {
	Data  []Album `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type Playlists struct {
	Data  []Playlist `json:"data"`
	Total int        `json:"total"`
	Next  string     `json:"next"`
}

type Fans struct {
	Data  []User `json:"data"`
	Total int    `json:"total"`
	Next  string `json:"next"`
}

type Related struct {
	Data  []Artist `json:"data"`
	Total int      `json:"total"`
	Next  string   `json:"next"`
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
