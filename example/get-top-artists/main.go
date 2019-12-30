package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Chart.GetArtists()
	artists, _ := json.Marshal(res)
	fmt.Println(string(artists))
}
