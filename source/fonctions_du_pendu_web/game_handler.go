package fonctions_du_pendu_web

import (
	"fmt"
	"net/http"
	"text/template"

	fdp "github.com/CookieG77/hangman/functions"
)

// Fonction qui gère le jeu. (Page dynamique)
func GamePage(w http.ResponseWriter, r *http.Request) {
	// Chargment du Template
	tmpl, err := template.ParseFiles("source/web/HangmanWebEdition.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Les Données dynamique de la page
	data := PageData{
		Word:       "", // Mot à trouver
		HiddenWord: "", // Mot cacher
		Health:     10, // Vie
	}
	// Vérifie si le joueur a entré une valeur
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to process form data", http.StatusBadRequest)
			return
		}
		// Recupère l'entré du Joueur
		input := r.FormValue("input")
		var res string
		if len(input) > 0 {
			// Fonction qui gère le jeu (expliqué si on passe notre curseur sur la fonction)
			res = fdp.CheckPlayerInput(&Game, &input, &Letter_list)

			// Met a jour le mot caché
			fdp.GameStrucSetHidden(&Game, res)
		}
		// Debug/Test
		fmt.Println(Letter_list, Game)

		// Met a jour la valeur "Input" de la page avec l'entré du joueur
		data.Input = convertListToString(Letter_list)

		if res == fdp.GameStrucGetWord(Game) {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
			return
		}
	}
	// Met à jour la page avec les nouvelles valeurs
	data.Word = fdp.GameStrucGetWord(Game)
	data.HiddenWord = fdp.GameStrucGetHidden(Game)
	data.Health = fdp.GameStrucGetHealth(Game)

	// Verifie la vie du joueur et renvoi sur la page de gameover si mort
	pl_health := fdp.GameStrucGetHealth(Game)
	if pl_health <= 0 {
		http.Redirect(w, r, "/gameover", http.StatusSeeOther)
		return
	}
	// Lance la page
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func convertListToString(list []string) string {
	res := ""
	for _, c := range list {
		res += string(c) + " "
	}
	return res
}
