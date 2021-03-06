package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Album.Get("302127")
	album, _ := json.Marshal(res)
	fmt.Println(string(album))
}
