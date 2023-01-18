package particles

import (
	"testing"
	"project-particles/config"
	
)

// Vérification que les particules ne sortent pas de l'écran avec CollisionsBords
func TestCollisionsBords(t *testing.T) {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.LifeSpan = false
	config.General.RenderMarges = true
	config.General.Marges = 50
	config.General.ColisionsBords = true
	config.General.MinSpeedX = -10
	config.General.MaxSpeedX = 10
	config.General.MinSpeedY = -10
	config.General.MaxSpeedY = 10
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 100
	s := NewSystem()
	s.ImageWidth, s.ImageHeight = 10, 10
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
			if p.Dead {
				t.Fail()
		}
	}
}


//Test verifiant que les mouvement de la particule sont bien pris en compte
func TestMove(t *testing.T){
	s:=NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		var X,Y float64
		p := e.Value.(*Particle)
		X,Y = p.PositionX,p.PositionY

		p.Move()

		if p.PositionX != X || p.PositionY != Y {
			t.Fail()
		}
	}
}

	

