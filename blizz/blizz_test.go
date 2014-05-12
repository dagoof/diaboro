package blizz

import (
	"encoding/json"
	"net/http"
	"testing"
)

const (
	url         = "http://us.battle.net/d3/en/data/calculator/"
	witchDoctor = "witch-doctor"
	demonHunter = "demon-hunter"
	monk        = "monk"
	barbarian   = "barbarian"
	crusader    = "crusader"
)

func BenchmarkSkillIndex(b *testing.B) {
	var v Calculator

	req, err := http.Get(url + witchDoctor)
	if err != nil {
		b.Fatal(err)
	}

	err = json.NewDecoder(req.Body).Decode(&v)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		NewSkillIndex(v)
	}

}
