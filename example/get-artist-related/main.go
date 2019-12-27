package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetRelated("27")
	related, _ := json.Marshal(res)
	fmt.Println(string(related))
}
