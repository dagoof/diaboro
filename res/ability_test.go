package res

import (
	"testing"
)

func TestAbility(t *testing.T) {
	r := Resolution{1920, 1200}

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
	r := Resolution{1920, 1200}

	total := 18
	for i := 0; i < total; i++ {
		t.Error(total, Passive{total, i}.Point(r))
	}
}

func TestRune(t *testing.T) {
	r := Resolution{1920, 1200}

	for i := 0; i < RuneColumns; i++ {
		t.Error(Rune{i}.Point(r))
	}
}
