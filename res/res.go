package res

type Direction int

const (
	Left Direction = iota + 1
	Right
)

type Button int

const (
	Accept Button = iota + 1
	Cancel
)

type Resolution struct {
	Width, Height int
}

func (r Resolution) Normalize(c Coordinate) Coordinate {
	return c
}

type Pointer interface {
	Point(Resolution) Coordinate
}

type Passive struct {
	Total, Offset int
}

type Ability struct {
	Total, Offset int
}

type Rune struct {
	Offset int
}

func (p Passive) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			PassiveBoundingLeft,
			PassiveBoundingTop,
		},
		PassiveWidth,
		PassiveHeight,
		PassiveGutter,
		PassivePadding,
		PassiveColumns,
		p.Total,
		p.Offset,
	}.Point(r)
}

/*
	  [0] [1]
    [0] [1] [2]
  [0] [1] [2] [3]
[0] [1] [2] [3] [4]
*/
func (a Ability) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			AbilityBoundingLeft,
			AbilityBoundingTop,
		},
		AbilityWidth,
		AbilityHeight,
		AbilityGutter,
		0,
		AbilityColumns,
		a.Total,
		a.Offset,
	}.Point(r)
}

func (roune Rune) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			RuneBoundingLeft,
			RuneBoundingTop,
		},
		RuneWidth,
		RuneHeight,
		RuneGutter,
		0,
		RuneColumns,
		RuneColumns,
		roune.Offset,
	}.Point(r)
}

func (b Button) Point(r Resolution) Coordinate {
	return Coordinate{}
}

func (d Direction) Point(r Resolution) Coordinate {
	return Coordinate{}
}

type MouseSlot struct {
	Button int
}

type KeySlot MouseSlot
