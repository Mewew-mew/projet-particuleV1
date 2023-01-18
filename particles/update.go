package particles

import(
"project-particles/config"
"github.com/hajimehoshi/ebiten/v2"
"math"
"math/rand"

)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	// Changer la couleur RGB de config (pour changer la couleur de base des particules lors du spawn) ainsi cela donne l'effet d'un arc en ciel 
	if config.General.Rainbow{
	if config.General.Red >= 1 && config.General.Green <= 1 && config.General.Blue <= 0 { // Rouge
		config.General.Green += 0.01
	} else if config.General.Red >= 0 && config.General.Green >= 1 && config.General.Blue <= 0 { // Jaune
		config.General.Red -= 0.01
	} else if config.General.Red <= 0 && config.General.Green >= 1 && config.General.Blue <= 1 { // Vert
		config.General.Blue += 0.01
	} else if config.General.Red <= 0 && config.General.Green >= 0 && config.General.Blue >= 1 { // Bleu clair
		config.General.Green -= 0.01
	} else if config.General.Red <= 1 && config.General.Green <= 0 && config.General.Blue >= 1 { // Bleu foncé
		config.General.Red += 0.01
	} else if config.General.Red >= 1 && config.General.Green <= 0 && config.General.Blue >= 0 { // Violet
		config.General.Blue -= 0.01
	}
}

	s.SpawnUp() // on rajoute des particule en fonction du spawn rate


		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			p.Move()
			p.newParticle()


		// on sauvegarde la position de la souris lors de la dernière frame 
		// et on actualise s.MouseX et s.MouseY avec la nouvelle position de la souris
		var LastX, LastY int = s.MouseX, s.MouseY
		s.MouseX, s.MouseY = ebiten.CursorPosition()

		// Si SpawnOnMouse est activé on change SpawnX et SpawnY dans la config à la position 
		// de la souris pendant cette frame. On retire la taille de l'image / 2 pour faire apparaître le centre
		// de la particule sous la souris et non le coin haut gauche
		if config.General.SpawnOnMouse {
			config.General.SpawnX, config.General.SpawnY = s.MouseX-int(p.ScaleX/2), s.MouseY-int(p.ScaleX/2)
		}

		// Si LockOnMouse est activé, alors on déplace le générateur en fonction de la souris
		// pour compenser la vitesse lorsque la souris se déplace, on ajoute à la position des particules
		// la différence entre la nouvelle et l'ancienne position
		if config.General.LockOnMouse {
			p.PositionX += float64(s.MouseX - LastX)
			p.PositionY += float64(s.MouseY - LastY) 
		}

		if config.General.Draw {
			config.General.SpawnRate = 0
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
				s.Create()
			}
		}

		//Si Circle est activé, on créé un cercle virtuel ensuite chaque particule spawn sur ce cercle 
		// de plus le centre du systeme deviens la souris , ainsi le generateur en forme de cercle suit le mouvement de la souris 
		if config.General.Circle{
			var cercle float64 = rand.Float64()*2*math.Pi
			centreX := float64(s.MouseX)-p.ScaleX/2
			centreY :=  float64(s.MouseY)-p.ScaleX/2
			p.PositionX = math.Cos(cercle)*config.General.CercleHauteur + centreX
			p.PositionY = math.Sin(cercle)*config.General.CercleLargeur + centreY
	
		}

		//Si Collisions Particule est activé, alors on créé une deuxieme boucle pour pas avoir deux fois la meme particule 
		//si la premiere et la deuxieme particule se rencontre on annule leurs vitesse pour donné l'illusion quelle se colle 
		if config.General.CollisionsParticules {
			for e := s.Content.Front(); e != nil; e = e.Next() {
				p1 := e.Value.(*Particle)
				if p1 != p{	
					if math.Sqrt((p.PositionX - p1.PositionX) * (p.PositionX - p1.PositionX) + (p.PositionY - p1.PositionY) * (p.PositionY - p1.PositionY)) < 45{//pour la taille 1
						p.SpeedX,p1.SpeedX = 0,0
						p.SpeedY,p1.SpeedY = 0,0
					}


				}
			}
		}
	}
}
	







