package fonctions_du_pendu_web

import (
	fdp "github.com/CookieG77/hangman/functions"
)

// Fonction qui crée la structure du jeu avec le package Fonctions Du Pendu
func CreateGameStructure(Game fdp.GameData) fdp.GameData {
	var hidden string
	Letter_list = nil
	wordList := fdp.ReadWordFile(Difficulty)               // Crée la liste de mots
	randomWord := fdp.GetRandomWord(wordList)              // Recupère un mot aléatoire dans la liste de mots
	Game = fdp.CreateGameStructure(10, hidden, randomWord) // Crée la structure du jeu
	fdp.CreateInvisibleWord(&Game)                         // Crée le mot invisible
	fdp.GetRevealedLetters(&Game, &Letter_list)            // Recupère les lettres déjà révélé
	return Game
}
