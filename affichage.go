package hangman

import (
	"fmt"
	"github.com/SPA-guetty/hangman/src/common"
)

func Showjosé(nb int) {
	var tab []string
	if nb == -1 {
		fmt.Print("\n\n")
		tab = append(tab, common.DisplayJosé+"  +---+"+common.Reset)
		tab = append(tab, common.DisplayJosé+"      |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"      |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" \\O/  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"  |   |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" / \\  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"========="+common.Reset)
	} else if nb == -2 {
		fmt.Print("\n\n")
		tab = append(tab, common.DisplayJosé+"  +---+"+common.Reset)
		tab = append(tab, common.DisplayJosé+"      |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" \\O/  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"  |   |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" / \\  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"      |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"========="+common.Reset)
	} else if nb == -3 {
		tab = append(tab, common.DisplayJosé+"  +---+"+common.Reset)
		tab = append(tab, common.DisplayJosé+"  |   |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"  |   |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" \\O/  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"  |   |"+common.Reset)
		tab = append(tab, common.DisplayJosé+" / \\  |"+common.Reset)
		tab = append(tab, common.DisplayJosé+"========="+common.Reset)
	} else {
		if nb <= 7 {
			tab = append(tab, common.DisplayJosé+"  +---+  "+common.Reset)
		}
		if nb <= 8 {
			tab = append(tab, common.DisplayJosé+"      |  "+common.Reset)
			tab = append(tab, common.DisplayJosé+"      |  "+common.Reset)
			tab = append(tab, common.DisplayJosé+"      |  "+common.Reset)
			tab = append(tab, common.DisplayJosé+"      |  "+common.Reset)
			tab = append(tab, common.DisplayJosé+"      |  "+common.Reset)
		}
		if nb <= 6 {
			tab[1] = common.DisplayJosé + "  |   |  " + common.Reset
		}
		if nb <= 5 {
			tab[2] = common.DisplayJosé + "  O   |  " + common.Reset
		}
		if nb == 4 {
			tab[3] = common.DisplayJosé + "  |   |  " + common.Reset
		}
		if nb == 3 {
			tab[3] = common.DisplayJosé + " /|   |  " + common.Reset
		}
		if nb <= 2 {
			tab[3] = common.DisplayJosé + " /|\\  |  " + common.Reset
		}
		if nb == 1 {
			tab[4] = common.DisplayJosé + " /    |  " + common.Reset
		}
		if nb == 0 {
			tab[4] = common.DisplayJosé + " / \\  |  " + common.Reset
		}
		if nb >= 0 {
			tab = append(tab, common.DisplayJosé+"========="+common.Reset)
		}
	}
	for _, e := range tab {
		fmt.Println(e)
	}
}
