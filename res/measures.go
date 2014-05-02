package res

const (
	SlotWidth
	SlotHeight
	SlotLeftOffset
	SlotTopOffset
	SlotGutter
	SlotVerticalGutter

	MouseSlotTopOffset

	PassiveSlotWidth
	PassiveSlotHeight
	PassiveSlotTopOffset

	ButtonWidth
	ButtonHeight
	ButtonGutter

	AbilityWidth
	AbilityHeight
	AbilityGutter
	AbilityTopOffset

	DirectionWidth
	DirectionHeight
	DirectionTopOffset
	LeftDirectionLeftOffset
	RightDirectionLeftOffset

	RuneWidth
	RuneHeight
	RuneLeftOffset
	RuneTopOffset
	RuneGutter
)

func Centered(left, top, width, height int) Coordinate {
	return Coordinate{
		left + (width / 2),
		top + (height / 2),
	}
}
