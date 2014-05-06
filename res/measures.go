package res

type Offset int

const (
	/*
		Width
		Height
		Gutter
		Padding
		Left
		Top
		Bounding
	*/
	ScreenWidth  = 1920
	ScreenHeight = 1200
	WindowTop    = 42
	WindowBottom = 974
	WindowLeft   = 410

	AbilityBoundingLeft = 602
	AbilityBoundingTop  = 185
	AbilityWidth        = 75
	AbilityHeight       = 74
	AbilityGutter       = 86
	AbilityColumns      = 6

	PassiveBoundingLeft = 483
	PassiveBoundingTop  = 315
	PassiveWidth        = 87
	PassiveHeight       = 83
	PassiveGutter       = 154
	PassivePadding      = 32
	PassiveColumns      = 4

	RuneBoundingLeft = 521
	RuneBoundingTop  = 442
	RuneWidth        = 78
	RuneHeight       = 77
	RuneGutter       = 82
	RuneColumns      = 6

	ButtonBoundingLeft = 703
	ButtonBoundingTop  = 880
	ButtonWidth        = 240
	ButtonHeight       = 46
	ButtonGutter       = 33
	ButtonColumns      = 2

	DirectionBoundingLeft = 491
	DirectionBoundingTop  = 194
	DirectionWidth        = 69
	DirectionHeight       = 61
	DirectionGutter       = 803
	DirectionColumns      = 2

	MouseSlotBoundingLeft = 555
	MouseSlotBoundingTop  = 185
	MouseSlotWidth        = 407
	MouseSlotHeight       = 98
	MouseSlotGutter       = 12
	MouseSlotColumns      = 2

	KeySlotBoundingLeft = MouseSlotBoundingLeft
	KeySlotBoundingTop  = 377
	KeySlotWidth        = MouseSlotWidth
	KeySlotHeight       = MouseSlotHeight
	KeySlotGutter       = MouseSlotGutter
	KeySlotPadding      = 63
	KeySlotColumns      = MouseSlotColumns
	KeySlotTotal        = 4
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Unshift(ref Coordinate) Coordinate {
	return Coordinate{
		c.X - ref.X,
		c.Y - ref.Y,
	}
}

func (c Coordinate) Shift(ref Coordinate) Coordinate {
	return Coordinate{
		c.X + ref.X,
		c.Y + ref.Y,
	}
}

func Unshift(c Coordinate) Coordinate {
	return c.Unshift(Coordinate{
		WindowLeft,
		WindowTop,
	})
}

func Shift(c Coordinate) Coordinate {
	return c.Shift(Coordinate{
		WindowLeft,
		WindowTop,
	})
}
