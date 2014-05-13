package blizz

import (
	"github.com/dagoof/diaboro/res"
)

type SkillMeta struct {
	Skill      Skill
	Page       int
	Total      int
	Offset     int
	RuneOffset int
}

func (s *SkillMeta) Ability() res.Ability {
	return res.Ability{s.Total, s.Offset}
}

func (s *SkillMeta) Rune() res.Rune {
	return res.Rune{s.RuneOffset}
}

type TraitMeta struct {
	trait         Trait
	Total, Offset int
}

type TraitIndex struct {
	traits []Trait
	index  map[string]TraitMeta
}

type SkillIndex struct {
	skills [][]Skill
	index  map[string]SkillMeta
}

func (i SkillIndex) MustGet(slug string) SkillMeta {
	return i.index[slug]
}

func (i SkillIndex) Get(slug string) (SkillMeta, bool) {
	m, ok := i.index[slug]
	return m, ok
}

func NewSkillIndex(c Calculator) SkillIndex {
	var skills [][]Skill
	categories := map[string]int{}

	for _, skill := range c.Skills {
		category := skill.CategorySlug
		if _, ok := categories[category]; !ok {
			categories[category] = len(skills)
			skills = append(skills, []Skill{})
		}

		skills[categories[category]] = append(
			skills[categories[category]],
			skill,
		)
	}

	index := map[string]SkillMeta{}
	for i, ss := range skills {
		for j, skill := range ss {
			for k, roune := range skill.Runes {
				index[Normalize(roune.Name)] = SkillMeta{
					skill,
					i,
					len(ss),
					j,
					k,
				}
			}
		}
	}

	return SkillIndex{skills, index}
}
