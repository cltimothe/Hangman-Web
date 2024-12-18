# 💀🪢 Hangman WEB ULTIMATE EDITION DELUXE (Projet Hangman WEB)
Ce Projet est un travail réalisé en duo pour notre première année d'étude à Ynov.<br>
Il s'agit d'un simple jeu du pendu. L'objectif est de deviné le mot à l'écran en ne voyant que le nombre de caractère et d'espace. À chaque erreur, le nombre de tentative augmente et un dessin d'un pendu se dessine. Quand le dessin du pendu est complet le joueur a perdu.<br>
## ⚙️🏳️‍🌈Liste des Fonctionnalités
- 🎲 **Mots aléatoires** : Chaques parties a un mot à deviner différent.
- 🌐 **Affichage en HTML/CSS** : L'interface du jeu a été réalisé intégralement en HTML et CSS.
- 🧠 **Un pendu en HTML/CSS** : Le pendu a lui aussi été adapté pour la version WEB.<br>
## 🧑‍🦼‍➡️ Installation
Premièrement il vous faut installer la dernière version dans la branche `main` en cliquant sur le bouton **CODE** (en haut à droite):
Il est aussi possible d'utiliser la commande ci-dessous si vous possédez git :
```bash
git clone https://github.com/cltimothe/Hangman-Web.git
cd Hangman-Web
```
## 🚀 Lancement
Pour lancer le jeu il suffit de lancer la fonction **main** avec la commande:
```bash
go run .\source\main.go
```
## 📁 Fichiers du Projet
- **source/main.go** : Fichier principale du projet, redirige sur le fonction de lancement.
- **source/pour_se_pendre.go** : Gère le lancement du jeu. 
- **source/create_game_structure.go** : Crée la structure du jeu (mot aléatoire, vie, mot caché, etc).
- **source/game_handler.go** : Fait le lien entre le frontend et le backend (Page et Jeu).
- **source/page_manager.go** : Sert à rediriger vers les différentes pages du jeu.
- **source/web/** : Dossier contenant tous les fichiers WEB (Html et CSS).
- **source/resource/** : Dossier contenant les liste des mots (.txt) ainsi que l'icône du jeu (.png).
## 👑 Auteurs
Développé par [Timothé CLEMENT](https://ytrack.learn.ynov.com/git/cltimothe), [Daniel NAKAD](https://ytrack.learn.ynov.com/git/ndaniel)<br>
