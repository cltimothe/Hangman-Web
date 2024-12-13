package fonctions_du_pendu_web

import (
	"fmt"
	"fonctions_du_pendu_web/source/hang"
	"net/http"
	"text/template"
)

// Fonction qui gère le jeu. (Page dynamique)
func GamePage(w http.ResponseWriter, r *http.Request) {
	// Chargment du Template
	tmpl, err := template.ParseFiles("source/web/PlaceholderHangmanUltimateWebEditionV02.html")
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
			res = hang.CheckPlayerInput(&Game, &input, &Letter_list)

			// Met a jour le mot caché
			hang.GameStrucSetHidden(&Game, res)
		}
		// Debug/Test
		fmt.Println(Letter_list, Game)

		// Met a jour la valeur "Input" de la page avec l'entré du joueur
		data.Input = input

		if res == hang.GameStrucGetWord(Game) {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
			return
		}
	}
	// Met à jour la page avec les nouvelles valeurs
	data.Word = hang.GameStrucGetWord(Game)
	data.HiddenWord = hang.GameStrucGetHidden(Game)
	data.Health = hang.GameStrucGetHealth(Game)

	// Verifie la vie du joueur et renvoi sur la page de gameover si mort
	pl_health := hang.GameStrucGetHealth(Game)
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
