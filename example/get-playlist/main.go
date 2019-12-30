package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Playlist.Get("908622995")
	playlist, _ := json.Marshal(res)
	fmt.Println(string(playlist))
}
