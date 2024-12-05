package functions

import (
	"fmt"
	"os"
	"os/signal"
	"slices"
)

func Launch() {
	// Simple code pour nettoyer le terminal à la fin du programme par Crtl+C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				ClearCmd()
				os.Exit(1)
			}
		}
	}()

	// Chargement des ressources du jeu
	var Logo string
	var Words []string
	ASCIIart := ReadASCIIArtFile("docs/hangman.txt")
	ASCIIart_alphabet := ReadASCIIArtFile("docs/standard.txt")

	// Gestion des arguments et lancement du jeu
	Args := GetArgs()
	if slices.Contains(Args, "-df") {
		go PlayLoopMusic("docs/deepfriedmusic.mp3")
		Logo = ReadLogoFile("docs/deepfriedlogo.txt")
		Words = ReadWordFile("docs/deepfriedwords.txt")
	} else {
		go PlayLoopMusic("docs/music.mp3")
		Logo = ReadLogoFile("docs/logo.txt")
		Words = ReadWordFile("docs/words.txt")
	}
	if slices.Contains(Args, "--letterFile") {
		ASCIIart_alphabet = ReadASCIIArtFile(Args[slices.Index(Args, "--letterFile")+1])
	}
	if slices.Contains(Args, "--startWith") {
		if len(Args) > slices.Index(Args, "--startWith")+1 {
			var save map[string]interface{}
			err := LoadSave(&save, Args[slices.Index(Args, "--startWith")+1])
			if err {
				ClearCmd()
				fmt.Println("Aucune sauvegarde n'a été trouvée, ou elle est corrompue.\n Assurez-vous que le fichier de sauvegarde soit correct.")
				return
			}
			Start(Words, ASCIIart, Logo, ASCIIart_alphabet, true, save)
		} else {
			ClearCmd()
			fmt.Println("Aucune sauvegarde n'a été fournit.")
			return
		}
	} else {
		Start(Words, ASCIIart, Logo, ASCIIart_alphabet, false, nil)
	}
	ClearCmd()
}
