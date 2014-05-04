package res

type Offset float64

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

	AbilityBoundingLeft  = 0
	AbilityBoundingTop  = 0
	AbilityBoundingWidth = 1.0

	AbilityWidth  = 0.2
	AbilityHeight = 0.2
	AbilityGutter = 0.1

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
