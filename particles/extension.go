package particles

import ("project-particles/config"

)

// gravity ajoute une valeur gravité à la vitesse Y de la particule pour donner un effet de gravité
func (p *Particle) gravity(gravity float64) {
	p.SpeedY += gravity
}

// decreaseOpacity diminue l'opacité d'une particule
func (p *Particle) decreaseOpacity(opacity float64) {
	p.Opacity -= opacity
}

// decreaseScale diminue la taille d'un particule
func (p *Particle) decreaseScale(scale float64) {
	p.ScaleX -= scale
	p.ScaleY -= scale
}

// decreaseRed diminue la couleur rouge d'une particule
func (p *Particle) decreaseRed(red float64) {
	p.ColorRed -= red
}

// decreaseGreen diminue la couleur verte d'une particule
func (p *Particle) decreaseGreen(green float64) {
	p.ColorGreen -= green
}

// decreaseBlue diminue la couleur blueue d'une particule
func (p *Particle) decreaseBlue(blue float64) {
	p.ColorBlue -= blue
}

// rotateParticle modifie la rotation d'une particule
func (p *Particle) rotateParticle(rotation float64) {
	p.Rotation += rotation
}


//Creation d'un system d'update uniquement les particules 
func (p *Particle) newParticle() {

	//Si la methode Graviter vaut true alors on lance la fonction de Gravity (voire debut de la page ligne 8)
	if config.General.Gravite {
		p.gravity(config.General.ConstanteGravite)
	}

	//Condition qui tue la particule lorsque elle sort de l'ecran (avec une petite marge de 100px)
		if p.PositionX < 0 - config.General.Marges ||
		p.PositionY < 0 - config.General.Marges ||
		p.PositionX > float64(config.General.WindowSizeX) +config.General.Marges ||
		p.PositionY > float64(config.General.WindowSizeY) + config.General.Marges {
			p.Dead = true
		}
	

	p.LifeSpan -= 1
	
	//Condition qui tue la particule lorsque le lifespan atteind 0
	if config.General.LifeSpan && p.LifeSpan <= 0 {
		p.Dead = true
	}

	//Condition qui fais respwan la particule lorsque le lifespan atteind -2 et qui permet de réutiliser une particule morte 
	if config.General.LifeSpan && p.LifeSpan <=-2{
		p.Respawn()
	}

	//Condition qui active le Rainbow (c'est a dire rotation de la particule, changement de couleur)
	if config.General.Rainbow {
		p.rotateParticle(0.08)
		p.PositionX += 1 	
	}


	//Condition qui fais rebondire les particule si elle entre en contact avec un bords 
	if config.General.ColisionsBords {
		if p.PositionX <= 0 || p.PositionX >= float64(config.General.WindowSizeX) - p.ScaleX {
			p.SpeedX = -p.SpeedX
	}
		if p.PositionY <= 0 || p.PositionY >= float64(config.General.WindowSizeY) - p.ScaleY {
			p.SpeedY = -p.SpeedY
		}
	}

	//Si le cercle est activé, les particule vont petit à petit perdre de la couleur pour devenir completement invisible
	if config.General.Circle {
		p.decreaseRed(0.006)
		p.decreaseGreen(0.001)
		p.decreaseBlue(0.002)
	}

	//Si le Snow est activé, les particule vont petit à petit perdre la couleur rouge et vont tournée sur elle-meme 
	if config.General.Snow{
		p.decreaseRed(0.006)
		p.rotateParticle(0.08)	}

}
