package game

import (
	"fmt"
	"github.com/SPA-guetty/hangman/src/common"
	"math/rand"
)

func IsACheatCode(msg string, gameInstance *common.Game) bool {
	for _, e := range gameInstance.CheatCodes {
		if msg == e {
			return true
		}
	}

	return false
}

func UseCheatCode(letter string, gameInstance *common.Game) {
	code := 0
	for i, e := range gameInstance.CheatCodes {
		if e == letter {
			code = i
		}
	}
	if code == 0 {
		// +1 vie
		fmt.Println(common.CheatCode + "Code activé, +1 vie" + common.Reset)
		gameInstance.AttemptsLeft++
		gameInstance.CountErrors--
	} else if code == 1 {
		// Afficher une lettre aléatoire
		fmt.Println(common.CheatCode + "Code activé, une nouvelle lettre est découverte" + common.Reset)
		CheatRandomLetter(gameInstance)
	} else if code == 2 {
		// Nombre de voyelles
		nb := 0
		for _, e := range gameInstance.Word {
			if e == 'a' || e == 'e' || e == 'i' || e == 'o' || e == 'u' || e == 'y' {
				nb++
			}
		}
		fmt.Println(common.CheatCode+"Il y a dans le mot", nb, "voyelles."+common.Reset)
	} else if code == 3 {
		// See used letters
		fmt.Println(common.CheatCode + "Voici les lettres inutilisées" + common.Reset)
		SeeLetters(UnusedLetters(gameInstance))
	} else if code == 4 {
		// Suicide
		gameInstance.AttemptsLeft = 0
	}
}

func CheatRandomLetter(gameInstance *common.Game) {
	//rand
	word := gameInstance.Word
	revealedletters := gameInstance.RevealedLetters
	var missingints []int
	for i := range word {
		if revealedletters[i] == '_' {
			missingints = append(missingints, i)
		}
	}
	randomint := missingints[rand.Intn(len(missingints))]
	gameInstance.RevealedLetters[randomint] = rune(word[randomint])
	fmt.Println(common.CheatCode+"Lettre découverte:", string(word[randomint])+common.Reset)
	fmt.Println(string(gameInstance.RevealedLetters))
}

func UnusedLetters(game *common.Game) []rune {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	var unused []rune
	for _, e := range alphabet {
		isin := false
		for _, letter := range game.GuessedLetters {
			if e == letter {
				isin = true
			}
		}
		if !isin {
			unused = append(unused, e)
		}
	}
	return unused
}
