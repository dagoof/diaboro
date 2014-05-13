package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/dagoof/diaboro"
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
	ability = flag.String(
		"ability",
		"zombie-bears",
		"ability to trigger",
	)
	target = flag.String(
		"target",
		"vampire-bats",
		"ability to switch to",
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

	si := blizz.NewSkillIndex(v)
	ti := blizz.NewTraitIndex(v)

	from := diaboro.Build{
		si.MustGet("firebomb"),
		si.MustGet("vampire-bats"),
		si.MustGet("burning-dogs"),
		si.MustGet("wrathful-protector"),
		si.MustGet("vengeful-spirit"),
		si.MustGet("jaunt"),
		ti.MustGet("fetish-sycophants"),
		ti.MustGet("grave-injustice"),
		ti.MustGet("midnight-feast"),
		ti.MustGet("pierce-the-veil"),
	}

	to := diaboro.Build{
		si.MustGet("well-of-souls"),
		si.MustGet("zombie-bears"),
		si.MustGet("piranhado"),
		si.MustGet("mass-hysteria"),
		si.MustGet("vengeful-spirit"),
		si.MustGet("jaunt"),
		ti.MustGet("spirit-vessel"),
		ti.MustGet("grave-injustice"),
		ti.MustGet("rush-of-essence"),
		ti.MustGet("spiritual-attunement"),
	}

	for _, part := range diaboro.Switch(from, to) {
		fmt.Println(part)
		part.Execute()

		time.Sleep(time.Millisecond * 500)
	}
}
