package gamelogic

import (
	"math/rand"
	"os"

	"ytrack.learn.ynov.com/git/orthomas/hangman/src/common"
)

func RevealRandomLetters(game *common.Game) {
	wordLen := len(game.Word)
	game.RevealedLetters = make([]rune, wordLen)
	revealedIndices := make([]int, 0)
	for i := 0; i < wordLen; i++ {
		game.RevealedLetters[i] = '_'
	}
	n := (wordLen / 2) - 1
	if n <= 0 {
		os.Exit(1)
	}
	// Boucle pour vérifier n lettres aléatoires au lancement du programme
	for len(revealedIndices) < n {
		randomIndex := rand.Intn(wordLen)
		if !common.Contains(revealedIndices, randomIndex) {
			revealedIndices = append(revealedIndices, randomIndex)
			game.RevealedLetters[randomIndex] = rune(game.Word[randomIndex])
		}
	}
}
