package res

import (
	"testing"
)

func TestAbility(t *testing.T) {
	total := 4

	r := Resolution{1920, 1200}
	for i := 0; i < total; i++ {
		t.Error(Ability{total, i}.Point(r))
	}
}
