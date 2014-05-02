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

func (r Resolution) X(n float64) {
	return r.Width * n
}


func (r Resolution) Y(n float64) {
	return r.Height * n
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
	blat := func(index, width, gutter int) int {
		if index == 0 {
			return 0
		}

		return index * width + (index - 1) * gutter
	}

	totalWidth := blat(a.Total, AbilityWidth, AbilityGutter)
	leftOffset := blat(a.Index, AbilityWidth, AbilityGutter)

	totalLeft := (AbilityBoundWidth - totalWidth) / 2 + leftOffset

	return Centered(totalLeft, AbilityTopOffset, AbilityWidth, AbilityHeight)
}

func (r Rune) Point(r Resolution) Coordinate {
}

func (b Button) Point(r Resolution) Coordinate {
}

func (d Direction) Point(r Resolution) Coordinate {
}

type MouseSlot struct {
	Button int
}

type KeySlot MouseSlot

type PassiveSlot
