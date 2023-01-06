package particles

import (
	"testing"
)

// Vérification que la fonction update met bien à jour la position des particules
// sur lesquelles on l'utilise
func TestParticleUpdate(t *testing.T) {
	s := NewSystem()

	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		posx := p.PositionX
		posy := p.PositionY
		s.Update()
		if p.PositionX == posx || p.PositionY == posy {
			t.Fail()
		}
	}
}
