package fonctions_du_pendu_web

import fdp "github.com/CookieG77/hangman/functions"

// Fonction qui crée la structure du jeu avec fdp
func CreateGameStructure(Game fdp.GameData) fdp.GameData {
	var hidden string
	wordList := fdp.ReadWordFile("source/resource/wordListUltimate.txt")
	randomWord := fdp.GetRandomWord(wordList)
	Game = fdp.CreateGameStructure(10, hidden, randomWord)
	fdp.CreateInvisibleWord(&Game)
	return Game
}
