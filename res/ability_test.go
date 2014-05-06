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

func TestMouseSlot(t *testing.T) {
	total := 2
	for i := 0; i < total; i++ {
		t.Error(total, MouseSlot{i}.Point(r))
	}
}

func TestKeySlot(t *testing.T) {
	total := 4
	for i := 0; i < total; i++ {
		t.Error(total, KeySlot{i}.Point(r))
	}
}

func TestPassiveSlot(t *testing.T) {
	total := 4
	for i := 0; i < total; i++ {
		t.Error(total, PassiveSlot{i}.Point(r))
	}
}
