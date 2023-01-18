#SAE Implementation / Particule / Groupe 2-2

##Lancement 

pour lancer le programme on doit utiliser la commande 

```
go build && ./project-particles
```
ensuite vous pouvez profité du projet 

##Guide

Notre SAE est très simple car tout est déjà automatisé donc lorsqu'on lance le projet, avec la commande vue ci-dessus, une page noire avec des énormes fleches apparaîtra.
Pour pouvoir continuer il suffit simplement de suivre l'instruction mise sur l'écran et de cliquer sur la petite fleche à droite.
une fois cette fleche sélectionnée, 7 petites icônes apparaîtra. il suffira de cliquer sur l'icône pour pouvoir lancer la fonction que vous désirez. Si vous voulez changer de fonction il suffit simplement de resélectionner la fleche à droite. Vous pouvez aussi mettre le système en Pause en cliquant sur le espace. et si jamais vous voulez réinitialiser le système vous avez un bouton spécial qui remet à zéro le système.    

##Information
Comme cité juste avant notre projet est un projet où il y a juste besoin de lancer le système car tout est automatisé.
bien évidemment si jamais on veut modifier des paramètres, il suffit de se rendre dans le draw à partir de la ligne 152. Ne vous inquiétez pas bien que le fichier ne soit pas un json il fonctionne de la même façon et il y a des commentaires pour pas que vous soyer perdu dû à la longueur de chaque fonction donc vous pouvez changer les paramètres comme bon vous sembles.

Dans notre projet il y a un total de 7 fonctions déjà intégrer
- [ ] une fonction Cercle
- [ ] une fonction Draw
- [ ] une fonction Glue
- [ ] une fonction Snow
- [ ] une fonction Pattern
- [ ] une fonction Gravite
- [ ] une fonction Rainbow

##Cercle 
La fonction cercle est une fonction qui crée un cercle et qui suit la souris 

##Draw
la fonction Draw est une petite fonction où chaque particule se suit en fonction de la souris ce qui donne l'illusion qu'on dessine 

##Glu
La fonction Glu est une fonction ou chaque particule est recouverte de colle et si elle rencontre une autre particule ou un bord se colle

##Snow 
La fonction Snow est une fonction qui donne l'impression que de la neige tombe avec un dégradé de couleur et une opacité random

##Pattern 
La fonction pattern est une fonction ou la particule se suit et rebondit contre chaque mur

##Gravite
La fonction gravite et une fonction où les particules apparaissent au niveau de la souris et sont attirées vers le bas et vont rebondir contre chaque bord

##Rainbow
La fonction Rainbow est une fonction où la particule créée une vaque arc-en-ciel


