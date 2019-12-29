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
