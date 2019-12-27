package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetRadio("27")
	radio, _ := json.Marshal(res)
	fmt.Println(string(radio))
}
