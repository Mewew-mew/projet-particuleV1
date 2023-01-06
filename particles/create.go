package particles

import "project-particles/config"

func (s *System) Create(count int) {// cette fonction permet de générer les particules et avec toute ses caractéristiques (position, vitesse, couleur, taille, ...) 
	var posx, posy float64

	if !config.General.RandomSpawn { 
		posx = float64(config.General.SpawnX)
		posy = float64(config.General.SpawnY)
	}

//test
	l := s.Content
	for i := 0; i < count; i++ { 
		if config.General.RandomSpawn {  
			posx = getFloatInBounds(0, float64(config.General.WindowSizeX))
			posy = getFloatInBounds(0, float64(config.General.WindowSizeY))
		}
		l.PushFront(&Particle{
			PositionX: posx,
			PositionY: posy,
			ScaleX:    2, ScaleY: 10,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
			SpeedX:  getFloatInBounds(-2, 2),
			SpeedY:  getFloatInBounds(-2, 2),
			BorderSpan: false,
		})
	}
}
