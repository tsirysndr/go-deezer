package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.Get("27")
	artist, _ := json.Marshal(res)
	fmt.Println(string(artist))
}
