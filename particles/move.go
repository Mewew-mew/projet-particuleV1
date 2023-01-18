package particles

//fonction qui fais bouger la particule
func (p *Particle) Move() {
		p.PositionX += p.SpeedX
		p.PositionY += p.SpeedY
}