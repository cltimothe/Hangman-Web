package functions

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/fatih/color"
)

var accents map[string]string = map[string]string{
	"é": "e",
	"è": "e",
	"ê": "e",
	"ë": "e",
	"à": "a",
	"â": "a",
	"ä": "a",
	"ù": "u",
	"û": "u",
	"ü": "u",
	"î": "i",
	"ï": "i",
	"ô": "o",
	"ö": "o",
	"ç": "c",
	"ñ": "n",
	"É": "E",
	"È": "E",
	"Ê": "E",
	"Ë": "E",
	"À": "A",
	"Â": "A",
	"Ä": "A",
	"Ù": "U",
	"Û": "U",
	"Ü": "U",
	"Î": "I",
	"Ï": "I",
	"Ô": "O",
	"Ö": "O",
	"Ç": "C",
	"Ñ": "N",
}

type GameData struct {
	health int
	hidden string
	word   string
}

// Fonction pour créer la structure de donnée d'une partie.
func CreateGameStructure(health int, hidden string, randomWord string) GameData {
	return GameData{
		health: health,
		hidden: hidden,
		word:   randomWord,
	}
}

func GameStrucGetHealth(game GameData) int {
	return game.health
}

func GameStrucGetHidden(game GameData) string {
	return game.hidden
}

func GameStrucGetWord(game GameData) string {
	return game.word
}

func GameStrucSetHealth(game *GameData, val int) {
	(*game).health = val
}

func GameStrucSetHidden(game *GameData, val string) {
	(*game).hidden = val
}

func GameStrucSetWord(game *GameData, val string) {
	(*game).word = val
}

// Fonction temporaire pour pas que le main ne pète un cable
func Start(words []string, asciiArt []string, logo string, alphabet []string, load_save bool, save map[string]interface{}) {
	var game GameData
	if load_save {
		game = CreateGameStructure((int)(save["health"].(float64)), save["hidden"].(string), save["randomWord"].(string))
	} else {
		randomWord := GetRandomWord(words)
		game = CreateGameStructure(9, "", randomWord)
		CreateInvisibleWord(&game)
	}
	ShowWord(game.hidden)
	GameLoop(&game, words, asciiArt, logo, alphabet)
}

// Boucle du jeu
func GameLoop(game *GameData, words []string, asciiArt []string, logo string, alphabet []string) {
	playing := true
	tmp := ""
	for playing {
		ClearCmd()
		PrintLogo(logo)
		PrintHangman(asciiArt, game.health)
		tmp = GetPlayerInput(game, &playing, words, asciiArt, logo, alphabet)
		if game.health < 0 || tmp == "stop" {
			playing = false
			break
		}
		(*game).hidden = tmp
	}
	if tmp == "stop" {
		SaveGame(game)
	}
}

// Récupère un mot aléatoire dans la liste des mots
func GetRandomWord(words []string) string {
	return words[rand.IntN(len(words))]
}

// Place la lettre entrée par le joueur dans le mot caché
func PlaceLetterInWord(game *GameData, letter string) string {
	var res string
	for i, l := range (*game).word {
		replace := false
		if string(l) == letter {
			replace = true
			res += string(l)
		}
		if !replace {
			res += string((*game).hidden[i])
		}
	}
	return res
}

// Créee le mot invisible, prend en argument le mot
// généré ainsi qu'un string vide
func CreateInvisibleWord(game *GameData) {
	for _, c := range (*game).word {
		if c == ' ' {
			(*game).hidden += " "
		} else {
			(*game).hidden += "_"
		}
	}
	// Enfin on révèle n lettres du mot basé sur la longueur du mot
	for revealedLetter := 0; revealedLetter < (len((*game).word)/2)-1; {
		letter := string((*game).word[rand.IntN(len((*game).word))])
		if letter != " " && !strings.Contains((*game).hidden, letter) {
			(*game).hidden = PlaceLetterInWord(game, letter)
			revealedLetter += 1
		}
	}
}

// Affiche le mot à trouver
func ShowWord(hidden string) {
	fmt.Println(hidden)
}

