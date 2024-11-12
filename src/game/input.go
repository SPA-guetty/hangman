package game

import (
	"fmt"
	"github.com/SPA-guetty/hangman"
	"github.com/SPA-guetty/hangman/src/common"
)

func Checkletter(game *common.Game, letter string) string {for {
		cheatcode := false
		if len(letter) != 1 {
			if letter == game.Word {
				Victory(game)
			} else if !IsACheatCode(letter, game) {
				return "err1"
				continue
			} else {
				cheatcode = true
			}
		}
		if letter >= "A" && letter <= "Z" {
			letter = string(letter[0] + 32)
		}
		if letter < "a" || letter > "z" {
			return "err2"
		}
		if !cheatcode && !IsUsed(letter, game) {
			game.GuessedLetters = append(game.GuessedLetters, rune(letter[0]))
		}
		return letter
	}
}

func PrintUpdatedWord(game *common.Game) {
	for _, char := range game.RevealedLetters {
		fmt.Print(string(char) + "")
	}
	fmt.Println()
}

func IsUsed(letter string, game *common.Game) bool {
	char := rune(letter[0])
	list := game.GuessedLetters
	for _, e := range list {
		if e == char {
			return true
		}
	}
	return false
}

func Input(game *common.Game, letter string) {
	isIn := common.Containsstr(game.Word, letter)
	if isIn {
		for i, char := range game.Word {
			if string(char) == letter {
				game.RevealedLetters[i] = rune(letter[0])
			}
		}
		PrintUpdatedWord(game)
	} else if IsACheatCode(letter, game) {
		UseCheatCode(letter, game)
	} else if len(letter)==1 {
		game.AttemptsLeft --
		fmt.Printf(common.WarningText+"Pas présent dans le mot, il vous reste %d essais\n"+common.Reset, game.AttemptsLeft)
		game.CountErrors++
		hangman.Showjosé(game.AttemptsLeft)
		PrintUpdatedWord(game)
	}
}
