package main

import (
	"fmt"
	"net/http"

	hgd "github.com/CookieG77/hangman/functions"
)

func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Hangman game!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Hangman-Web!")
		hgd.ClearCmd()
	})

	http.HandleFunc("/hangman", hangmanHandler)

	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