// Récupère la lettre entrée par le joueur et vérifie si elle est dans le mot
// et renvoi le mot caché avec des lettres révélé ou non.
// Cette fonction gère presque l'entièreté du jeu.
func GetPlayerInput(game *GameData, playing *bool, words []string, asciiArt []string, logo string, alphabet []string) string {
	var res string
	var letter string
	fmt.Println(GetASCIIString(alphabet, (*game).hidden))
	fmt.Println("Ecrivez une lettre:")
	fmt.Scan(&letter)                       // Recupère l'entrée du joueur
	letter = NormalizeText(letter, accents) // retire les accents/majuscule
	if letter == "stop" {                   // Permet l'arret de la partie en cours et de sauvegarder la partie en cours
		return "stop"
	}
	if !IsInputValid(letter) {
		return (*game).hidden
	}
	if IsInHidden((*game).hidden, letter) { // Vérifie si la lettre est déjà visible
		fmt.Println("test")
		return (*game).hidden
	}
	letter = SetLower(letter)
	if len(letter) == 1 {
		// Si la lettre entrée par le joueur est une seule lettre on tente de la remplacer dans le mot caché
		res = PlaceLetterInWord(game, letter)
	} else if letter == (*game).word {
		// Si le joueur tente de deviné le mot complet
		res = letter
	} else {
		res = (*game).hidden
	}
	if res == (*game).hidden && len(letter) == 1 {
		// Si le le mot caché n'as pas changer c'est que le joueur à deviné une mauvaise lettre, on enlève donc un point de vie
		(*game).health -= 1

	} else if letter != (*game).word && len(letter) > 1 {
		// Si le joueur tente de deviné le mot complet et qu'il se trompe, on enlève 2 points de vie
		(*game).health -= 2
	}
	if (*game).word == res || (*game).health == 0 || letter == (*game).word {
		for bonneReponse := false; !bonneReponse; {
			ClearCmd()
			PrintLogo(logo)
			PrintHangman(asciiArt, (*game).health)
			var answer string
			switch (*game).word {
			case res:
				color.HiGreen("\nVous avez gagné !\nLe mot était:\n%s\n", GetASCIIString(alphabet, (*game).word))
			default:
				color.Red("\nVous avez perdu !\nLe mot était:\n%s\n", GetASCIIString(alphabet, (*game).word))
			}
			fmt.Println("Voulez-vous rejouer (oui/non) ?")
			fmt.Scan(&answer)
			answer = strings.ToLower(answer)
			switch answer {
			case "oui":
				*playing = false
				bonneReponse = true
				Start(words, asciiArt, logo, alphabet, false, nil)
				return "restart"
			case "non":
				*playing = false
				bonneReponse = true
				return "fin"
			}
		}

	}
	return res
}

func CheckPlayerInput(game *GameData, input string) {
	input = NormalizeText(strings.ToLower(input), accents) // Retire les accents
	res := PlaceLetterInWord(game, input)
	if len(input) == 1 {
		if res == (*game).hidden {
			(*game).health -= 1
		}
	} else {
		if res != (*game).word {
			(*game).health -= 2
		}
	}
}

// Fonction pour vérifier si la lettre entrée est déjà révélé dans le mot
// Args: hidden = mot vu par le joueur, letter = lettre entrée en paramètre
// Condition: à lancer uniquement si l'input du joueur est égal à 1
func IsInHidden(hidden string, letter string) bool {
	for _, c := range hidden {
		if letter == string(c) {
			return true
		}
	}
	return false
}

// Modifie la lettre entrée par le joueur pour la mettre en minuscule
func SetLower(letter string) string {
	//res := ""
	for _, c := range letter {
		if c >= 'A' && c <= 'Z' {
			c += 32
			return string(c)
		} else if c >= 'a' && c <= 'z' {
			return letter
		}
	}
	return "error"
}

// Vérifie si le mot entrée est valide ou non
func IsInputValid(input string) bool {
	if len(input) == 0 {
		return false
	}
	for _, c := range input {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= 129 && c <= 165) {
			return false
		}
	}
	return true
}

func NormalizeText(input string, accents map[string]string) string {
	result := input
	for accent, replacement := range accents {
		result = strings.ReplaceAll(result, accent, replacement)
	}
	return result
}

// Affiche le logo du jeu
func PrintLogo(logo string) {
	logo_color := color.New(color.FgHiYellow)
	logo_color.Println(logo)
}

func PrintHangman(ASCIIart []string, hp int) {
	hangman_color := color.New(color.FgGreen)
	switch hp {
	case 0:
		hangman_color = color.New(color.FgRed)
	case 1:
		hangman_color = color.New(color.FgHiRed)
	case 2:
		hangman_color = color.New(color.FgYellow)
	case 3:
		hangman_color = color.New(color.FgHiYellow)
	case 4:
		hangman_color = color.New(color.FgHiGreen)
	}
	hangman_color.Println(ASCIIart[9-hp])
}

func GetASCIIString(ASCIIart []string, hidden string) string {
	var res string
	var lines []string
	for _, c := range hidden {
		tmp := strings.Split(ASCIIart[c-31], "\n")
		for i, line := range tmp {
			if i >= len(lines) {
				lines = append(lines, "")
			}
			lines[i] += line
		}
	}
	for _, line := range lines {
		res += line + "\n"
	}
	return res
}

// Permet de sauvegarder la partie en cours
func SaveGame(game *GameData) {
	save := map[string]interface{}{
		"health":     (*game).health,
		"hidden":     (*game).hidden,
		"randomWord": (*game).word,
	}
	file, _ := json.MarshalIndent(save, "", " ")
	os.WriteFile("save.txt", file, 0644)
}
