package diaboro

import (
	"github.com/dagoof/diaboro/blizz"
	"github.com/dagoof/diaboro/res"
)

type Command interface {
	Execute() error
}

type Routine []Command

type Button int

type Key string

type Inputer interface {
	Click(Button) Command
	Press(Key) Command
	Move(res.Pointer) Command
}

type Ability struct {
	Total, Index int
}

type Passive Ability

type Build struct {
	Mouse0   blizz.SkillMeta
	Mouse1   blizz.SkillMeta
	Key0     blizz.SkillMeta
	Key1     blizz.SkillMeta
	Key2     blizz.SkillMeta
	Key3     blizz.SkillMeta
	Passive0 blizz.TraitMeta
	Passive1 blizz.TraitMeta
	Passive2 blizz.TraitMeta
	Passive3 blizz.TraitMeta
}

func moveSkill(from, to blizz.SkillMeta) Routine {
	var r Routine

	direction := res.Right
	pageTurns := to.Page - from.Page

	if to.Page < from.Page {
		direction = res.Left
		pageTurns = -pageTurns
	}

	r = append(r, TheInputer.Move(direction))
	for i := 0; i < pageTurns; i++ {
		r = append(r, TheInputer.Click(1))
	}

	r = append(r, TheInputer.Move(to.Ability()))
	r = append(r, TheInputer.Click(1))
	r = append(r, TheInputer.Move(to.Rune()))
	r = append(r, TheInputer.Click(1))
	r = append(r, TheInputer.Move(res.Accept))
	r = append(r, TheInputer.Click(1))

	return r
}

func MoveSkill(from, to blizz.SkillMeta) Routine {
	return moveSkill(from, to)
}

/*
func movePassive(p Passive) Routine {
	var r Routine

	r = append(r, TheInputer.Move(p.Coordinates()))
	r = append(r, TheInputer.Click(1))
	r = append(r, TheInputer.Move(buttons.Accept))
	r = append(r, TheInputer.Click(1))

	return r
}
*/

func moveActives(from, to Build) Routine {
	var r Routine

	r = append(r, TheInputer.Move(res.MouseSlot{0}))
	r = append(r, moveSkill(from.Mouse0, to.Mouse0)...)

	r = append(r, TheInputer.Move(res.MouseSlot{1}))
	r = append(r, moveSkill(from.Mouse1, to.Mouse1)...)

	r = append(r, TheInputer.Move(res.KeySlot{0}))
	r = append(r, moveSkill(from.Key0, to.Key0)...)

	r = append(r, TheInputer.Move(res.KeySlot{1}))
	r = append(r, moveSkill(from.Key1, to.Key1)...)

	r = append(r, TheInputer.Move(res.KeySlot{2}))
	r = append(r, moveSkill(from.Key2, to.Key2)...)

	r = append(r, TheInputer.Move(res.KeySlot{3}))
	r = append(r, moveSkill(from.Key3, to.Key3)...)

	return r
}

func Switch(from, to Build) Routine {
	var r Routine

	//r = append(r, clearActives(from, to)...)
	r = append(r, moveActives(from, to)...)

	//r = append(r, clearPassives(from, to)...)
	//r = append(r, movePassives(from, to)...)

	return r
}
