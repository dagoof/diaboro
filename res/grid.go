package res

type BoundingBox struct {
	Left, Top, Width Offset
}

type Grid struct {
	BoundingBox

	Width, Height         Offset
	Gutter, Padding       Offset
	Columns, Total, Index int
}

func ChangeMe(index int, width, gutter Offset) Offset {
	if index == 0 {
		return 0
	}

	oIndex := Offset(index)

	return oIndex*width + Offset(index-1)*gutter
}

func (g Grid) Centered(r Resolution, left, top Offset) Coordinate {
	return Coordinate{
		r.X(left + (g.Width / Offset(2))),
		r.Y(top + (g.Height / Offset(2))),
	}
}

func (g Grid) Point(r Resolution) Coordinate {
	totalWidth := ChangeMe(g.Total%g.Columns, g.Width, g.Gutter)
	offsetLeft := ChangeMe(g.Index%g.Columns, g.Width, g.Gutter)

	left := (g.BoundingBox.Width-totalWidth)/2 + offsetLeft
	top := Offset(g.Index / g.Columns) * (g.Height + g.Padding)

	return g.Centered(r, left, top)
}
