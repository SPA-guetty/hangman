package main

import (
	"hangman/src/game"
	"hangman/src/initalization"
)

func main() {
	gameInstance, err := initalization.InitGame("dictionnary/medium.txt")
	if err != nil {
		game.HandleError(err)
	}
	game.Menu(gameInstance)
}
