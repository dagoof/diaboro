package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"

	"github.com/dagoof/diaboro/blizz"
)

const (
	url = "http://us.battle.net/d3/en/data/calculator/"
)

var (
	class = flag.String(
		"class",
		"witch-doctor",
		"class to gather information for",
	)
)

func init() {
	flag.Parse()
}

func main() {
	var v blizz.Calculator

	req, err := http.Get(url + *class)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(req.Body).Decode(&v)
	if err != nil {
		panic(err)
	}

	index := blizz.NewSkillIndex(v)
	body, err := json.MarshalIndent(
		index.MustGet("black-blood"),
		"", "  ",
	)

	if err != nil {
		panic(err)
	}

	os.Stdout.Write(body)

}
