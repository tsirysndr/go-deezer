package deezer

import "fmt"

type GenreService service

type Genre struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Picture       string `json:"picture,omitempty"`
	PictureBig    string `json:"picture_big,omitempty"`
	PictureMedium string `json:"picture_medium,omitempty"`
	PictureSmall  string `json:"picture_small,omitempty"`
	PictureXL     string `json:"picture_xl,omitempty"`
	Type          string `json:"type,omitempty"`
}

type Genres struct {
	Data  []Genre `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

func (s *GenreService) Get(ID int) (*Genre, error) {
	var err error
	genre := new(Genre)
	path := fmt.Sprintf("genre/%d", ID)
	s.client.base.Get(path).Receive(genre, err)
	return genre, err
}

func (s *GenreService) List() (*Genres, error) {
	var err error
	genres := new(Genres)
	s.client.base.Get("genre").Receive(genres, err)
	return genres, err
}

func (s *GenreService) GetArtists(ID int) (*Artists, error) {
	var err error
	artists := new(Artists)
	path := fmt.Sprintf("genre/%d/artists", ID)
	s.client.base.Get(path).Receive(artists, err)
	return artists, err
}
