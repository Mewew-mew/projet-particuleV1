package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	Maxpartciles			 int
	RandomSpawnX             bool
	RandomSpawnY             bool

	ScaleX					 float64
	ScaleY					 float64

	CercleHauteur            float64
	CercleLargeur			 float64

	SpawnOnMouse			 bool
	LockOnMouse				 bool

	SpawnX, SpawnY           int
	SpawnRate                float64

	Speedfix				 bool
	SpeedUnique				 float64
	MinSpeedX				 float64
	MaxSpeedX				 float64
	MinSpeedY				 float64
	MaxSpeedY				 float64

	Gravite 				 bool
	ConstanteGravite		 float64

	RenderMarges			 bool
	Marges					 float64

	LifeSpan				 bool
	LifeSpanDuration		 int


	RandomOpacity			 bool
	Opacity					 float64
	Red						 float64
	Green					 float64
	Blue					 float64

	Multicolore 			 bool

	Rainbow					 bool
	ColisionsBords			 bool
	Circle					 bool

	CollisionsParticules     bool

	Pluie					 bool

	Draw					 bool			


}

var General Config
