package particles

import "container/list"

// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content *list.List // contenue du System
	tick int		   // nombre de fois que la fonction update
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
type Particle struct {
	PositionX, PositionY            float64 // Position à l'écran
	Rotation                        float64 // Orientation
	ScaleX, ScaleY                  float64 // Taille
	ColorRed, ColorGreen, ColorBlue float64 // Couleur
	Opacity                         float64 // Transparence
	SpeedX, SpeedY            		float64 // Vitesse
}

