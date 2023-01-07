package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// cette fonction créé aussi une particule en fonction du nombre de InitNumParticles
//
func NewSystem() System {
	l := list.New()
	s:= System{Content: l}
	s.Create(config.General.InitNumParticles)
	return s
}
