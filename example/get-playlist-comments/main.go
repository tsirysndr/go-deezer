package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Playlist.GetComments("908622995")
	comments, _ := json.Marshal(res)
	fmt.Println(string(comments))
}
