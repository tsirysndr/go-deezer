package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-deezer"
)

func main() {
	client := deezer.NewClient()
	res, _ := client.Genre.List()
	genres, _ := json.Marshal(res)
	fmt.Println(string(genres))
}
