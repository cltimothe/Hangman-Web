package fonctions_du_pendu_web

import (
	"log"
	"net/http"
	"text/template"

	fdp "github.com/CookieG77/hangman/functions"
)

// Rien de vulgaire btw, fdp = Fonctions du Pendu (≧∀≦)ゞ
//   					 fdpw = Fonctions du Pendu Web (UwU)

type PageData struct {
	Word       string
	HiddenWord string
	Name       string
}

var Game fdp.GameData

// Fonction qui fait ce que devrait faire le main (lancement du jeu) mais pas dans
// le main parce que sinon c'est pas jolie n'est-ce pas ?
func MainMaisPasDansLeMain() {
	Game = CreateGameStructure(Game)
	http.HandleFunc("/", handler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// J'ai le droit de mettre cette fonction ici parce que je suis pas dans le main UwU
// Sinon la fonction sert a charger la page.
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/web/PlaceholderHangmanUltimateWebEditionV02.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Initialize PageData with default values
	data := PageData{
		Word:       fdp.GameStrucGetWord(Game),   // Example value for the word
		HiddenWord: fdp.GameStrucGetHidden(Game), // Example hidden word
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to process form data", http.StatusBadRequest)
			return
		}
		// Update PageData with the submitted name
		data.Name = r.FormValue("name")
		// playing := true
		// fdp.GetPlayerInput(Game, &playing)
	}
	// Render the template with the dynamic data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
