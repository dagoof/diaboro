package res

import (
	"testing"
)

var r = Resolution{1920, 1200}

func TestAbility(t *testing.T) {
	total := 5
	for i := 0; i < total; i++ {
		t.Error(total, Ability{total, i}.Point(r))
	}

	total = 4
	for i := 0; i < total; i++ {
		t.Error(total, Ability{total, i}.Point(r))
	}

}

func TestPassive(t *testing.T) {
	total := 18
	for i := 0; i < total; i++ {
		t.Error(total, Passive{total, i}.Point(r))
	}
}

func TestRune(t *testing.T) {
	for i := 0; i < RuneColumns; i++ {
		t.Error(Rune{i}.Point(r))
	}
}

func TestButton(t *testing.T) {
	t.Error("accept", Accept.Point(r))
	t.Error("cancel", Cancel.Point(r))
}

func TestDirection(t *testing.T) {
	t.Error("left", Left.Point(r))
	t.Error("right", Right.Point(r))
}
