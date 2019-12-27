package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetComments("27")
	comments, _ := json.Marshal(res)
	fmt.Println(string(comments))
}
