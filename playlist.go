package deezer

import "fmt"

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

func (s *PlaylistService) Get(ID string) (*Playlist, error) {
	var err error
	playlist := new(Playlist)
	s.client.base.Path("playlist/").Get(ID).Receive(playlist, err)
	return playlist, err
}

func (s *PlaylistService) GetComments(ID string) (*Comments, error) {
	var err error
	comments := new(Comments)
	path := fmt.Sprintf("playlist/%s/comments", ID)
	s.client.base.Get(path).Receive(comments, err)
	return comments, err
}

func (s *PlaylistService) GetFans(ID string) (*Fans, error) {
	var err error
	fans := new(Fans)
	path := fmt.Sprintf("playlist/%s/fans", ID)
	s.client.base.Get(path).Receive(fans, err)
	return fans, err
}
