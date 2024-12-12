package fonctions_du_pendu_web

import (
	"fonctions_du_pendu_web/source/hang"

	fdp "github.com/CookieG77/hangman/functions"
)

// Fonction qui crée la structure du jeu avec le package Fonctions Du Pendu
func CreateGameStructure(Game hang.GameData) hang.GameData {
	var hidden string
	Letter_list = nil
	wordList := fdp.ReadWordFile("source/resource/wordListUltimate.txt") // Crée la liste de mots
	randomWord := fdp.GetRandomWord(wordList)                            // Recupère un mot aléatoire dans la liste de mots
	Game = hang.CreateGameStructure(10, hidden, randomWord)              // Crée la structure du jeu
	hang.CreateInvisibleWord(&Game)                                      // Crée le mot invisible
	hang.GetRevealedLetters(&Game, &Letter_list)                         // Recupère les lettres déjà révélé
	return Game
}
