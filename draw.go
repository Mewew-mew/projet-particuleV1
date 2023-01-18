package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"strconv"
	"image/color"
	"golang.org/x/image/font/basicfont"	
)

func (g *game) resetSystem() {
	assets.Get()
	g.system = particles.NewSystem()
}

// extensionChanged est une méthode qui change le nom de l'extension active
func (g *game) extensionChanged(extension string) {
	g.system.ActiveExtension = extension
}

// setConfig est une fonction qui modifie la config avec les arguments données
func setConfig(ParticleImage string, InitNumParticles int, Maxpartciles int, RandomSpawnX bool,
	RandomSpawnY bool, ScaleX float64 , ScaleY float64, 
	CercleHauteur float64, CercleLargeur float64, SpawnOnMouse bool,
	LockOnMouse bool, SpawnX int, SpawnY int, SpawnRate float64,
	Speedfix bool, SpeedUnique float64, MinSpeedX float64, MaxSpeedX float64, 
	MinSpeedY float64,MaxSpeedY float64,  Gravite bool, ConstanteGravite float64, 
	RenderMarges bool, Marges float64, LifeSpan bool, LifeSpanDuration int,
	RandomOpacity bool, Opacity float64,
	Red float64, Green float64, Blue float64, Multicolore bool, Rainbow bool,
	ColisionsBords bool,Circle bool,
	CollisionsParticules bool, Snow bool) {
		config.General.ParticleImage = ParticleImage
		config.General.InitNumParticles = InitNumParticles
		config.General.Maxpartciles = Maxpartciles
		config.General.RandomSpawnX = RandomSpawnX
		config.General.RandomSpawnY = RandomSpawnY
		config.General.ScaleX = ScaleX
		config.General.ScaleY = ScaleY
		config.General.CercleHauteur = CercleHauteur
		config.General.CercleLargeur = CercleLargeur
		config.General.SpawnOnMouse = SpawnOnMouse
		config.General.LockOnMouse = LockOnMouse
		config.General.SpawnX = SpawnX
		config.General.SpawnY = SpawnY
		config.General.SpawnRate = SpawnRate
		config.General.Speedfix = Speedfix
		config.General.SpeedUnique = SpeedUnique
		config.General.MinSpeedX = MinSpeedX
		config.General.MaxSpeedX = MaxSpeedX
		config.General.MinSpeedY = MinSpeedY
		config.General.MaxSpeedY = MaxSpeedY
		config.General.Gravite = Gravite
		config.General.ConstanteGravite = ConstanteGravite
		config.General.RenderMarges = RenderMarges
		config.General.Marges = Marges
		config.General.LifeSpan = LifeSpan
		config.General.LifeSpanDuration = LifeSpanDuration
		config.General.RandomOpacity = RandomOpacity
		config.General.Opacity = Opacity
		config.General.Red = Red
		config.General.Green = Green
		config.General.Blue = Blue
		config.General.Multicolore = Multicolore
		config.General.Rainbow = Rainbow
		config.General.ColisionsBords = ColisionsBords
		config.General.Circle = Circle
		config.General.CollisionsParticules = CollisionsParticules
		config.General.Snow = Snow


	}
	


// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			if !p.Dead{
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}


		
	// Afficher l'rrière plan du bouton avec la flèche à droite de l'écran
	ExtensionsButtonBackgroundOptions := ebiten.DrawImageOptions{}
	ExtensionsButtonBackgroundOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 50, float64(config.General.WindowSizeY) / 2 - 125)
	screen.DrawImage(assets.ExtensionsButtonBackground, &ExtensionsButtonBackgroundOptions)

	// Afficher la flèche à droite de l'écran
	ExtensionsButtonOptions := ebiten.DrawImageOptions{}
	// Retourner la flèche lors d'un appui sur le bouton
	if g.system.ExtensionsButtonPressed {
		ExtensionsButtonOptions.GeoM.Rotate(3.14)
		ExtensionsButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 8, float64(config.General.WindowSizeY) / 2 + 33.5)
	} else {
		ExtensionsButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) - 42.5, float64(config.General.WindowSizeY) / 2 - 33.5)
	}
	screen.DrawImage(assets.ExtensionsButton, &ExtensionsButtonOptions)




	// Taille de l'arrière plan lorsque le menu de sélection des extensions est activé
	// Il a finalement été retiré et n'est donc pas visible
	var BackgroundSizeX = float64(config.General.WindowSizeX) * 0.8
	var BackgroundSizeY = BackgroundSizeX / 1.78


	// Appui sur le bouton gauche de la souris
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		CursorXPosition, CursorYPosition := ebiten.CursorPosition()

		// Si la souris est au dessus le bouton avec la flèche à droite de l'écran
		if float64(CursorXPosition) > float64(config.General.WindowSizeX) - 42.5 && CursorYPosition > config.General.WindowSizeY / 2 - 62 && CursorYPosition < config.General.WindowSizeY / 2 + 62 {
			g.system.ExtensionsButtonPressed = !g.system.ExtensionsButtonPressed
		}


		if g.system.ExtensionsButtonPressed {
			// Appui sur le bouton réinitaliser le système
			if float64(CursorXPosition) > float64(config.General.WindowSizeX) / 2 - 227.5 &&
			float64(CursorXPosition) < float64(config.General.WindowSizeX) / 2 + 227.5 &&
			float64(CursorYPosition) > float64(config.General.WindowSizeY) - 50 &&
			float64(CursorYPosition) < float64(config.General.WindowSizeY) - 10 {
				g.resetSystem()
			}
		
			//1er extension
			//  Selectonner l'image corespondante pour sélectionner l'extension Cercle
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig(
					"assets/particle.png", //"ParticleImage"
					1000, //"InitNumParticles"
					100000, //Maxpartciles
					false, //RandomSpawnX
					false, //RandomSpawnY
					3,//ScaleX
					3,//ScaleY
					200, //CercleHauteur
					200, //CercleLargeur
					true, //SpawnOnMouse
					true,//LockOnMouse
					920, //SpawnX
					400, //SpawnY
					2,  //SpawnRate
					false, //Speedfix
					12, //SpeedUnique
					-5.0, //MinSpeedX			 
					5.0, //MaxSpeedX				 
					-5.0, //MinSpeedY				 
					5.0, //MaxSpeedY		 		
					false, //Gravite
					0.2, //ConstanteGravite
					false, //RenderMarges
					100, //Marges
					false, //LifeSpan
					1000, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					false, //Multicolore
					false, //Rainbow
					false, //ColisionsBords
					true, //Circle
					false, //CollisionsParticules
					false,//Snow
				)				
				g.resetSystem()
				g.extensionChanged("Cercle")
			}

			//2eme extensions
			//  Selectonner l'image corespondante pour sélectionner l'extension pour dessiner
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig(
					"assets/particle_cercle.png", //"ParticleImage"
					0, //"InitNumParticles"
					100000, //Maxpartciles
					false, //RandomSpawnX
					false, //RandomSpawnY
					5,//ScaleX
					5,//ScaleY
					100, //CercleHauteur
					100, //CercleLargeur
					true, //SpawnOnMouse
					true,//LockOnMouse
					920, //SpawnX
					400, //SpawnY
					0.95,  //SpawnRate
					true, //Speedfix
					0, //SpeedUnique
					-5.0, //MinSpeedX			 
					5.0, //MaxSpeedX				 
					-5.0, //MinSpeedY				 
					5.0, //MaxSpeedY		 		
					false, //Gravite
					0.2, //ConstanteGravite
					false, //RenderMarges
					100, //Marges
					false, //LifeSpan
					1000, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					false, //Multicolore
					false, //Rainbow
					false, //ColisionsBords
					false, //Circle
					true, //CollisionsParticules
					false,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Draw")
			}
			// 3eme extension
			//  Selectonner l'image corespondante pour sélectionner l'extension balle collante
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig(
					"assets/particle_cercle.png", //"ParticleImage"
					1, //"InitNumParticles"
					2000, //Maxpartciles
					true, //RandomSpawnX
					true, //RandomSpawnY
					5,//ScaleX
					5,//ScaleY
					700, //CercleHauteur
					800, //CercleLargeur
					false, //SpawnOnMouse
					false,//LockOnMouse
					950, //SpawnX
					50, //SpawnY
					2,  //SpawnRate
					false, //Speedfix
					12, //SpeedUnique
					-1.0, //MinSpeedX			 
					1.0, //MaxSpeedX				 
					-1.0, //MinSpeedY				 
					1.0, //MaxSpeedY		 		
					true, //Gravite
					0.2, //ConstanteGravite
					false, //RenderMarges
					100, //Marges
					true, //LifeSpan
					1000, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					false, //Multicolore
					false, //Rainbow
					true, //ColisionsBords
					false, //Circle
					true, //CollisionsParticules
					false,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Particule collante")
			}

			// 4eme extension
			//  Selectonner l'image corespondante pour sélectionner l'extension Graviter avec rebond
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301{
				setConfig(
					"assets/particle_outline.png", //"ParticleImage"
					10, //"InitNumParticles"
					100000, //Maxpartciles
					false, //RandomSpawnX
					false, //RandomSpawnY
					2,//ScaleX
					2,//ScaleY
					100, //CercleHauteur
					100, //CercleLargeur
					true, //SpawnOnMouse
					true,//LockOnMouse
					920, //SpawnX
					400, //SpawnY
					0.1,  //SpawnRate
					false, //Speedfix
					12, //SpeedUnique
					-5.0, //MinSpeedX			 
					5.0, //MaxSpeedX				 
					-5.0, //MinSpeedY				 
					5.0, //MaxSpeedY		 		
					true, //Gravite
					0.2, //ConstanteGravite
					false, //RenderMarges
					100, //Marges
					true, //LifeSpan
					100, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					true, //Multicolore
					false, //Rainbow
					true, //ColisionsBords
					false, //Circle
					false, //CollisionsParticules
					false,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Gravite avec rebondis")
			}

			//5eme extension
			//  Selectonner l'image corespondante pour sélectionner l'extension Snow 

			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig(
					"assets/particle_flocon.png", //"ParticleImage"
					0, //"InitNumParticles"
					1000000, //Maxpartciles
					true, //RandomSpawnX
					false, //RandomSpawnY
					2,//ScaleX
					1,//ScaleY
					100, //CercleHauteur
					100, //CercleLargeur
					false, //SpawnOnMouse
					false,//LockOnMouse
					0, //SpawnX
					-50, //SpawnY
					20,  //SpawnRate
					false, //Speedfix
					12, //SpeedUnique
					0, //MinSpeedX			 
					0, //MaxSpeedX				 
					-0.1, //MinSpeedY				 
					1, //MaxSpeedY		 		
					true, //Gravite
					0.09, //ConstanteGravite
					true, //RenderMarges
					200, //Marges
					false, //LifeSpan
					1000, //LifeSpanDuration
					true,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					false, //Multicolore
					false, //Rainbow
					false, //ColisionsBords
					false, //Circle
					false, //CollisionsParticules
					true,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Snow avec des petit flocon")
			}
			// 6eme extension
			//  Selectonner l'image corespondante pour sélectionner l'extension d'un patterne du rebond en continue exemple le logo DVD
				if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 &&
				float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05 + BackgroundSizeX / 1006 * 0.42 * 569 &&
				float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 &&
				float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 + BackgroundSizeY / 566 * 0.42 * 301{
				setConfig(
					"assets/particle.png", //"ParticleImage"
					1, //"InitNumParticles"
					100000, //Maxpartciles
					false, //RandomSpawnX
					false, //RandomSpawnY
					1,//ScaleX
					1,//ScaleY
					100, //CercleHauteur
					100, //CercleLargeur
					false, //SpawnOnMouse
					false,//LockOnMouse
					920, //SpawnX
					400, //SpawnY
					0.6,  //SpawnRate
					true, //Speedfix
					12, //SpeedUnique
					-5.0, //MinSpeedX			 
					5.0, //MaxSpeedX				 
					-5.0, //MinSpeedY				 
					5.0, //MaxSpeedY		 		
					false, //Gravite
					0.2, //ConstanteGravite
					false, //RenderMarges
					100, //Marges
					false, //LifeSpan
					1000, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					1.0, //Green
					1.0, //Blue
					true, //Multicolore
					false, //Rainbow
					true, //ColisionsBords
					false, //Circle
					false, //CollisionsParticules
					false,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Patterne d'un rebond en continue")
			}


			//7me
			//  Selectonner l'image corespondante pour sélectionner l'extension d'une vague multicolore
			if float64(CursorXPosition) > (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 &&
			float64(CursorXPosition) < (float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71 + BackgroundSizeX / 1006 * 0.42 * 569 &&
			float64(CursorYPosition) > (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 &&
			float64(CursorYPosition) < (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7 + BackgroundSizeY / 566 * 0.42 * 301 {
				setConfig(
					"assets/particle_luciole.png", //"ParticleImage"
					20, //"InitNumParticles"
					1000000, //Maxpartciles
					false, //RandomSpawnX
					true, //RandomSpawnY
					2,//ScaleX
					2,//ScaleY
					100, //CercleHauteur
					100, //CercleLargeur
					false, //SpawnOnMouse
					false,//LockOnMouse
					-10, //SpawnX
					400, //SpawnY
					20,  //SpawnRate
					false, //Speedfix
					12, //SpeedUnique
					3.0, //MinSpeedX			 
					3.0, //MaxSpeedX				 
					0, //MinSpeedY				 
					0, //MaxSpeedY		 		
					false, //Gravite
					0.2, //ConstanteGravite
					true, //RenderMarges
					100, //Marges
					false, //LifeSpan
					1000, //LifeSpanDuration
					false,//RandomOpacity
					1.0,//Opacity
					1.0, //Red
					0, //Green
					0, //Blue
					false, //Multicolore
					true, //Rainbow
					false, //ColisionsBords
					false, //Circle
					false, //CollisionsParticules
					false,//Snow


				)				
				g.resetSystem()
				g.extensionChanged("Vagues multicolore")
			}
		}
	}
	// Appui sur la barre d'espace
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.system.PauseGame = !g.system.PauseGame
	}
	
	if g.system.ExtensionsButtonPressed {

/*1*/		ResetButtonOptions := ebiten.DrawImageOptions{}
		ResetButtonOptions.GeoM.Translate(float64(config.General.WindowSizeX) / 2 - 227.5, float64(config.General.WindowSizeY) - 60)
		screen.DrawImage(assets.ResetButton, &ResetButtonOptions)
		
/*2*/			CercleExtensionOptions := ebiten.DrawImageOptions{}
		CercleExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		CercleExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055)
		screen.DrawImage(assets.CercleExtension, &CercleExtensionOptions)	
	

/*4*/			DrawExtensionOptions3 := ebiten.DrawImageOptions{}
		DrawExtensionOptions3.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		DrawExtensionOptions3.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.055)
		screen.DrawImage(assets.DrawExtension, &DrawExtensionOptions3)	

/*5*/			GlueExtensionOptions4 := ebiten.DrawImageOptions{}
		GlueExtensionOptions4.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		GlueExtensionOptions4.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.GlueExtension, &GlueExtensionOptions4)	

/*6*/			RainExtensionOptions5 := ebiten.DrawImageOptions{}
		RainExtensionOptions5.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		RainExtensionOptions5.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.38, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.RainExtension, &RainExtensionOptions5)	

