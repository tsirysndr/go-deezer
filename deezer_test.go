package deezer

import (
	"encoding/json"
	"testing"
)

func TestTrackUnmarshal(t *testing.T) {
	raw := `{
		"id": 3135556,
		"readable": true,
		"title": "Harder, Better, Faster, Stronger",
		"title_short": "Harder, Better, Faster, Stronger",
		"title_version": "",
		"isrc": "GBDUW0000059",
		"duration": 226,
		"rank": 814839,
		"explicit_lyrics": false,
		"explicit_content_lyrics": 0,
		"explicit_content_cover": 0,
		"preview": "https://cdnt-preview.dzcdn.net/api/...",
		"artist": {
			"id": 27,
			"name": "Daft Punk",
			"tracklist": "https://api.deezer.com/artist/27/top?limit=50",
			"type": "artist"
		},
		"album": {
			"id": 302127,
			"title": "Discovery",
			"cover": "https://api.deezer.com/album/302127/image",
			"type": "album"
		},
		"type": "track"
	}`
	var track Track
	if err := json.Unmarshal([]byte(raw), &track); err != nil {
		t.Fatal("Track unmarshal:", err)
	}
	if track.ID != 3135556 {
		t.Errorf("Track.ID = %d, want 3135556", track.ID)
	}
	if track.Title != "Harder, Better, Faster, Stronger" {
		t.Errorf("Track.Title = %q, want Harder, Better, Faster, Stronger", track.Title)
	}
	if track.Duration != 226 {
		t.Errorf("Track.Duration = %d, want 226", track.Duration)
	}
	if track.Artist == nil || track.Artist.ID != 27 {
		t.Errorf("Track.Artist.ID = %v, want 27", track.Artist)
	}
	if track.Album == nil || track.Album.ID != 302127 {
		t.Errorf("Track.Album.ID = %v, want 302127", track.Album)
	}
}

func TestAlbumUnmarshal(t *testing.T) {
	raw := `{
		"id": 302127,
		"title": "Discovery",
		"upc": "724384960650",
		"genre_id": 106,
		"nb_tracks": 14,
		"duration": 3662,
		"fans": 336027,
		"release_date": "2001-03-07",
		"record_type": "album",
		"available": true,
		"explicit_lyrics": false,
		"artist": {
			"id": 27,
			"name": "Daft Punk",
			"type": "artist"
		},
		"type": "album"
	}`
	var album Album
	if err := json.Unmarshal([]byte(raw), &album); err != nil {
		t.Fatal("Album unmarshal:", err)
	}
	if album.ID != 302127 {
		t.Errorf("Album.ID = %d, want 302127", album.ID)
	}
	if album.Title != "Discovery" {
		t.Errorf("Album.Title = %q, want Discovery", album.Title)
	}
	if album.NbTracks != 14 {
		t.Errorf("Album.NbTracks = %d, want 14", album.NbTracks)
	}
	if album.Artist == nil || album.Artist.ID != 27 {
		t.Errorf("Album.Artist.ID = %v, want 27", album.Artist)
	}
}

func TestArtistUnmarshal(t *testing.T) {
	raw := `{
		"id": 27,
		"name": "Daft Punk",
		"nb_album": 38,
		"nb_fan": 5170167,
		"radio": true,
		"tracklist": "https://api.deezer.com/artist/27/top?limit=50",
		"type": "artist"
	}`
	var artist Artist
	if err := json.Unmarshal([]byte(raw), &artist); err != nil {
		t.Fatal("Artist unmarshal:", err)
	}
	if artist.ID != 27 {
		t.Errorf("Artist.ID = %d, want 27", artist.ID)
	}
	if artist.Name != "Daft Punk" {
		t.Errorf("Artist.Name = %q, want Daft Punk", artist.Name)
	}
	if artist.NbAlbum != 38 {
		t.Errorf("Artist.NbAlbum = %d, want 38", artist.NbAlbum)
	}
	if artist.NbFan != 5170167 {
		t.Errorf("Artist.NbFan = %d, want 5170167", artist.NbFan)
	}
}

