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

// Vérifictation que graity ajoute une valeur gravité à la vitesse Y d'une particule
func TestGravity(t *testing.T) {
	s := NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)

		p.SpeedY = getFloatInBounds(0,1) * 100
		var SpeedY = p.SpeedY
		p.gravity(-getFloatInBounds(0,1))
		

		if p.SpeedY > SpeedY {
			t.Fail()
		}
	}
}

// Vérifictation que rotateParticle modifie la rotation d'une particule
func TestRotateParticle(t *testing.T) {
	s := NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)

		p.Rotation = getFloatInBounds(0,1)
		var Rotation = p.Rotation
		p.rotateParticle(getFloatInBounds(0,1))

		if p.Rotation < Rotation {
			t.Fail()
		}
	}
}


// Vérifictation que decreaseOpacity diminue l'opacité d'une particule
func TestDecreaseOpactiy(t *testing.T) {
	s := NewSystem()
		for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)


		p.Opacity = getFloatInBounds(0,1)
		var Opacity = p.Opacity
		p.decreaseOpacity(getFloatInBounds(0,1))

		if p.Opacity > Opacity {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseRed diminue la couleur rouge d'une particule
func TestDecreaseRed(t *testing.T) {
	s := NewSystem()
		for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)


		p.ColorRed = getFloatInBounds(0,1)
		var ColorRed = p.ColorRed
		p.decreaseRed(getFloatInBounds(0,1))

		if p.ColorRed > ColorRed {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseGreen diminue la couleur verte d'une particule
func TestDecreaseGreen(t *testing.T) {
	s := NewSystem()
		for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)


		p.ColorGreen = getFloatInBounds(0,1)
		var ColorGreen = p.ColorGreen
		p.decreaseGreen(getFloatInBounds(0,1))

		if p.ColorGreen > ColorGreen {
			t.Fail()
		}
	}
}

// Vérifictation que decreaseBlue diminue la couleur bleue d'une particule
func TestDecreaseBlue(t *testing.T) {
	s := NewSystem()
		for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)


		p.ColorBlue = getFloatInBounds(0,1)
		var ColorBlue = p.ColorBlue
		p.decreaseBlue(getFloatInBounds(0,1))

		if p.ColorBlue > ColorBlue {
			t.Fail()
		}
	}
}