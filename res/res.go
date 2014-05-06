package res

type Direction int

const (
	Left Direction = iota
	Right
)

type Button int

const (
	Accept Button = iota
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

type Ability struct {
	Total, Offset int
}

type Rune struct {
	Offset int
}

type Passive struct {
	Total, Offset int
}

type MouseSlot struct {
	Offset int
}

type KeySlot struct {
	Offset int
}

type PassiveSlot struct {
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
	return Grid{
		Coordinate{
			ButtonBoundingLeft,
			ButtonBoundingTop,
		},
		ButtonWidth,
		ButtonHeight,
		ButtonGutter,
		0,
		ButtonColumns,
		ButtonColumns,
		int(b),
	}.Point(r)
}

func (d Direction) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			DirectionBoundingLeft,
			DirectionBoundingTop,
		},
		DirectionWidth,
		DirectionHeight,
		DirectionGutter,
		0,
		DirectionColumns,
		DirectionColumns,
		int(d),
	}.Point(r)
}

func (s MouseSlot) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			MouseSlotBoundingLeft,
			MouseSlotBoundingTop,
		},
		MouseSlotWidth,
		MouseSlotHeight,
		MouseSlotGutter,
		0,
		MouseSlotColumns,
		MouseSlotColumns,
		s.Offset,
	}.Point(r)
}

func (s KeySlot) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			KeySlotBoundingLeft,
			KeySlotBoundingTop,
		},
		KeySlotWidth,
		KeySlotHeight,
		KeySlotGutter,
		KeySlotPadding,
		KeySlotColumns,
		KeySlotTotal,
		s.Offset,
	}.Point(r)
}

func (s PassiveSlot) Point(r Resolution) Coordinate {
	return Grid{
		Coordinate{
			PassiveSlotBoundingLeft,
			PassiveSlotBoundingTop,
		},
		PassiveSlotWidth,
		PassiveSlotHeight,
		PassiveSlotGutter,
		0,
		PassiveSlotColumns,
		PassiveSlotColumns,
		s.Offset,
	}.Point(r)
}
