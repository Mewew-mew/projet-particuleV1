package particles

import (
	"project-particles/config"
	"testing"
)


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

		

