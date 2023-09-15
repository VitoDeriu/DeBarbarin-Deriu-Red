package menu

import (
	char "ProjetRed/Character"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func LoadingScreen() {

	ClearTerminal()
	fmt.Println("                                                                               ")
	fmt.Println("                                                                               ")
	fmt.Println("          PROJET  RED                           Antoine de Barbarin            ")
	fmt.Println("                                                                               ")
	fmt.Println("                                                      Vito Deriu               ")
	fmt.Println("                             -|             |-                                 ")
	fmt.Println("         -|                  [-_-_-_-_-_-_-_-]                  |-             ")
	fmt.Println("         [-_-_-_-_-]          |             |          [-_-_-_-_-]             ")
	fmt.Println("          | o   o |           [  0   0   0  ]           | o   o |              ")
	fmt.Println("           |     |    -|       |           |       |-    |     |               ")
	fmt.Println("           |     |_-___-___-___-|         |-___-___-___-_|     |               ")
	fmt.Println("           |  o  ]              [    0    ]              [  o  |               ")
	fmt.Println("           |     ]   o   o   o  [ _______ ]  o   o   o   [     | ----__________")
	fmt.Println("_____----- |     ]              [ ||||||| ]              [     |               ")
	fmt.Println("           |     ]              [ ||||||| ]              [     |               ")
	fmt.Println("       _-_-|_____]--------------[_|||||||_]--------------[_____|-_-_           ")
	fmt.Println("      ( (__________------------_____________-------------_________) )          ")
	fmt.Println()

	time.Sleep(time.Second * 5)
}

func PrincipalMenu() {

	pointingAt := 1
	var selectedOption int
	var displayNumber int

	ClearTerminal()
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                           ➵  Nouvelle partie                               ▓")
	fmt.Println("▓                              Paramètres                                    ▓")
	fmt.Println("▓                              Bonus                                         ▓")
	fmt.Println("▓                              Crédits                                       ▓")
	fmt.Println("▓                              Quitter                                       ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("▓                                                                            ▓")
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	for selectedOption == 0 {
		for displayNumber == 0 {
			consoleReader := bufio.NewReaderSize(os.Stdin, 2)
			input, _ := consoleReader.ReadByte()
			ascii := input

			switch ascii {

			// Arrow up
			case 38:
				if pointingAt > 1 {
					pointingAt--
					displayNumber = pointingAt
				}

				// Arrow down
			case 40:
				if pointingAt < 5 {
					pointingAt++
					displayNumber = pointingAt
				}

				//Enter
			case 13:
				selectedOption = pointingAt

			}
		}

		switch displayNumber {
		case 1:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                           ➵  Nouvelle partie                               ▓")
			fmt.Println("▓                              Paramètres                                    ▓")
			fmt.Println("▓                              Bonus                                         ▓")
			fmt.Println("▓                              Crédits                                       ▓")
			fmt.Println("▓                              Quitter                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			displayNumber = 0

		case 2:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Nouvelle partie                               ▓")
			fmt.Println("▓                           ➵  Paramètres                                    ▓")
			fmt.Println("▓                              Bonus                                         ▓")
			fmt.Println("▓                              Crédits                                       ▓")
			fmt.Println("▓                              Quitter                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			displayNumber = 0

		case 3:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Nouvelle partie                               ▓")
			fmt.Println("▓                              Paramètres                                    ▓")
			fmt.Println("▓                           ➵  Bonus                                         ▓")
			fmt.Println("▓                              Crédits                                       ▓")
			fmt.Println("▓                              Quitter                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			displayNumber = 0

		case 4:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Nouvelle partie                               ▓")
			fmt.Println("▓                              Paramètres                                    ▓")
			fmt.Println("▓                              Bonus                                         ▓")
			fmt.Println("▓                           ➵  Crédits                                       ▓")
			fmt.Println("▓                              Quitter                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			displayNumber = 0

		case 5:
			ClearTerminal()
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Nouvelle partie                               ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Paramètres                                    ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Bonus                                         ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                              Crédits                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓   [O\\\\\\[===============-    Quitter                                       ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("▓                                                                            ▓")
			fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
			displayNumber = 0
		}
	}
	ClearTerminal()
	fmt.Println("Tu as sélectionné le menu nº ", selectedOption)

	switch selectedOption {
	case 1:
		var Thorgan char.Character
		Thorgan.CreateCharacter("Thorgan", "Elfe", 999, 90000, 80000, map[string]int{"Potion de vie": 3, "Épée de la mort": 1, "Hache des ténèbres": 2})
		Thorgan.DisplayInfo()

	case 2:
		ClearTerminal()
		fmt.Println("Paramètres...")
		fmt.Println("En cours de conception")

	case 3:
		ClearTerminal()
		fmt.Println("Bonus")
		fmt.Println("Désolé, il n'y a rien pour le moment !")

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
		fmt.Println("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 5:
		ClearTerminal()
		fmt.Println("Bye bye !")
		os.Exit(0)
	}
}
