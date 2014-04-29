package diaboro

/*
A Position can be transformed into a coordinate according to a known Resolution.

res := Resolution{1920, 1200}
pos := Ability_3_0

input.Move(res.Adjust(pos))
*/

type Position Coordinate

var (
	Left
	Right

	Ability_3_0
	Ability_3_1
	Ability_3_2

	Ability_3_0
	Ability_3_1
	Ability_3_2

	Ability_3_0
	Ability_3_1
	Ability_3_2
)
