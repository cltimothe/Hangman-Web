package fonctions_du_pendu_web

import (
	"fmt"
	"fonctions_du_pendu_web/source/hang"
	"net/http"
	"strconv"
	"text/template"
)

// J'ai le droit de mettre cette fonction ici parce que je suis pas dans le main.
// Sinon la fonction sert a (re)charger la page.
func GamePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/web/PlaceholderHangmanUltimateWebEditionV02.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Les Données dynamique de la page
	data := PageData{
		Word:       "", // Example value for the word
		HiddenWord: "", // Example hidden word
		Health:     "", // Example hidden word
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to process form data", http.StatusBadRequest)
			return
		}
		// Mettre a jour la page avec l'input du joueur
		input := r.FormValue("name")
		var res string
		if len(input) > 0 {
			res = hang.CheckPlayerInput(&Game, &input, &Letter_list)
			hang.GameStrucSetHidden(&Game, res)
		}
		fmt.Println(Letter_list)

		data.Name = input

		if res == hang.GameStrucGetWord(Game) {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
			return
		}
	}
	data.Word = hang.GameStrucGetWord(Game)
	data.HiddenWord = hang.GameStrucGetHidden(Game)
	data.Health = strconv.Itoa(hang.GameStrucGetHealth(Game))

	// Verifie la vie du joueur
	pl_health := hang.GameStrucGetHealth(Game)
	if pl_health <= 0 {
		http.Redirect(w, r, "/gameover", http.StatusSeeOther)
		return
	}
	// Render the template with the dynamic data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
