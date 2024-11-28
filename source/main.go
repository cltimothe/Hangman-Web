package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Debug: Check if accueil.html exists
	if _, err := os.Stat("./web/accueil.html"); err != nil {
		fmt.Println("accueil.html not found:", err)
	} else {
		fmt.Println("accueil.html is ready to be served!")
	}

	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/accueil.html")
	})

	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
