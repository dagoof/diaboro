package blizz

import (
	"strings"
)

func Normalize(s string) string {
	return strings.ToLower(
		strings.Join(strings.Fields(s), "-"),
	)
}

type Rune struct {
	RequiredLevel int
	Name          string
	Type          string
}

type Skill struct {
	RequiredLevel int
	OrderIndex    int
	CategoryName  string
	CategorySlug  string
	Name          string
	Icon          string
	Slug          string
	Runes         []Rune
}

type Trait struct {
	RequiredLevel int
	OrderIndex    int
	Name          string
	Icon          string
	Slug          string
}

type Calculator struct {
	Skills []Skill
	Traits []Trait
}

type Offset int

type TraitOffset struct {
	Row, Column Offset
}
