package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// dans cette fonction on créé un nouveau systeme en fonction des particules

func NewSystem() System {
	l := list.New()
	s:= System{Content: l}
	s.ActiveExtension = "Si vous voulez arrétez le systeme, appuyer sur ESPACE"

	for i:=0 ; i < config.General.InitNumParticles ; i++{
		s.Create()
	}
	return s
}
