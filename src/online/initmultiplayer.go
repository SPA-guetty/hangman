package online

import (
	"context"
	"fmt"
	"ytrack.learn.ynov.com/git/orthomas/hangman/src/common"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// InitMultiplayer initialise Firebase pour le mode multijoueur
func InitMultiplayer() error {
	opt := option.WithCredentialsFile("/home/guireg/hangman/data/hangman.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf(common.ErrorText+"error initializing app: %v"+common.Reset, err)
	}
	_, err = app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf(common.ErrorText+"error initializing Firebase Auth: %v"+common.Reset, err)
	}
	return nil
}
