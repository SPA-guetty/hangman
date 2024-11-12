package main

import (
	"github.com/SPA-guetty/hangman/src/game"
	"github.com/SPA-guetty/hangman/src/initalization"
)

func main() {
	gameInstance, err := initalization.InitGame("dictionnary/medium.txt")
	if err != nil {
		game.HandleError(err)
	}
	game.Menu(gameInstance)
}
