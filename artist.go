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
