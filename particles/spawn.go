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
	if s.Content.Len() <= config.General.Maxpartciles{
		// à chaque appel à la fonction Update().
		s.NumToSpawn += config.General.SpawnRate
		for s.NumToSpawn >= 1 {
			s.Create()
			s.NumToSpawn--
		}
	}
}
