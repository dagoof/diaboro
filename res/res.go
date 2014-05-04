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

type Coordinate struct {
	X, Y int
}

type Resolution struct {
	Width, Height int
}

func (r Resolution) Scale() Offset {
	return Offset(r.Height) * 0.8
}

func (r Resolution) X(n Offset) int {
	return (r.Width - r.Height) + int(r.Scale() * n)
}

func (r Resolution) Y(n Offset) int {
	return int(r.Scale() * n)
}

type Pointer interface {
	Point(Resolution) Coordinate
}

type Passive struct {
	Row, Offset int
}

type Ability struct {
	Total, Offset int
}

type Rune struct {
	Offset int
}

func (p Passive) Point(r Resolution) Coordinate {
	return Coordinate{}
}

/*

	  [0] [1]
    [0] [1] [2]
  [0] [1] [2] [3]
[0] [1] [2] [3] [4]

total width = total * width + (total - 1) * gutter
left offset = index * width + (index - 1) * gutter

top-left = (total-width / 2)

*/
func (a Ability) Point(r Resolution) Coordinate {
	return Grid{
		BoundingBox{
			AbilityBoundingLeft,
			AbilityBoundingTop,
			AbilityBoundingWidth,
		},
		AbilityWidth,
		AbilityHeight,
		AbilityGutter,
		0,
		a.Total,
		a.Total,
		a.Offset,
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

