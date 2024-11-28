package hang

import (
	"html/template"
	"net/http"
)

// Data structure to hold dynamic data
type PageData struct {
	Content string
}

func test() {
	print("est")
}

func startGameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/web/hangmanUltimateWebEditionV02.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Data to inject into the template
	data := PageData{
		Content: "This content is dynamically generated using Go templates.",
	}

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
