package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"time"

	rwidth "github.com/mattn/go-runewidth"
)

var MenuLateralBar,
	LogoLittle,
	Sword,
	SoldierArt,
	PrincipalMenuOptions,
	CharCreationMenuText,
	ElfDescription,
	HumanDescription,
	DwarfDescription,
	CharMenuText,
	CharMenuMainGrid [][]rune

var BottomBar,
	CharMenuTitleBar,
	PrincipalMenuTitleBar,
	CharCreationMenuTitleBar,
	CharCreationName,
	CharCreationNameError,
	CharCreationMenuCursor []rune

const (
	PRINCIPAL_MENU     = 0
	CHAR_CREATION_MENU = 1
	CHAR_MENU          = 2
)

func CreateDisplayVariables() {

	MenuLateralBar = append(MenuLateralBar, []rune(" /┃\\ "))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune(" \\┃/ "))

	LogoLittle = append(LogoLittle, []rune("    /"))
	LogoLittle = append(LogoLittle, []rune("\\  /\\"))
	LogoLittle = append(LogoLittle, []rune(" \\/¯¯\\"))

	Sword = append(Sword, []rune("    /("))
	Sword = append(Sword, []rune("O\\\\\\{}============-"))
	Sword = append(Sword, []rune("    \\("))

	SoldierArt = append(SoldierArt, []rune("   .-."))
	SoldierArt = append(SoldierArt, []rune(" __|=|__"))
	SoldierArt = append(SoldierArt, []rune("(_/`-`\\_)"))
	SoldierArt = append(SoldierArt, []rune("//\\___/\\\\"))
	SoldierArt = append(SoldierArt, []rune("<>/   \\<>"))
	SoldierArt = append(SoldierArt, []rune(" \\|_._|/"))
	SoldierArt = append(SoldierArt, []rune("  <_I_>"))
	SoldierArt = append(SoldierArt, []rune("   |||"))
	SoldierArt = append(SoldierArt, []rune("  /_|_\\"))

	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Nouvelle partie"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Écran de chargement"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Bonus"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Crédits"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Quitter"))

	CharCreationMenuText = append(CharCreationMenuText, []rune("Choix de la race :"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Humain"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Elfe"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Nain"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Description :"))

	ElfDescription = append(ElfDescription, []rune("Les elfes se spécialisent dans la magie            "))
	ElfDescription = append(ElfDescription, []rune("et la perception : 80 HP et 120 Mana     "))

	HumanDescription = append(HumanDescription, []rune("Les humains sont équilibrés dans leurs statistiques"))
	HumanDescription = append(HumanDescription, []rune("et leurs compétences : 100 HP et 100 Mana"))

	DwarfDescription = append(DwarfDescription, []rune("Les nains se spécialisent dans l'endurance         "))
	DwarfDescription = append(DwarfDescription, []rune("et leur constitution : 120 HP et 80 Mana "))

	CharMenuText = append(CharMenuText, []rune("Inventaire"))
	CharMenuText = append(CharMenuText, []rune("Équipement"))
	CharMenuText = append(CharMenuText, []rune("Retour"))
	CharMenuText = append(CharMenuText, []rune("Quitter"))
	CharMenuText = append(CharMenuText, []rune("Niveau :"))
	CharMenuText = append(CharMenuText, []rune("Force :"))
	CharMenuText = append(CharMenuText, []rune("Endurance :"))
	CharMenuText = append(CharMenuText, []rune("Perception :"))
	CharMenuText = append(CharMenuText, []rune("Expérience :"))
	CharMenuText = append(CharMenuText, []rune("HP :"))
	CharMenuText = append(CharMenuText, []rune("MP :"))
	CharMenuText = append(CharMenuText, []rune("Or :"))
	CharMenuText = append(CharMenuText, []rune("Équipement :"))
	CharMenuText = append(CharMenuText, []rune("Compétences :"))

	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│ ├────────┬───────────────────────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┬───────────────┤ │"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│├──────────────────────┼─────┼─────┼─────────┼─────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰┴──────────────────────┴─────┴─────┴─────────┴─────╯"))

	BottomBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	PrincipalMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationName = []rune("Nom du personnage (max 20 caractères) :")

	CharCreationNameError = []rune("Erreur : veuillez rentrer un nom valide !")

	CharCreationMenuCursor = []rune("⎯{===-")
}

func LoadingScreen() {
	ClearTerminal()
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
	fmt.Println()
	fmt.Println("Chargement :")
	loadingBar()
	time.Sleep(time.Second * 1)
}