func TestPlaylistUnmarshal(t *testing.T) {
	raw := `{
		"id": 908622995,
		"title": "En mode 60",
		"public": true,
		"nb_tracks": 50,
		"fans": 22034,
		"checksum": "bd6fb0f18f3dc8af04b462d74220c3a3",
		"creation_date": "2014-06-27 11:09:31",
		"type": "playlist"
	}`
	var playlist Playlist
	if err := json.Unmarshal([]byte(raw), &playlist); err != nil {
		t.Fatal("Playlist unmarshal:", err)
	}
	if playlist.ID != 908622995 {
		t.Errorf("Playlist.ID = %d, want 908622995", playlist.ID)
	}
	if playlist.Title != "En mode 60" {
		t.Errorf("Playlist.Title = %q, want En mode 60", playlist.Title)
	}
	if !playlist.Public {
		t.Error("Playlist.Public = false, want true")
	}
}

func TestGenreUnmarshal(t *testing.T) {
	raw := `{
		"id": 0,
		"name": "All",
		"picture": "https://api.deezer.com/genre/0/image",
		"type": "genre"
	}`
	var genre Genre
	if err := json.Unmarshal([]byte(raw), &genre); err != nil {
		t.Fatal("Genre unmarshal:", err)
	}
	if genre.ID != 0 {
		t.Errorf("Genre.ID = %d, want 0", genre.ID)
	}
	if genre.Name != "All" {
		t.Errorf("Genre.Name = %q, want All", genre.Name)
	}
}

func TestUserUnmarshal(t *testing.T) {
	raw := `{
		"id": 1687950,
		"name": "Alexandre - Pop & Hits Editor",
		"tracklist": "https://api.deezer.com/user/1687950/flow",
		"type": "user"
	}`
	var user User
	if err := json.Unmarshal([]byte(raw), &user); err != nil {
		t.Fatal("User unmarshal:", err)
	}
	if user.ID != 1687950 {
		t.Errorf("User.ID = %d, want 1687950", user.ID)
	}
	if user.Name != "Alexandre - Pop & Hits Editor" {
		t.Errorf("User.Name = %q, want Alexandre - Pop & Hits Editor", user.Name)
	}
}

func TestContributorUnmarshal(t *testing.T) {
	raw := `{
		"id": 27,
		"name": "Daft Punk",
		"radio": true,
		"type": "artist",
		"role": "Main"
	}`
	var c Contributor
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatal("Contributor unmarshal:", err)
	}
	if c.ID != 27 {
		t.Errorf("Contributor.ID = %d, want 27", c.ID)
	}
	if c.Role != "Main" {
		t.Errorf("Contributor.Role = %q, want Main", c.Role)
	}
	if !c.Radio {
		t.Error("Contributor.Radio = false, want true")
	}
}

func TestGenresUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{"id": 0, "name": "All", "type": "genre"},
			{"id": 132, "name": "Pop", "type": "genre"},
			{"id": 116, "name": "Rap/Hip Hop", "type": "genre"}
		]
	}`
	var gs Genres
	if err := json.Unmarshal([]byte(raw), &gs); err != nil {
		t.Fatal("Genres unmarshal:", err)
	}
	if len(gs.Data) != 3 {
		t.Errorf("Genres.Data len = %d, want 3", len(gs.Data))
	}
	if gs.Data[1].Name != "Pop" {
		t.Errorf("Genres.Data[1].Name = %q, want Pop", gs.Data[1].Name)
	}
}

func TestTopUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 66609426,
				"title": "Get Lucky (Radio Edit - feat. Pharrell Williams and Nile Rodgers)",
				"duration": 248,
				"rank": 954234,
				"artist": {"id": 27, "name": "Daft Punk", "type": "artist"},
				"type": "track"
			}
		],
		"total": 1
	}`
	var top Top
	if err := json.Unmarshal([]byte(raw), &top); err != nil {
		t.Fatal("Top unmarshal:", err)
	}
	if len(top.Data) != 1 {
		t.Errorf("Top.Data len = %d, want 1", len(top.Data))
	}
	if top.Data[0].ID != 66609426 {
		t.Errorf("Top.Data[0].ID = %d, want 66609426", top.Data[0].ID)
	}
	if top.Total != 1 {
		t.Errorf("Top.Total = %d, want 1", top.Total)
	}
}

func TestAlbumsUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 494309801,
				"title": "Random Access Memories (Drumless Edition)",
				"record_type": "album",
				"type": "album"
			}
		],
		"total": 10
	}`
	var albums Albums
	if err := json.Unmarshal([]byte(raw), &albums); err != nil {
		t.Fatal("Albums unmarshal:", err)
	}
	if len(albums.Data) != 1 {
		t.Errorf("Albums.Data len = %d, want 1", len(albums.Data))
	}
	if albums.Data[0].ID != 494309801 {
		t.Errorf("Albums.Data[0].ID = %d, want 494309801", albums.Data[0].ID)
	}
}

func TestArtistsUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 6404,
				"name": "Justice",
				"nb_fan": 801480,
				"radio": true,
				"type": "artist"
			}
		],
		"total": 2
	}`
	var artists Artists
	if err := json.Unmarshal([]byte(raw), &artists); err != nil {
		t.Fatal("Artists unmarshal:", err)
	}
	if len(artists.Data) != 1 {
		t.Errorf("Artists.Data len = %d, want 1", len(artists.Data))
	}
	if artists.Data[0].Name != "Justice" {
		t.Errorf("Artists.Data[0].Name = %q, want Justice", artists.Data[0].Name)
	}
}

func TestTracksUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 3135553,
				"title": "One More Time",
				"duration": 320,
				"rank": 931282,
				"artist": {"id": 27, "name": "Daft Punk", "type": "artist"},
				"type": "track"
			}
		],
		"total": 14,
		"next": "https://api.deezer.com/album/302127/tracks?index=25"
	}`
	var tracks Tracks
	if err := json.Unmarshal([]byte(raw), &tracks); err != nil {
		t.Fatal("Tracks unmarshal:", err)
	}
	if len(tracks.Data) != 1 {
		t.Errorf("Tracks.Data len = %d, want 1", len(tracks.Data))
	}
	if tracks.Data[0].ID != 3135553 {
		t.Errorf("Tracks.Data[0].ID = %d, want 3135553", tracks.Data[0].ID)
	}
	if tracks.Next == "" {
		t.Error("Tracks.Next is empty, expected pagination URL")
	}
}

func TestPlaylistsUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 908622995,
				"title": "En mode 60",
				"public": true,
				"type": "playlist"
			}
		],
		"total": 1
	}`
	var pls Playlists
	if err := json.Unmarshal([]byte(raw), &pls); err != nil {
		t.Fatal("Playlists unmarshal:", err)
	}
	if len(pls.Data) != 1 {
		t.Errorf("Playlists.Data len = %d, want 1", len(pls.Data))
	}
	if pls.Data[0].Title != "En mode 60" {
		t.Errorf("Playlists.Data[0].Title = %q, want En mode 60", pls.Data[0].Title)
	}
}

func TestRelatedUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 6404,
				"name": "Justice",
				"nb_fan": 801480,
				"radio": true,
				"type": "artist"
			}
		],
		"total": 20
	}`
	var rel Related
	if err := json.Unmarshal([]byte(raw), &rel); err != nil {
		t.Fatal("Related unmarshal:", err)
	}
	if len(rel.Data) != 1 {
		t.Errorf("Related.Data len = %d, want 1", len(rel.Data))
	}
	if rel.Data[0].Name != "Justice" {
		t.Errorf("Related.Data[0].Name = %q, want Justice", rel.Data[0].Name)
	}
}

func TestRadioUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": 68442677,
				"title": "Get Lucky (Daft Punk Remix)",
				"duration": 631,
				"rank": 281313,
				"artist": {"id": 27, "name": "Daft Punk", "type": "artist"},
				"type": "track"
			}
		],
		"total": 25
	}`
	var radio Radio
	if err := json.Unmarshal([]byte(raw), &radio); err != nil {
		t.Fatal("Radio unmarshal:", err)
	}
	if len(radio.Data) != 1 {
		t.Errorf("Radio.Data len = %d, want 1", len(radio.Data))
	}
	if radio.Data[0].Title != "Get Lucky (Daft Punk Remix)" {
		t.Errorf("Radio.Data[0].Title = %q, want Get Lucky (Daft Punk Remix)", radio.Data[0].Title)
	}
}

