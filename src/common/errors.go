package common

import (
	"fmt"
	"os"
)

func Errors(game *Game, inputLetter string) {
	if len(os.Args) != 2 {
		fmt.Printf(ErrorText + "Vous avez besoin d'un seul argument pour que le programme fonctionne, assurez-vous d'utiliser 2 arguments et réessayez." + Reset)
		os.Exit(84)
	}
	if !Containsstr(game.Word, inputLetter) {
		game.AttemptsLeft--
		fmt.Printf(WarningText+"Pas présent dans le mot, il vous reste %d essais\n"+Reset, game.AttemptsLeft)
		game.CountErrors++
	}
}
