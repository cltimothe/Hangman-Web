package functions

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

// Renvoie la liste des arguments fournie au programme
func GetArgs() []string {
	return os.Args[1:]
}

// Fonction qui lit chaque ligne du fichier 'file' et les renvoie dans un slice de string
func ReadWordFile(file string) []string {
	file_content, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var res []string
	scanner := bufio.NewScanner(file_content)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return res
}

// Fonction qui lit chaques ligne dans le fichier 'file',
// Et les regroupe par paquet pour formé des dessins en ASCII.
// La fonction considère que les packets sont séparer par une ligne vide.
func ReadASCIIArtFile(file string) []string {
	file_content, err := os.Open(file)
	if err != nil {
		print("Veuillez vérifier que le fichier existe.")
		panic(err)
	}
	var res []string
	scanner := bufio.NewScanner(file_content)
	temp := ""
	for scanner.Scan() {
		if scanner.Text() != "" {
			if temp != "" {
				temp += "\n"
			}
			temp += scanner.Text()
		} else {
			res = append(res, temp)
			temp = ""
		}
	}

	if err = scanner.Err(); err != nil {
		print("Veuillez vérifier que ce fichier est bien un alphabet ascii valide!")
		panic(err)
	}
	file_content.Close()
	return res
}

// Permet de lire le logo du jeu
func ReadLogoFile(file string) string {
	file_content, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var res string
	scanner := bufio.NewScanner(file_content)
	for scanner.Scan() {
		if scanner.Text() != "" {
			res += scanner.Text() + "\n"
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return res[:len(res)-2]
}

// Permet de charger la partie en cours
func LoadSave(res *map[string]interface{}, file_dir string) bool {
	file, err1 := os.ReadFile(file_dir)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

// Permet de vider le terminal
func ClearCmd() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Joue la music donner en boucle sans jamais s'arreter.
func PlayLoopMusic(adresse string) {
	f, err := os.Open(adresse)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	loop := beep.Loop(-1, streamer)
	volume := &effects.Volume{
		Streamer: loop,
		Base:     2,
		Volume:   1,
		Silent:   false,
	}
	speaker.Play(beep.Seq(volume, beep.Callback(func() {
		done <- true
	})))

	<-done
}
