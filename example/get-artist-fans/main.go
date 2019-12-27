package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Artist.GetFans("27")
	fans, _ := json.Marshal(res)
	fmt.Println(string(fans))
}