func TestFansUnmarshal(t *testing.T) {
	raw := `{"data": [], "total": 0}`
	var fans Fans
	if err := json.Unmarshal([]byte(raw), &fans); err != nil {
		t.Fatal("Fans unmarshal:", err)
	}
	if len(fans.Data) != 0 {
		t.Errorf("Fans.Data len = %d, want 0", len(fans.Data))
	}
}

func TestCommentUnmarshal(t *testing.T) {
	raw := `{
		"id": 1,
		"text": "Great track!",
		"date": 1609459200,
		"author": {"id": 1, "name": "User", "type": "user"},
		"type": "comment"
	}`
	var c Comment
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatal("Comment unmarshal:", err)
	}
	if c.ID != 1 {
		t.Errorf("Comment.ID = %d, want 1", c.ID)
	}
	if c.Text != "Great track!" {
		t.Errorf("Comment.Text = %q, want Great track!", c.Text)
	}
}

func TestAlbumWithGenresUnmarshal(t *testing.T) {
	raw := `{
		"id": 302127,
		"title": "Discovery",
		"genres": {
			"data": [
				{"id": 106, "name": "Electro", "type": "genre"}
			]
		},
		"type": "album"
	}`
	var album Album
	if err := json.Unmarshal([]byte(raw), &album); err != nil {
		t.Fatal("Album with genres unmarshal:", err)
	}
	if album.Genres == nil {
		t.Fatal("Album.Genres is nil")
	}
	if len(album.Genres.Data) != 1 {
		t.Errorf("Album.Genres.Data len = %d, want 1", len(album.Genres.Data))
	}
	if album.Genres.Data[0].Name != "Electro" {
		t.Errorf("Album.Genres.Data[0].Name = %q, want Electro", album.Genres.Data[0].Name)
	}
}

func TestTrackWithContributorsUnmarshal(t *testing.T) {
	raw := `{
		"id": 66609426,
		"title": "Get Lucky",
		"contributors": [
			{
				"id": 27,
				"name": "Daft Punk",
				"radio": true,
				"type": "artist",
				"role": "Main"
			}
		],
		"type": "track"
	}`
	var track Track
	if err := json.Unmarshal([]byte(raw), &track); err != nil {
		t.Fatal("Track with contributors unmarshal:", err)
	}
	if len(track.Contributors) != 1 {
		t.Errorf("Track.Contributors len = %d, want 1", len(track.Contributors))
	}
	if track.Contributors[0].Name != "Daft Punk" {
		t.Errorf("Track.Contributors[0].Name = %q, want Daft Punk", track.Contributors[0].Name)
	}
}

func TestAlbumWithTracksUnmarshal(t *testing.T) {
	raw := `{
		"id": 302127,
		"title": "Discovery",
		"tracks": {
			"data": [
				{
					"id": 3135553,
					"title": "One More Time",
					"artist": {"id": 27, "name": "Daft Punk", "type": "artist"},
					"type": "track"
				}
			]
		},
		"type": "album"
	}`
	var album Album
	if err := json.Unmarshal([]byte(raw), &album); err != nil {
		t.Fatal("Album with tracks unmarshal:", err)
	}
	if album.Tracks == nil {
		t.Fatal("Album.Tracks is nil")
	}
	if len(album.Tracks.Data) != 1 {
		t.Errorf("Album.Tracks.Data len = %d, want 1", len(album.Tracks.Data))
	}
	if album.Tracks.Data[0].Title != "One More Time" {
		t.Errorf("Album.Tracks.Data[0].Title = %q, want One More Time", album.Tracks.Data[0].Title)
	}
}
