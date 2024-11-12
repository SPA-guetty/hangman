package game

import (
	"fmt"
	"github.com/SPA-guetty/hangman"
	"github.com/SPA-guetty/hangman/src/common"
	"os"
	"time"

	"github.com/01-edu/z01"
)

func DisplayWord(selectedWord string) {
	for range selectedWord {
		fmt.Print("_ ")
	}
	fmt.Println()
}

func SeeLetters(list []rune) {
	for _, e := range list {
		z01.PrintRune(e)
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func CheatCodes(gameInstance *common.Game) {
	common.ClearTerminal()
	fmt.Print(common.Title +`
	
 ██████╗██╗  ██╗███████╗ █████╗ ████████╗ ██████╗ ██████╗ ██████╗ ███████╗███████╗
██╔════╝██║  ██║██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔═══██╗██╔══██╗██╔════╝██╔════╝
██║     ███████║█████╗  ███████║   ██║   ██║     ██║   ██║██║  ██║█████╗  ███████╗
██║     ██╔══██║██╔══╝  ██╔══██║   ██║   ██║     ██║   ██║██║  ██║██╔══╝  ╚════██║
╚██████╗██║  ██║███████╗██║  ██║   ██║   ╚██████╗╚██████╔╝██████╔╝███████╗███████║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝
                                                                                  
`+common.Reset)

	fmt.Println(common.CheatCode+"Pour rajouter une vie: addonelife")
	fmt.Println("Pour afficher une lettre aléatoirement: giveletter")
	fmt.Println("Pour connaître le nombre de voyelles du mot: muchvowels")
	fmt.Println("Pour voir les lettres inutilisées: unusedletters")
	fmt.Println("Pour abandonner: sacrifice")

	fmt.Print("\nRentrez n'importe quoi pour fermer les cheatcodes..."+common.Reset)
	var answer string
	fmt.Scan(&answer)
	common.ClearTerminal()
	Menu(gameInstance)
}

func AddWordToDictionary(gameInstance *common.Game) {
    // Demander à l'utilisateur d'entrer un mot à ajouter
    var word string
    fmt.Print("Entrez le mot à ajouter au dictionnaire : ")
    fmt.Scan(&word)

    file, err := os.OpenFile(gameInstance.Path, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier:", err)
        return
    }
    defer file.Close()

    if _, err := file.WriteString(word); err != nil {
        fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
    } else {
        fmt.Println("Mot ajouté au dictionnaire:", word)
    }
}

func Victory(gameInstance *common.Game) {
    for i := 0; i < 6; i++ {
        common.ClearTerminal()
        hangman.Showjosé(-1)
        fmt.Printf(common.DisplayJosé+"Bravo, vous avez trouvé le mot qui était : %v\n"+common.Reset, gameInstance.Word)
        time.Sleep(300 * time.Millisecond)
        common.ClearTerminal()
        hangman.Showjosé(-2)
        fmt.Printf(common.DisplayJosé+"Bravo, vous avez trouvé le mot qui était : %v\n"+common.Reset, gameInstance.Word)
        time.Sleep(300 * time.Millisecond)
    }
    common.ClearTerminal()
    hangman.Showjosé(-1)
    fmt.Printf(common.DisplayJosé+"Bravo, vous avez trouvé le mot qui était : %v\n"+common.Reset, gameInstance.Word)
    time.Sleep(300 * time.Millisecond)
    common.ClearTerminal()

    // Demander à l'utilisateur s'il souhaite ajouter le mot au dictionnaire
    var addWord string
    fmt.Print("Souhaitez-vous ajouter un mot au dictionnaire ? (oui/non) : ")
    fmt.Scan(&addWord)

    if addWord == "oui" {
        AddWordToDictionary(gameInstance)
    }
	Menu(gameInstance)
}

func Sacrifice(gameInstance *common.Game) {
	common.ClearTerminal()
	fmt.Println(common.WarningText+"Vous décidez d'en finir avec José..."+common.Reset)
	hangman.Showjosé(-3)
	fmt.Println(common.DisplayJosé+"Perdu ! Le mot était :"+common.Reset, gameInstance.Word)
	time.Sleep(1 * time.Second)
	common.ClearTerminal()
	fmt.Println(common.WarningText+"Vous décidez d'en finir avec José..."+common.Reset)
	hangman.Showjosé(0)
	fmt.Println(common.DisplayJosé+"Perdu ! Le mot était :"+common.Reset, gameInstance.Word)
	time.Sleep(4 * time.Second)
	common.ClearTerminal()
	Menu(gameInstance)
}
