package main

import (
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

// Vérification que resetSystem réinitialise bien le système
func TestResetSystem (t *testing.T) {
	config.General.InitNumParticles = 100
	config.General.ParticleImage = "assets/particle.png"
	g := game{system: particles.NewSystem()}
	config.General.InitNumParticles = 0
	g.resetSystem()
	if g.system.Content.Len() != 0 {
		t.Fail()
	}
}


// Vérification que extensionChanged change bien le nom de l'extension
func TestExtensionName (t *testing.T) {
	g := game{system: particles.NewSystem()}
	g.extensionChanged("Test")
	if g.system.ActiveExtension[0] != 84 || g.system.ActiveExtension[1] != 101 ||
	g.system.ActiveExtension[2] != 115 || g.system.ActiveExtension[3] != 116 {
		t.Fail()
	}
}

// Vérification que setConfig modifie bien la configuration
func TestSetConfig (t *testing.T) {
	setConfig("assets/particle.png", 1000,100000, true, true,
	3, 3, 200, 200, true, true, 500, 500, 2, true, 20, 0, 
	0, 0, 0, true, 1, true, 100, true, 1000, true, 1.0,
	1.0, 1.0, 1.0, true, true, true, true, true, true)
	
	if config.General.InitNumParticles != 1000 || config.General.Maxpartciles != 100000 || !config.General.RandomSpawnX || !config.General.RandomSpawnY || 
	config.General.ScaleX != 3 || config.General.ScaleY != 3 || config.General.CercleHauteur !=200 || config.General.CercleLargeur !=200 || 
	!config.General.SpawnOnMouse || !config.General.LockOnMouse ||
	config.General.SpawnX != 500 || config.General.SpawnY != 500 || config.General.SpawnRate != 2 || !config.General.Speedfix || config.General.SpeedUnique !=20 ||
	config.General.MinSpeedX !=0 || config.General.MaxSpeedX !=0 || config.General.MinSpeedY !=0 ||
	config.General.MaxSpeedY !=0 || !config.General.Gravite || config.General.ConstanteGravite !=1 ||
	!config.General.RenderMarges || config.General.Marges !=100 || !config.General.LifeSpan||
	config.General.LifeSpanDuration !=1000 || !config.General.RandomOpacity ||
	config.General.Opacity !=1.0 || config.General.Red !=1.0 || config.General.Green !=1.0||
	config.General.Blue!=1.0 || !config.General.Multicolore || !config.General.Rainbow ||
	!config.General.ColisionsBords || !config.General.Circle || !config.General.CollisionsParticules || !config.General.Snow {
		t.Fail()
	}
}



