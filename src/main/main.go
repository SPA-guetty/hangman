package main

import (
	"ytrack.learn.ynov.com/git/orthomas/hangman/src/game"
	"ytrack.learn.ynov.com/git/orthomas/hangman/src/initalization"
)

func main() {
	gameInstance, err := initalization.InitGame("dictionnary/medium.txt")
	if err != nil {
		game.HandleError(err)
	}
	game.Menu(gameInstance)
}
