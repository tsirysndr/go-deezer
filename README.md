<h1 align="center">Welcome to go-deezer 👋 *** Work In Progress ***</h1>
<p align="center">
  <a href="https://github.com/tsirysndr/go-spotify/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-deezer" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-deezer">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-deezer">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-deezer">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-deezer">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-deezer">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-deezer">
  <a href="https://github.com/tsirysndr/go-deezer/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

go-deezer is a Go client library for accessing the [Deezer API](https://developers.deezer.com/api)

## 🚚 Install

```sh
go get github.com/tsirysndr/go-deezer
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-deezer"
```

Construct a new Deezer client, then use the various services on the client to access different parts of the Deezer API. For example:

```Go
client := deezer.NewClient()
res, _ := client.Artist.Get("27")
artist, _ := json.Marshal(res)
fmt.Println(string(artist))
```

## ✨ Coverage

Currently the following services are supported:

- [x] Get an Album
- [x] Return a list of album's comments
- [x] Return a list of album's fans
- [x] Return a list of album's tracks
- [x] Get an Artist
- [x] Get the top 5 tracks of an artist
- [x] Return a list of artist's albums
- [x] Return a list of artist's comments
- [x] Return a list of artist's fans
- [x] Return a list of related artists
- [x] Return a list of tracks
- [x] Return a list of artist's playlist
- [x] Returns the Top tracks
- [x] Returns the Top albums
- [x] Returns the Top artists
- [x] Returns the Top playlists
- [ ] Returns the Top podcasts
- [ ] Remove a comment
- [ ] Add a comment to the album
- [ ] Add a comment to the artist
- [ ] Add a comment to the playlist
- [ ] Return a list of albums selected every week by the Deezer Team
- [x] Returns all artists for a genre
- [x] Get a Genre
- [ ] Returns all radios for a genre
- [ ] Get the information about the API in the current country
- [ ] Get the user's options
- [x] Get a playlist
- [ ] Rate the playlist
- [ ] Update the playlist
- [ ] Delete the playlist
- [ ] Add a playlist to the folder
- [ ] Remove a playlist from the folder
- [ ] Create a playlist
- [ ] Add a playlist to the user's favorites
- [ ] Remove a playlist from the user's favorites
- [x] Return a list of playlist's comments
- [x] Return a list of playlist's fans
- [x] Return a list of playlist's tracks
- [ ] Return a list of playlist's recommendation tracks
- [ ] Get a radio
- [ ] Add a radio to the user's favorites
- [ ] Remove a radio from the user's favorites
- [ ] Returns a list of radio splitted by genre
- [ ] Return the top radios (25 radios)
- [ ] Get first 40 tracks in the radio
- [ ] Returns a list of personal radio splitted by genre (as MIX in website)
- [ ] Search tracks
- [x] Get track
- [ ] Update a personal track
- [ ] Delete a personal track
- [ ] Add a track to the playlist
- [ ] Order tracks in the playlist
- [ ] Remove tracks from the playlist
- [ ] Add a track to the user's favorites
- [ ] Remove a track from the user's favorites
- [ ] Return a list of user's favorite albums
- [ ] Return a list of user's favorite artists
- [ ] Returns a list of the user's top 25 tracks
- [ ] Returns a list of the user's top albums 
- [ ] Returns a list of the user's top playlists
- [ ] Returns a list of the user's top artists

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!

