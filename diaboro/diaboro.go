package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dagoof/diaboro"
	"github.com/dagoof/diaboro/blizz"
	"github.com/dagoof/diaboro/res"
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
	ability = flag.String(
		"ability",
		"zombie-bears",
		"ability to trigger",
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
	abil := index.MustGet(*ability)

	inp := diaboro.TheInputer

	fmt.Println(inp.Click(1))
	inp.Move(res.Ability{abil.Total, abil.Offset}).Execute()
	time.Sleep(time.Second)
	inp.Move(res.Rune{abil.Rune}).Execute()
}
