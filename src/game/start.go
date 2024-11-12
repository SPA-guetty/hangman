package game

import (
	"fmt"
	"github.com/SPA-guetty/hangman/src/common"
	"github.com/SPA-guetty/hangman/src/wordmanager"
	"os"
	"time"

	"golang.org/x/term"
)

func StartGame(gameInstance *common.Game) {
	common.ClearTerminal()
	
	// Art ASCII à afficher
	asciiArt := `
██████╗ ███████╗██████╗ ██╗   ██╗████████╗    ██████╗ ██╗   ██╗         ██╗███████╗██╗   ██╗
██╔══██╗██╔════╝██╔══██╗██║   ██║╚══██╔══╝    ██╔══██╗██║   ██║         ██║██╔════╝██║   ██║
██║  ██║█████╗  ██████╔╝██║   ██║   ██║       ██║  ██║██║   ██║         ██║█████╗  ██║   ██║
██║  ██║██╔══╝  ██╔══██╗██║   ██║   ██║       ██║  ██║██║   ██║    ██   ██║██╔══╝  ██║   ██║
██████╔╝███████╗██████╔╝╚██████╔╝   ██║       ██████╔╝╚██████╔╝    ╚█████╔╝███████╗╚██████╔╝
╚═════╝ ╚══════╝╚═════╝  ╚═════╝    ╚═╝       ╚═════╝  ╚═════╝      ╚════╝ ╚══════╝ ╚═════╝ `
	
	// Obtenir la largeur du terminal
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Erreur: Le terminal était trop petit.")
		width = 80 // Définir une largeur par défaut en cas d'erreur
	}

	// Centrer l'art ASCII et l'afficher
	centeredAsciiArt := centerText(asciiArt, width)
    fmt.Print(common.Title + centeredAsciiArt + "\n" + common.Reset)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println(common.DisplayJosé + "Bonne chance, vous avez 10 essais." + common.Reset)
	gameInstance.AttemptsLeft = 10
	var newguess []rune

	filepath := gameInstance.Path
	dictionnary := wordmanager.LoadWords(filepath)
	selectedWord, err := wordmanager.SelectRandomWord(dictionnary)
	if err != nil {
		HandleError(err)
	}
	gameInstance.Word = selectedWord
	gameInstance.GuessedLetters = newguess
	gameInstance.RevealedLetters = make([]rune, len(selectedWord))
	for i := range gameInstance.RevealedLetters {
		gameInstance.RevealedLetters[i] = '_'
	}

	fmt.Println(string(gameInstance.RevealedLetters))

	for gameInstance.AttemptsLeft > 0 {
		Input(gameInstance, "")
		if string(gameInstance.RevealedLetters) == gameInstance.Word {
			Victory(gameInstance)
		}
	}
	if gameInstance.AttemptsLeft == 0 {
		fmt.Printf(common.WarningText+"Perdu ! Le mot était : %v"+common.Reset, gameInstance.Word)
		time.Sleep(5 * time.Second)
		common.ClearTerminal()
		Menu(gameInstance)
		fmt.Print('\n')
	}
}
