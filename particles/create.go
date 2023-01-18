package particles

import ("project-particles/config"
)

func (s *System) Create() { // cette fonction permet de générer les particules et avec toute ses caractéristiques (position, vitesse, couleur, taille, ...)
	var posx, posy float64 = float64(config.General.SpawnX), float64(config.General.SpawnY)
	var spdx float64 = getFloatInBounds(float64(config.General.MinSpeedX), float64(config.General.MaxSpeedX))
	var spdy float64 = getFloatInBounds(float64(config.General.MinSpeedY), float64(config.General.MaxSpeedY))
	var Lifetime int = config.General.LifeSpanDuration
	var red,blue,green float64 = config.General.Red, config.General.Blue, config.General.Green
	var scalX,scalY float64 = config.General.ScaleX,config.General.ScaleY
	var opa float64 = config.General.Opacity
	
	l := s.Content

		if config.General.RandomOpacity{
			opa = getFloatInBounds(0,opa)
		}
		if config.General.RandomSpawnX {
			posx = getFloatInBounds(0, float64(config.General.WindowSizeX))
		}
		if config.General.RandomSpawnY {
			posy = getFloatInBounds(0, float64(config.General.WindowSizeY))
		}
		if config.General.Multicolore {
			red = getFloatInBounds(0,config.General.Red)
			blue = getFloatInBounds(0,config.General.Blue)
			green = getFloatInBounds(0,config.General.Green)
		}

		if config.General.Speedfix{
			spdx,spdy = config.General.SpeedUnique,config.General.SpeedUnique
			}
		
		l.PushFront(&Particle{
			PositionX: posx,
			PositionY: posy,
			ScaleX:    scalX, ScaleY: scalY,
			ColorRed: red, ColorGreen: green, ColorBlue: blue,
			Opacity: opa,
			SpeedX:  spdx,
			SpeedY:  spdy,
			LifeSpan: Lifetime,
			Dead: false,

		})
	}


func (p *Particle) Respawn(){
	var posx, posy float64 = float64(config.General.SpawnX), float64(config.General.SpawnY)
	if config.General.RandomSpawnX {
		posx = getFloatInBounds(0, float64(config.General.WindowSizeX))
	}
	if config.General.RandomSpawnY {
		posy = getFloatInBounds(0, float64(config.General.WindowSizeY))
	}
	p.PositionX = posx
	p.PositionY = posy
	p.LifeSpan = config.General.LifeSpanDuration
	p.Dead=false
	p.SpeedX=getFloatInBounds(-5, 5)
	p.SpeedY=getFloatInBounds(-5, 5)

}
