package deezer

type SearchService service

type SearchParams struct {
	Q string `url:"q"`
}

func (s *SearchService) Get(text string) (*Tracks, error) {
	var err error
	params := SearchParams{
		Q: text,
	}
	tracks := new(Tracks)
	s.client.base.Get("search").QueryStruct(params).Receive(tracks, err)
	return tracks, err
}
