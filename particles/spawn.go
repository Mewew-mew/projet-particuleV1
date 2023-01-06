package particles

import (
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
//
func (s *System) SpawnUp() { // cette fonction met à jours le systeme qui indique combien de nouvelles particules qui seront engendrées
	// à chaque appel à la fonction Update().
	s.tick++

	if config.General.SpawnRate <= 0 {
		s.Create(0)

	} else if config.General.SpawnRate < 1 || config.General.SpawnRate > 0 { // si la valeur du tick qui sera engendree est inferieur à 1 alors on engendra une particule sur deux tick
		if s.tick%int(1/config.General.SpawnRate) == 0 {
			s.Create(1)
		}
	} else {
		for value := 0; value < (int(config.General.SpawnRate)); value++ { // si la valeur du tick qui sera engendree est superieur à 1 alors on engendra une particule par rapport au nombre de tick et de la value
			s.Create(1)
		}
	}

}
