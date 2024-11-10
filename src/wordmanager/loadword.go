package wordmanager

import (
	"bufio"
	"log"
	"os"
	"strings"

	"hangman/src/common"
)

func LoadWords(filePath string) *common.Dictionnary {
	file, err := os.Open(filePath)
	dictionnary := &common.Dictionnary{}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		dictionnary.Words = append(dictionnary.Words, word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dictionnary
}
