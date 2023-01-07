package particles

import "math/rand"

// getFloatInBounds est une fonction qui retourne un nombre décimal aleatoire compris
// entre un min et un max
func getFloatInBounds(min, max float64) float64 { 
	return (max-min)*rand.Float64() + min
}
