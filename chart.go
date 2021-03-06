package deezer

type ChartService service

func (s *ChartService) Get() {

}

func (s *ChartService) GetTracks() (*Tracks, error) {
	var err error
	tracks := new(Tracks)
	s.client.base.Get("chart/0/tracks").Receive(tracks, err)
	return tracks, err
}

func (s *ChartService) GetAlbums() (*Albums, error) {
	var err error
	albums := new(Albums)
	s.client.base.Get("chart/0/albums").Receive(albums, err)
	return albums, err
}

func (s *ChartService) GetArtists() (*Artists, error) {
	var err error
	artists := new(Artists)
	s.client.base.Get("chart/0/artists").Receive(artists, err)
	return artists, err
}

func (s *ChartService) GetPlaylists() (*Playlists, error) {
	var err error
	playlists := new(Playlists)
	s.client.base.Get("chart/0/playlists").Receive(playlists, err)
	return playlists, err
}
