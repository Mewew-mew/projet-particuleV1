package particles

import (
	"project-particles/config"
)

func (s *System) SpawnUp() { // cette fonction met à jours le systeme qui indique combien de nouvelles particules qui seront engendrées
	// à chaque appel à la fonction Update().
	s.tick++

	if config.General.SpawnRate <= 0 { // si la valeur du spawnRate est à 0 on initialise 0 particule
		s.Create(0)

	} else if config.General.SpawnRate < 1 || config.General.SpawnRate > 0 { // si la valeur du SpawnRate est inferieur à 1 alors on engendra une particule sur deux en fonction de la variable tick
		if s.tick%int(1/config.General.SpawnRate) == 0 {
			s.Create(1)
		}
	} else {
		for value := 0; value < (int(config.General.SpawnRate)); value++ { // sinon on engendre des particule en fonction du spawnRate 
			s.Create(1)
		}
	}

}
