package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"project-particles/config"
)

// ParticleImage est une variable globale pour stocker l'image d'une particule
var ParticleImage, ExtensionsButtonBackground, ExtensionsButton, ExtensionsBackground,
ResetButton, CercleExtension, RainExtension, RainbowExtension, GlueExtension, DrawExtension,
PatternExtension, GravityExtension, CustomExtension *ebiten.Image


// Get charge en mémoire l'image de la particule. (particle.png)
// Vous pouvez changer cette image si vous le souhaitez, et même en proposer
// plusieurs, qu'on peut choisir via le fichier de configuration. Cependant
// ceci n'est pas du tout central dans le projet et ne devrait être fait que
// si vous avez déjà bien avancé sur tout le reste.
func Get() {

	if config.General.Rainbow  {
		config.General.ParticleImage = "assets/particle_luciole.png"
	}

	if config.General.Circle {
		config.General.ParticleImage = "assets/particle_cercle.png"
	}

	var err error
	ParticleImage, _, err = ebitenutil.NewImageFromFile(config.General.ParticleImage)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}

	ExtensionsButtonBackground, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Button_Background.png")
	if err != nil {
		log.Fatal("Problem while loading extension button background image: ", err)
	}

	ExtensionsButton, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Button.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	ExtensionsBackground, _, err = ebitenutil.NewImageFromFile("assets/Extensions_Background.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	ResetButton, _, err = ebitenutil.NewImageFromFile("assets/Reset_Button.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	CustomExtension, _, err = ebitenutil.NewImageFromFile("assets/Custom_Extension.png")
	if err != nil {
		log.Fatal("Problem while loading extension button image: ", err)
	}

	CercleExtension, _, err = ebitenutil.NewImageFromFile("assets/cercle.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}
	GlueExtension, _, err = ebitenutil.NewImageFromFile("assets/glue.png")
	if err != nil {
		log.Fatal("Problem while loading extension glue image: ", err)
	}
	GravityExtension, _, err = ebitenutil.NewImageFromFile("assets/graviter.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}
	DrawExtension, _, err = ebitenutil.NewImageFromFile("assets/draw.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}
	PatternExtension, _, err = ebitenutil.NewImageFromFile("assets/pattern.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}
	RainbowExtension, _, err = ebitenutil.NewImageFromFile("assets/rainbow.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}
	RainExtension, _, err = ebitenutil.NewImageFromFile("assets/pluie.png")
	if err != nil {
		log.Fatal("Problem while loading extension cercle image: ", err)
	}



}
