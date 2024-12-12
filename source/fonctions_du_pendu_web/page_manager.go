package fonctions_du_pendu_web

import (
	"net/http"
	"text/template"
)

func GameoverPage(w http.ResponseWriter, r *http.Request) {
	Game = CreateGameStructure(Game) // Recréer une nouvelle partie
	tmpl := template.Must(template.ParseFiles("source/web/gameover.html"))
	tmpl.Execute(w, nil)
}

func WinPage(w http.ResponseWriter, r *http.Request) {
	Game = CreateGameStructure(Game) // Recréer une nouvelle partie
	tmpl := template.Must(template.ParseFiles("source/web/win.html"))
	tmpl.Execute(w, nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("source/web/accueil.html"))
	tmpl.Execute(w, nil)
}
