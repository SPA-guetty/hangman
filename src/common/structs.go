package common

type Game struct {
	Path			string
	Word            string
	AttemptsLeft    int
	GuessedLetters  []rune
	RevealedWord    []rune
	RevealedLetters []rune
	CountErrors     int
	CheatCodes      []string
}

type Hangman struct {
	Stages []string
}

type Dictionnary struct {
	Words []string
}
