package particles

import (
	"project-particles/config"
	"testing"
)

// Vérification que la fonction SpawnUp met bien à jour le nombre des particules
// qui'il y a dans le systeme en le mettant à jours grace a la fonction update
func TestSpawnRate(t *testing.T) {
	config.Get("../config.json")
	config.General.SpawnRate = 1.0

	s := NewSystem()
	var init = config.General.InitNumParticles
	s.Update()

	if s.Content.Len() == init {
		t.Fail()
	}

}

// fonction qui test la mise à jour d'un système pour SpawnRate négatif ou nul
func TestNoSpawnRate(t *testing.T) {
	var s System

	config.General.InitNumParticles = 5
	config.General.SpawnRate = -1
	s = NewSystem()
	s.Update()

	if s.Content.Len() != 5 {
		t.Fail()
	}

}


// fonction qui test la mise à jour d'un système pour SpawnRate inferieur à 1 et positif
func TestNoSpawnRate(t *testing.T) {
	var s System

	config.General.InitNumParticles = 5
	config.General.SpawnRate = 0.5
	s = NewSystem()
	s.Update()

	if s.Content.Len() != 5 {
		t.Fail()
	}

}

// fonction qui test la mise à jour d'un système pour SpawnRate superieur à 1 et positif
func TestNoSpawnRate(t *testing.T) {
	var s System

	config.General.InitNumParticles = 5
	config.General.SpawnRate = 2
	s = NewSystem()
	s.Update()

	if s.Content.Len() != 7 {
		t.Fail()
	}

}
