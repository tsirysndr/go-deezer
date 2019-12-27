package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetPlaylists("27")
	playlists, _ := json.Marshal(res)
	fmt.Println(string(playlists))
}
