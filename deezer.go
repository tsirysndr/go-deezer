package deezer

import (
	"github.com/dghubble/sling"
)

type Client struct {
	base      *sling.Sling
	common    service
	Album     *AlbumService
	Artist    *ArtistService
	Chart     *ChartService
	Comment   *CommentService
	Editorial *EditorialService
	Genre     *GenreService
	Infos     *InfosService
	Options   *OptionsService
	Playlist  *PlaylistService
	Radio     *RadioService
	Search    *SearchService
	Track     *TrackService
	User      *UserService
}

type service struct {
	client *Client
}

func NewClient() *Client {
	base := sling.New().Base("https://api.deezer.com/")
	c := &Client{}
	c.base = base
	c.common.client = c
	c.Album = (*AlbumService)(&c.common)
	c.Artist = (*ArtistService)(&c.common)
	c.Chart = (*ChartService)(&c.common)
	c.Comment = (*CommentService)(&c.common)
	c.Editorial = (*EditorialService)(&c.common)
	c.Genre = (*GenreService)(&c.common)
	c.Infos = (*InfosService)(&c.common)
	c.Options = (*OptionsService)(&c.common)
	c.Playlist = (*PlaylistService)(&c.common)
	c.Radio = (*RadioService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Track = (*TrackService)(&c.common)
	c.User = (*UserService)(&c.common)
	return c
}
