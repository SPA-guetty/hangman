package game

import (
	"fmt"
	"os"
	"strings"

	"ytrack.learn.ynov.com/git/orthomas/hangman/src/common"

	"golang.org/x/term"
)

func ChangeDifficulty(gameInstance *common.Game) {
    // Obtenir la largeur du terminal
    width, _, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        fmt.Println("Erreur: Le terminal est trop petit.")
        width = 80 // Définir une largeur par défaut en cas d'erreur
    }

    common.ClearTerminal()

    asciiArt := `
██████╗ ██╗███████╗███████╗██╗ ██████╗██╗   ██╗██╗     ████████╗██╗   ██╗
██╔══██╗██║██╔════╝██╔════╝██║██╔════╝██║   ██║██║     ╚══██╔══╝╚██╗ ██╔╝
██║  ██║██║█████╗  █████╗  ██║██║     ██║   ██║██║        ██║    ╚████╔╝ 
██║  ██║██║██╔══╝  ██╔══╝  ██║██║     ██║   ██║██║        ██║     ╚██╔╝  
██████╔╝██║██║     ██║     ██║╚██████╗╚██████╔╝███████╗   ██║      ██║   
╚═════╝ ╚═╝╚═╝     ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝   ╚═╝      ╚═╝   `
    centeredAsciiArt := centerText(asciiArt, width)

    // Afficher l'art ASCII centré
    fmt.Print(common.Title + centeredAsciiArt + "\n" + common.Reset)

    // Ajouter une séparation visuelle
    fmt.Println(centerText(strings.Repeat("-", width), width))

    // Afficher les options de difficulté, centrées
    fmt.Println(common.DisplayJosé + centerText("Facile: 1", width))
    fmt.Println(centerText("Moyen: 2", width))
    fmt.Println(centerText("Difficile: 3", width))
    fmt.Print(centerText("Votre choix : ", width))

    var answer string
    fmt.Scan(&answer)

    // Mise à jour du chemin en fonction du choix de l'utilisateur
    if answer == "1" {
        gameInstance.Path = "dictionnary/easy.txt"
    } else if answer == "2" {
        gameInstance.Path = "dictionnary/medium.txt"
    } else if answer == "3" {
        gameInstance.Path = "dictionnary/hard.txt"
    }

    common.ClearTerminal()
    Menu(gameInstance)
}

func PrintDif(gameInstance *common.Game) string {
	if gameInstance.Path == "dictionnary/easy.txt" {
		return "Facile"
	} else if gameInstance.Path == "dictionnary/medium.txt" {
		return "Moyen"
	} else {
		return "Difficile"
	}
}
