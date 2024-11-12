package game

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"

	"github.com/SPA-guetty/hangman/src/common"
)

func centerText(text string, width int) string {
    lines := strings.Split(text, "\n")
    centeredLines := make([]string, len(lines))

    for i, line := range lines {
        // Garde la ligne telle qu'elle est pour l'ASCII art
        trimmedLine := line
        displayWidth := runewidth.StringWidth(trimmedLine)
        if displayWidth < width {
            padding := (width - displayWidth) / 2
            centeredLines[i] = strings.Repeat(" ", padding) + trimmedLine
        } else {
            centeredLines[i] = trimmedLine
        }
    }

    return strings.Join(centeredLines, "\n")
}

func Menu(gameInstance *common.Game) {
	var answer string

	common.ClearTerminal()
	// Obtenir la largeur du terminal
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Erreur: Le terminal était trop petit.")
	}

	asciiArt := `
██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`

	centeredAsciiArt := centerText(asciiArt, width)

	fmt.Print(common.Title + centeredAsciiArt + "\n" + common.Reset)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println(centerText(common.DisplayJosé+"Pour lancer le jeu: 1", width))
	fmt.Println(centerText(fmt.Sprintf("Difficulté: 2 (actuellement: %s)", PrintDif(gameInstance)), width))
	fmt.Println(centerText("Code de triche: 3", width))
	fmt.Println(centerText("Jeu en ligne: 4", width))
	fmt.Println(centerText("Sortir du jeu: 5", width))
	fmt.Print(centerText("Choix joueur :", width))
	fmt.Scan(&answer)
	for {
		if answer == "1" {
			StartGame(gameInstance)
		} else if answer == "2" {
			ChangeDifficulty(gameInstance)
		} else if answer == "3" {
			CheatCodes(gameInstance)
			return
		} else if answer == "4" {
			// OnlineMenu(gameInstance)
			common.ClearTerminal()
			Menu(gameInstance)
		} else if answer == "5" {
			common.ClearTerminal()
			os.Exit(0)
		} else {
			fmt.Println(centerText(common.ErrorText+"Invalid input", width))
			fmt.Print(centerText("Choix joueur : ", width))
			fmt.Scan(&answer)
		}
	}
}
