package res

type Grid struct {
	BoundingBox Coordinate

	Width, Height         int
	Gutter, Padding       int
	Columns, Total, Index int
}

func TotalRowWidth(index, width, gutter int) int {
	if index == 0 {
		return 0
	}

	return index*width + (index-1)*gutter
}

func RowOffset(index, width, gutter int) int {
	return index*width + index*gutter
}

func (g Grid) Centered(c Coordinate) Coordinate {
	return Coordinate{
		c.X + (g.Width / 2),
		c.Y + (g.Height / 2),
	}
}

func (g Grid) BoxWidth() int {
	return TotalRowWidth(g.Columns, g.Width, g.Gutter)
}

func (g Grid) Point(r Resolution) Coordinate {
	row := g.Index / g.Columns
	col := g.Index % g.Columns

	// If we are on the last row, we may have less cols
	cols := g.Columns
	if row == (g.Total / g.Columns) {
		cols = g.Total % g.Columns
	}

	totalWidth := TotalRowWidth(cols, g.Width, g.Gutter)
	offsetLeft := RowOffset(col, g.Width, g.Gutter)

	x := (g.BoxWidth()-totalWidth)/2 + offsetLeft
	y := row * (g.Height + g.Padding)

	c := Coordinate{x, y}
	c = g.BoundingBox.Shift(c)
	c = g.Centered(c)
	return r.Normalize(c)
}