/*7*/			PatternExtensionOptions := ebiten.DrawImageOptions{}
		PatternExtensionOptions.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		PatternExtensionOptions.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.05, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7)
		screen.DrawImage(assets.PatternExtension, &PatternExtensionOptions)	

/*8*/			GravityExtensionOptions8 := ebiten.DrawImageOptions{}
		GravityExtensionOptions8.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		GravityExtensionOptions8.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.3775)
		screen.DrawImage(assets.GravityExtension, &GravityExtensionOptions8)	

		RainbowExtensionOptions10 := ebiten.DrawImageOptions{}
		RainbowExtensionOptions10.GeoM.Scale(BackgroundSizeX / 1006 * 0.42, BackgroundSizeX / 1006 * 0.42)
		RainbowExtensionOptions10.GeoM.Translate((float64(config.General.WindowSizeX) - BackgroundSizeX) / 2 + BackgroundSizeX * 0.71, (float64(config.General.WindowSizeY) - (BackgroundSizeY)) / 2 + BackgroundSizeY * 0.7)
		screen.DrawImage(assets.RainbowExtension, &RainbowExtensionOptions10)
		}	
		
		text.Draw(screen, strconv.Itoa(g.system.Content.Len()), basicfont.Face7x13, config.General.WindowSizeX / 2 - 15, 13, color.RGBA{0, 255, 0, 255})
		
	
		// Affichage de l'extension en bas à gauche de l'écran
		text.Draw(screen, "Extension : " + g.system.ActiveExtension, basicfont.Face7x13, 20, config.General.WindowSizeY - 13, color.RGBA{255, 255, 255, 255})
	}