func loadingBar() {
	var loadingBar string
	var loadingPercentage float64
	for i := 0; i < 74; i++ {
		time.Sleep(time.Millisecond * 40)
		loadingBar += "₪"
		loadingPercentage += 1.36
		fmt.Print("\r", int(loadingPercentage), "% ", loadingBar)
	}
}

func displayTopBar(menuNb int) {
	var CurrentTopBar []rune
	switch menuNb {

	case PRINCIPAL_MENU:
		CurrentTopBar = PrincipalMenuTitleBar

	case CHAR_CREATION_MENU:
		CurrentTopBar = CharCreationMenuTitleBar

	case CHAR_MENU:
		CurrentTopBar = CharMenuTitleBar

	}
	column := 0
	for _, char := range CurrentTopBar {
		DisplayRune(column, 0, char)
		column += rwidth.RuneWidth(char)
	}
}

func displayMenuBars() {
	for i := range MenuLateralBar {
		column := 0
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char)
			column += rwidth.RuneWidth(char)
		}
	}
	for i := range MenuLateralBar {
		column := 77
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayBottomBar() {
	column := 0
	for _, char := range BottomBar {
		DisplayRune(column, 17, char)
		column += rwidth.RuneWidth(char)
	}
}

func displayLogo() {
	column := 69
	line := 13
	for i := range LogoLittle {
		column = 69
		for _, char := range LogoLittle[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displaySoldier() {
	column := 57
	line := 3
	for i := range SoldierArt {
		column = 57
		for _, char := range SoldierArt[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayBlankMenu(menuNb int) {
	displayTopBar(menuNb)
	displayMenuBars()
	displayBottomBar()
	if menuNb != CHAR_MENU {
		displayLogo()
	}
	if menuNb == CHAR_CREATION_MENU {
		displaySoldier()
	}
}

func DisplayMenuOptions(menuNb int) {
	var menuOptions [][]rune
	switch menuNb {
	case PRINCIPAL_MENU:
		menuOptions = PrincipalMenuOptions
	}
	column := 36
	line := 4
	for i := range menuOptions {
		column = 36
		for _, char := range menuOptions[i] {
			DisplayRune(column, line+(2*i), char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayPrincipalCursor(option, previousOption int) {

	var cursor [][]rune
	var column, line int
	column = 11

	if previousOption != 0 {

		cursor = append(cursor, []rune("      "))
		cursor = append(cursor, []rune("                   "))
		cursor = append(cursor, []rune("      "))

		switch previousOption {

		case 1:
			line = 3
		case 2:
			line = 5
		case 3:
			line = 7
		case 4:
			line = 9
		case 5:
			line = 11
		}
		for i := range cursor {
			column = 11
			for _, char := range cursor[i] {
				DisplayRune(column, line+i, char)
				column += rwidth.RuneWidth(char)
			}
		}

		previousOption = 0

		DisplayPrincipalCursor(option, previousOption)

	} else {

		switch option {

		case 1:
			line = 3
		case 2:
			line = 5
		case 3:
			line = 7
		case 4:
			line = 9
		case 5:
			line = 11
		}

		for i := range Sword {
			column = 11
			for _, char := range Sword[i] {
				DisplayRune(column, line+i, char)
				column += rwidth.RuneWidth(char)
			}
		}
	}
}

func displayCredits() {
	ClearTerminal()
	fmt.Println("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Crédits ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println(" /┃\\                                                                       /┃\\  ")
	fmt.Println("//┃\\\\                      /                              PROJET RED      //┃\\\\")
	fmt.Println("\\\\┃//                     / ITO                                           \\\\┃//")
	fmt.Println("/\\┃/\\             \\      /\\                           SEPTEMBRE 2023      /\\┃/\\")
	fmt.Println("\\/┃\\/              \\    /  \\                                              \\/┃\\/")
	fmt.Println("//┃\\\\               \\  /¯¯¯¯\\ NTOINE                                      //┃\\\\")
	fmt.Println("\\\\┃//                \\/      \\                                            \\\\┃//")
	fmt.Println("/\\┃/\\                                                                     /\\┃/\\")
	fmt.Println("\\/┃\\/                                                                     \\/┃\\/")
	fmt.Println("//┃\\\\  |¯¯\\  |¯¯\\  /¯¯\\  |¯¯\\  |    |  /¯¯\\ ¯¯¯¯|¯¯¯¯ ¯¯|¯¯ /¯¯\\  |\\   |  //┃\\\\")
	fmt.Println("\\\\┃//  |__/  |__/ /    \\ |   \\ |    | /         |       |  /    \\ | \\  |  \\\\┃//")
	fmt.Println("/\\┃/\\  |     |\\   \\    / |   / |    | \\         |       |  \\    / |  \\ |  /\\┃/\\")
	fmt.Println("\\/┃\\/ _|_    | \\   \\__/  |__/  |____|  \\__/     |     __|__ \\__/  |   \\|  \\/┃\\/")
	fmt.Println("//┃\\\\                                                                     //┃\\\\")
	fmt.Println("\\\\┃//                                                                     \\\\┃//")
	fmt.Println(" \\┃/                                                                  Ⓒ Ⓡ  \\┃/  ")
	fmt.Printf("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	time.Sleep(time.Second * 3)
}

func CharCreationNameDisplay(err bool) {
	column := 8
	for _, char := range CharCreationName {
		DisplayRune(column, 3, char)
		column += rwidth.RuneWidth(char)
	}
	if err {
		column = 8
		for _, char := range CharCreationNameError {
			DisplayRune(column, 2, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func CharCreationMenuMainDisplay(name string) {

	DisplayText(13, 5, name)

	var column int
	var lines = []int{7, 9, 10, 11, 13}

	for i := range CharCreationMenuText {

		switch i {

		case 0, 4:
			column = 8

		case 1, 2, 3:
			column = 16

		}
		for _, char := range CharCreationMenuText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func CharCreationMenuChangingDisplay(option, previousOption int) {

	column := 7
	var lines = []int{9, 10, 11}

	DisplayText(column, lines[option-1], string(CharCreationMenuCursor)) //➸

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "      ")
	}

	var description [][]rune

	switch option {

	case 1:
		description = HumanDescription

	case 2:
		description = ElfDescription

	case 3:
		description = DwarfDescription

	}
	line := 14
	column = 16
	for i := range description {
		column = 16
		for _, char := range description[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func CharMenu(char *char.Character) {
	DisplayBlankMenu(CHAR_MENU)

}

// ╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮
//                                   │ ├────────┬───────────────────────┤
// ├─────────────────┬───────────────┧ ╽
// ├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯
// ╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮
//                  │├──────────────────────┼─────┼─────┼─────────┼─────┧
//                  ╰┴──────────────────────┴─────┴─────┴─────────┴─────╯
//
// │
// ┃
// ┃
// ╿
// │
//
// │
// │
// │
// │
//
// ┃ ┃
// ┃ ┃
// ╿ ╿
// │ │
//
// │
// │
// │
// │
// │
//
// ╽
// ┃
// ┃
// ╿
// │
//
// │┃
// │┃
// │╿
// ││
// ││
//
// │
// │
// │
// │
// │
//
// │
// │
// │
// │
// │
//
// │
// │
// │
// │
// │
//
// │
// │
// │
// │
// │
//
// ┃
// ┃
// ╿
// │
// │
//
// Niveau :
// Attaque :
// Défense :
// Agilité :
//
// XP :
// HP :
// MP :
// Or :
//
// Tête
// Torse
// Mains
// Jambes
// Pieds
//
// ⎯{==-
//
// Inventaire
// Équipement
// Retour
// Quitter

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮  /┃\
//  //┃\\ │  MonPseudoEstTropLong   Humain  │ ├────────┬───────────────────────┤ //┃\\
//  \\┃// ┟─────────────────┬───────────────┧ ╽ Tête   │ Heaume du chasseur    ╽ \\┃//
//  /\┃/\ ┃ Niveau :    100 │ XP : 24.785   ┃ ┃ Torse  │ Armure de fantassin   ┃ /\┃/\
//  \/┃\/ ┃ Attaque :   150 │ HP : 350/400  ┃ ┃ Mains  │ Épée longue en mithril┃ \/┃\/
//  //┃\\ ╿ Défense :   170 │ MP : 430/600  ╿ ╿ Jambes │ Jambières elfiques    ╿ //┃\\
//  \\┃// │ Agilité :   275 │ Or : 12.570   │ │ Pieds  │ Chausses légères      │ \\┃//
//  /\┃/\ ├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯ /\┃/\
//  \/┃\/ ╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮ \/┃\/
//  //┃\\ ⎯{==- Inventaire │┟──────────────────────┼─────┼─────┼─────────┼─────┧ //┃\\
//  \\┃//      (Croissance)│┃ Coup de poing humain │ 240 │ 100 │ Agilité │ 105 ┃ \\┃//
//  /\┃/\       Équipement │┃ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 ┃ /\┃/\
//  \/┃\/                  │╿ Coup de poing humain │ 240 │ 100 │ Défense │ 105 ╿ \/┃\/
//  //┃\\       Retour     ││ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 │ //┃\\
//  \\┃//       Quitter    ││ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 │ \\┃//
//   \┃/                   ╰┴──────────────────────┴─────┴─────┴─────────┴─────╯  \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//
