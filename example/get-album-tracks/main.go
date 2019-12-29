package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Album.GetTracks("302127")
	tracks, _ := json.Marshal(res)
	fmt.Println(string(tracks))
}
