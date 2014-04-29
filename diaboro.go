package diaboro

type Command string

type Routine []Command

type Coordinate struct {
	X, Y int
}

type Button int

type Key string

type Inputer interface {
	Click(Button) Command
	Press(Key) Command
	Move(Coordinate) Command
}

type Ability struct {
	Total, Index int
}

type Skill struct {
	Page    int
	Ability Ability
	Rune    int
}

type Passive Ability

type Build struct {
	Mouse0   Skill
	Mouse1   Skill
	Key0     Skill
	Key1     Skill
	Key2     Skill
	Key3     Skill
	Passive0 Passive
	Passive1 Passive
	Passive2 Passive
	Passive3 Passive
}

func moveSkill(from, to Skill) Routine {
	var r Routine

	direction := buttons.Right
	pageTurns := to.Page - from.Page

	if to.Page < from.Page {
		direction = buttons.Left
		pageTurns = -pageTurns
	}

	r = append(r, inputer.Move(direction))
	for i := 0; i < pageTurns; i++ {
		r = append(r, inputer.Click(1))
	}

	r = append(r, inputer.Move(to.Ability.Coordinates()))
	r = append(r, inputer.Click(1))
	r = append(r, inputer.Move(to.Rune.Coordinates()))
	r = append(r, inputer.Click(1))
	r = append(r, inputer.Move(buttons.Accept))
	r = append(r, inputer.Click(1))

	return r
}

func movePassive(p Passive) Routine {
	var r Routine

	r = append(r, inputer.Move(p.Coordinates()))
	r = append(r, inputer.Click(1))
	r = append(r, inputer.Move(buttons.Accept))
	r = append(r, inputer.Click(1))

	return r
}

func moveActives(from, to Build) Routine {
	var r Routine

	r = append(r, inputer.Move(buttons.Mouse0))
	r = append(r, moveSkill(from.Mouse0, to.Mouse0)...)

	r = append(r, inputer.Move(buttons.Mouse1))
	r = append(r, moveSkill(from.Mouse1, to.Mouse1)...)

	r = append(r, inputer.Move(buttons.Key0))
	r = append(r, moveSkill(from.Key0, to.Key0)...)

	r = append(r, inputer.Move(buttons.Key1))
	r = append(r, moveSkill(from.Key1, to.Key1)...)

	r = append(r, inputer.Move(buttons.Key2))
	r = append(r, moveSkill(from.Key2, to.Key2)...)

	r = append(r, inputer.Move(buttons.Key3))
	r = append(r, moveSkill(from.Key3, to.Key3)...)

	return r
}

func Switch(from, to Build) Routine {
	var r Routine

	r = append(r, clearActives(from, to)...)
	r = append(r, moveActives(from, to)...)

	r = append(r, clearPassives(from, to)...)
	r = append(r, movePassives(from, to)...)

	return r
}
