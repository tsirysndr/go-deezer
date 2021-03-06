package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Album.GetFans("302127")
	fans, _ := json.Marshal(res)
	fmt.Println(string(fans))
}
