package fonctions_du_pendu_web

import (
	"fonctions_du_pendu_web/source/hang"
	"log"
	"net/http"
)

// Rien de vulgaire btw, fdp = Fonctions du Pendu (≧∀≦)ゞ
//   					 fdpw = Fonctions du Pendu Web (UwU)

type PageData struct {
	Word       string
	HiddenWord string
	Name       string
	Health     string
}

var Game hang.GameData

var Letter_list []string

// Fonction qui fait ce que devrait faire le main (lancement du jeu) mais pas dans
// le main parce que sinon c'est pas jolie.
func Launch() {
	// Crée la structure du jeu (hp, mot, mot caché)
	Game = CreateGameStructure(Game)

	// Support du fichier css
	fs := http.FileServer(http.Dir("./source/web"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Lie les pages a une fonction
	http.HandleFunc("/", launchHomePage)
	http.HandleFunc("/hardcore-gaming", launchGamePage)
	http.HandleFunc("/gameover", launchGameoverPage)
	http.HandleFunc("/win", launchWinPage)

	// Démarre l'hébergement
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
Les fonctions en dessus servent a rediriger vers les bonnes pages.
(Decalé dans un autre fichier pour pas que le fichier soit trop grand)
*/

// Charge la page d'Accueil
func launchHomePage(w http.ResponseWriter, r *http.Request) {
	HomePage(w, r)
}

// Charge la page du Jeu
func launchGamePage(w http.ResponseWriter, r *http.Request) {
	GamePage(w, r)
}

// Charge la page de Victoire
func launchWinPage(w http.ResponseWriter, r *http.Request) {
	WinPage(w, r)
}

// Charge la page de Défaite
func launchGameoverPage(w http.ResponseWriter, r *http.Request) {
	GameoverPage(w, r)
}
