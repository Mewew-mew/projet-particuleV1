package particles



///"math/rand"
///"project-particles/config"
///"time"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	
	s.SpawnUp() // on rajoute des particule en fonction du spawn rate
	
	for e := s.Content.Front(); e != nil; e= e.Next(){
		p := e.Value.(*Particle)
		p.PositionX+=p.SpeedX
		p.PositionY+=p.SpeedY
	}

}
	


