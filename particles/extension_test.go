package particles

import (
	"project-particles/config"
	"testing"
)


// Vérification que multicolore crée des couleurs RGB comprises entre 0 et 1
func TestMulticolore(t *testing.T) {
	config.General.Multicolore = true
	s := NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)

		if p.ColorRed < 0 || p.ColorRed > 1 ||
		p.ColorGreen < 0 || p.ColorGreen > 1 ||
		p.ColorBlue < 0 || p.ColorBlue > 1 {
			t.Fail()
		}
	
	}
}

// Vérification que les marges tue les particules
func TestParticleMarges(t *testing.T) {
	config.General.Gravite = false
	config.General.RenderMarges = true
	config.General.Marges = 0
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.SpawnX = 400
	config.General.SpawnY = 300
	s := NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		for p.PositionX > 0 && p.PositionX < float64(config.General.WindowSizeX) && 
		p.PositionY > 0 && p.PositionY < float64(config.General.WindowSizeY) {
			p.newParticle()
		}

		if !p.Dead {
			t.Fail()
		}
	}
}
