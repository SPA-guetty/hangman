package initalization

import (
	"fmt"
	"hangman/src/common"
	"hangman/src/gamelogic"
	"hangman/src/wordmanager"
)

func InitGame(filepath string) (*common.Game, error) {
	gameInstance, err := Init(filepath)
	if err != nil {
		return nil, fmt.Errorf(common.ErrorText+"error initializing game: %w"+common.Reset, err)
	}
	return gameInstance, nil
}

func Init(filePath string) (*common.Game, error) {
	dictionnary := wordmanager.LoadWords(filePath)
	selectedWord, err := wordmanager.SelectRandomWord(dictionnary)
	if err != nil {
		return nil, err
	}
	gameInstance := &common.Game{
		Path: filePath,
		Word: selectedWord,
		AttemptsLeft: 10,
		GuessedLetters: []rune{},
		RevealedLetters: make([]rune, len(selectedWord)),
	}
	gamelogic.RevealRandomLetters(gameInstance)
	var cheattab []string
	cheattab = append(cheattab, "addonelife")
	cheattab = append(cheattab, "giveletter")
	cheattab = append(cheattab, "muchvowels")
	cheattab = append(cheattab, "unusedletters")
	cheattab = append(cheattab, "sacrifice")
	gameInstance.CheatCodes = cheattab

	return gameInstance, nil
}
