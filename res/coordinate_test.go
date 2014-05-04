package res

import (
	"testing"
)

func TestCoordinate(t *testing.T) {
	box := Coordinate{2, 2}
	point := Coordinate{4, 3}

	relative := point.Unshift(box)

	if (relative != Coordinate{2, 1}) {
		t.Error(relative)
	}

	original := relative.Shift(box)

	if original != point {
		t.Error("shifted unshift not same")
	}
}
