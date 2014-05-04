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

	AbilityBoundingLeft  = 602
	AbilityBoundingTop   = 185

	AbilityWidth   = 75
	AbilityHeight  = 74
	AbilityGutter  = 86
	AbilityColumns = 6

	SlotWidth = iota
	SlotHeight
	SlotGutter
	SlotPadding
	SlotLeft
	SlotTop
	SlotBounding

	MouseSlotTopOffset

	PassiveSlotWidth
	PassiveSlotHeight
	PassiveSlotTopOffset

	ButtonWidth
	ButtonHeight
	ButtonGutter

	RuneWidth
	RuneHeight
	RuneLeftOffset
	RuneTopOffset
	RuneGutter
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
