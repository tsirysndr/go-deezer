package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Chart.GetAlbums()
	albums, _ := json.Marshal(res)
	fmt.Println(string(albums))
}
