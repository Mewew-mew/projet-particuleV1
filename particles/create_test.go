package particles

import (
	"project-particles/config"
	"testing"
)

/*
// Vérification que les particules générées quand config.General.RandomSpawn
// vaut true sont bien toutes à l'intérieur de l'écran
func TestTrue(t *testing.T) {
	config.General.RandomSpawnX = true
	config.General.RandomSpawnY =true
	config.General.InitNumParticles = 100

	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	s := NewSystem()

	for e := s.Content.Front(); e != nil; e= e.Next(){
		p := e.Value.(*Particle)
		if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX) ||
			p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY) {
			t.Fail()
		}        
	}
}
*/
// Vérification que les particules générées quand config.General.RandomSpawn
// vaut false sont bien toutes à la position demandée, c'est-à-dire aux
// coordonnées (config.General.SpawnX, config.General.SpawnY)
func TestFalse(t *testing.T) {
	config.General.RandomSpawnX = false
	config.General.RandomSpawnY = false


	s := NewSystem()

	for e := s.Content.Front(); e != nil; e= e.Next(){
		p := e.Value.(*Particle)
		if int(p.PositionX) != config.General.SpawnX || int(p.PositionY) != config.General.SpawnY {
			t.Fail()
		}        
	}
}

		

