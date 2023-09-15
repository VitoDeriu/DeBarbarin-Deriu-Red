package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"os"
	"os/exec"
	"time"

	term "github.com/nsf/termbox-go"
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func LoadingScreen() {

	ClearTerminal()
	fmt.Println("                                                                               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("          PROJET  RED                                    /ito Deriu            ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                                                     \\  /\\                     ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                                                      \\/¯¯\\ntoine de Barbarin  ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                             -|             |-                                 ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("         -|                  [-_-_-_-_-_-_-_-]                  |-             ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("         [-_-_-_-_-]          |             |          [-_-_-_-_-]             ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("          | o   o |           [  0   0   0  ]           | o   o |              ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     |    -|       |           |       |-    |     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     |_-___-___-___-|         |-___-___-___-_|     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |  o  ]              [    0    ]              [  o  |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     ]   o   o   o  [ _______ ]  o   o   o   [     | ----__________")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("_____----- |     ]              [ ||||||| ]              [     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     ]              [ ||||||| ]              [     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("       _-_-|_____]--------------[_|||||||_]--------------[_____|-_-_           ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("      ( (__________------------_____________-------------_________) )          ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                                                                               ")
	time.Sleep(time.Second * 2)
}

func PrincipalMenu() {

	pointingAt := 1
	var selectedOption int

	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	ClearTerminal()
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓              /(                                                            ▓")
	fmt.Println("▓          O\\\\\\{}============-      Nouvelle partie                          ▓")
	fmt.Println("▓              \\(                                                            ▓")
	fmt.Println("▓                                   Écran de chargement                      ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                   Bonus                                    ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                   Crédits                                  ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                   Quitter                             /    ▓")
	fmt.Println("▓                                                                   \\  /\\    ▓")
	fmt.Println("▓                                                                    \\/¯¯\\   ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	for selectedOption == 0 {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowUp:
				if pointingAt > 1 {
					pointingAt--
				}
			case term.KeyArrowDown:
				if pointingAt < 5 {
					pointingAt++
				}
			case term.KeyEnter:
				selectedOption = pointingAt
			}
		case term.EventError:
			panic(ev.Err)
		}

		switch pointingAt {
		case 1:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓              /(                                                            ▓")
			fmt.Println("▓          O\\\\\\{}============-      Nouvelle partie                          ▓")
			fmt.Println("▓              \\(                                                            ▓")
			fmt.Println("▓                                   Écran de chargement                      ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Bonus                                    ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Crédits                                  ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Quitter                             /    ▓")
			fmt.Println("▓                                                                   \\  /\\    ▓")
			fmt.Println("▓                                                                    \\/¯¯\\   ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

		case 2:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Nouvelle partie                          ▓")
			fmt.Println("▓              /(                                                            ▓")
			fmt.Println("▓          O\\\\\\{}============-      Écran de chargement                      ▓")
			fmt.Println("▓              \\(                                                            ▓")
			fmt.Println("▓                                   Bonus                                    ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Crédits                                  ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Quitter                             /    ▓")
			fmt.Println("▓                                                                   \\  /\\    ▓")
			fmt.Println("▓                                                                    \\/¯¯\\   ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

		case 3:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Nouvelle partie                          ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Écran de chargement                      ▓")
			fmt.Println("▓              /(                                                            ▓")
			fmt.Println("▓          O\\\\\\{}============-      Bonus                                    ▓")
			fmt.Println("▓              \\(                                                            ▓")
			fmt.Println("▓                                   Crédits                                  ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Quitter                             /    ▓")
			fmt.Println("▓                                                                   \\  /\\    ▓")
			fmt.Println("▓                                                                    \\/¯¯\\   ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

		case 4:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Nouvelle partie                          ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Écran de chargement                      ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Bonus                                    ▓")
			fmt.Println("▓              /(                                                            ▓")
			fmt.Println("▓          O\\\\\\{}============-      Crédits                                  ▓")
			fmt.Println("▓              \\(                                                            ▓")
			fmt.Println("▓                                   Quitter                             /    ▓")
			fmt.Println("▓                                                                   \\  /\\    ▓")
			fmt.Println("▓                                                                    \\/¯¯\\   ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

		case 5:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Nouvelle partie                          ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Écran de chargement                      ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Bonus                                    ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                   Crédits                                  ▓")
			fmt.Println("▓              /(                                                            ▓")
			fmt.Println("▓          O\\\\\\{}============-      Quitter                             /    ▓")
			fmt.Println("▓              \\(                                                   \\  /\\    ▓")
			fmt.Println("▓                                                                    \\/¯¯\\   ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		}
	}

	switch selectedOption {
	case 1:
		var Thorgan char.Character
		Thorgan.CreateCharacter("Thorgan", "Elfe", 999, 90000, 80000, map[string]int{"Potion de vie": 3, "Épée de la mort": 1, "Hache des ténèbres": 2})
		ClearTerminal()
		Thorgan.DisplayInfo()
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 2:
		ClearTerminal()
		LoadingScreen()
		PrincipalMenu()

	case 3:
		ClearTerminal()
		fmt.Println("Bonus")
		fmt.Println("Désolé, il n'y a rien pour le moment !")
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 4:
		ClearTerminal()
		fmt.Println("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Crédits ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println("   ▓                                                                        ▓")
		fmt.Println("   ▓                       /                                PROJET RED      ▓")
		fmt.Println("   ▓                      / ITO                                             ▓")
		fmt.Println("   ▓              \\      /\\                             SEPTEMBRE 2023      ▓")
		fmt.Println("   ▓               \\    /  \\                                                ▓")
		fmt.Println("   ▓                \\  /¯¯¯¯\\ NTOINE                                        ▓")
		fmt.Println("   ▓                 \\/      \\                                              ▓")
		fmt.Println("   ▓                                                                        ▓")
		fmt.Println("   ▓                                                                        ▓")
		fmt.Println("   ▓   |¯¯\\  |¯¯\\  /¯¯\\  |¯¯\\  |    |  /¯¯\\  ¯¯¯¯|¯¯¯¯ ¯¯|¯¯  /¯¯\\  |\\   |  ▓")
		fmt.Println("   ▓   |__/  |__/ /    \\ |   \\ |    | /          |       |   /    \\ | \\  |  ▓")
		fmt.Println("   ▓   |     |\\   \\    / |   / |    | \\          |       |   \\    / |  \\ |  ▓")
		fmt.Println("   ▓  _|_    | \\   \\__/  |__/  |____|  \\__/      |     __|__  \\__/  |   \\|  ▓")
		fmt.Println("   ▓                                                                        ▓")
		fmt.Println("   ▓                                                                        ▓")
		fmt.Println("   ▓                                                                   Ⓒ Ⓡ  ▓")
		fmt.Printf("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 5:
		ClearTerminal()
		fmt.Println("Bye bye !")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}
}
