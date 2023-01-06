package particles

import (
	"project-particles/config"
	"testing"
)

// Vérification que NewSystem crée bien un système avec le nombre de particules
// initiales indiqué dans config.General.InitNumParticles
func TestNew(t *testing.T) {
	s := NewSystem()
	var init = config.General.InitNumParticles 

	if s.Content.Len() != init {
		t.Fail()
	}
}


// Vérification que NewSystem crée bien un système vide
func TestVide(t *testing.T){
	config.General.InitNumParticles = 0
	s := NewSystem()

	if s.Content.Len()  != 0 {
		t.Fail()
	}
}
