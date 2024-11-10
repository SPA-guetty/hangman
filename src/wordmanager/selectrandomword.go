package wordmanager

import (
	"fmt"
	"math/rand"

	"ytrack.learn.ynov.com/git/orthomas/hangman/src/common"
)

func SelectRandomWord(dictionnary *common.Dictionnary) (string, error) {
	if len(dictionnary.Words) <= 0 {
		return "", fmt.Errorf(common.ErrorText+"dictionnary is empty"+common.Reset)
	}
	randomWord := rand.Intn(len(dictionnary.Words))
	selectedWord := dictionnary.Words[randomWord]
	return selectedWord, nil
}

