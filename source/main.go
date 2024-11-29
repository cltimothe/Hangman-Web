package main

import (
	"log"
	"net/http"
	"text/template"
)

// PageData holds the dynamic values for the page
type PageData struct {
	Word       string
	HiddenWord string
	Name       string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Initialize PageData with default values
	data := PageData{
		Word:       "PENDU",     // Example value for the word
		HiddenWord: "_ _ _ _ _", // Example hidden word
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
	}

	// Render the page with the dynamic data
	tmpl := `
<!DOCTYPE html>
<html lang="fr">    
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="styles.css" rel="stylesheet" />
    <title>Hangman Ultimate Web Edition</title>   
  </head>
  <body>
    <h1>Ceci est un magnifique placeholder</h1>
    <p>Mot actuel : {{.Word}}</p>
    <p>Mot caché : {{.HiddenWord}}</p>
    {{if .Name}}
      <p>Bienvenue, {{.Name}}!</p>
    {{end}}
    <form action="/" method="post">
      <label for="name">Entrez votre nom :</label>
      <input type="text" id="name" name="name" required>
      <button type="submit">Envoyer</button>
    </form>
  </body>
</html>`
	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
