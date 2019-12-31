package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Playlist.GetTracks("908622995")
	tracks, _ := json.Marshal(res)
	fmt.Println(string(tracks))
}
