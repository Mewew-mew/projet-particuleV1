package particles

import (
	"testing"
)



// Vérification que tout se passe bien dans quelques cas particuliers
func TestGetFloatInBoundsRangeSpecial(t *testing.T) {
	if getFloatInBounds(0, 0) != 0 {
		t.Fail()
	}
}
