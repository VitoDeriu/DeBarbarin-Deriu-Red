package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"strconv"
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
	CharMenuMainGrid,
	CharInventoryGrid,
	CharInventoryText [][]rune

var BottomBar,
	CharMenuTitleBar,
	PrincipalMenuTitleBar,
	CharCreationMenuTitleBar,
	CharInventoryTitleBar,
	CharCreationName,
	CharCreationNameError,
	CharCreationMenuCursor,
	CharMenuCursor,
	CharInventoryItemsCursor []rune

const (
	PRINCIPAL_MENU     = 0
	CHAR_CREATION_MENU = 1
	CHAR_MENU          = 2
	CHAR_INVENTORY     = 3
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

	CharMenuText = append(CharMenuText, []rune("Niveau :"))
	CharMenuText = append(CharMenuText, []rune("Attaque :"))
	CharMenuText = append(CharMenuText, []rune("Défense :"))
	CharMenuText = append(CharMenuText, []rune("Agilité :"))
	CharMenuText = append(CharMenuText, []rune("XP :"))
	CharMenuText = append(CharMenuText, []rune("HP :"))
	CharMenuText = append(CharMenuText, []rune("MP :"))
	CharMenuText = append(CharMenuText, []rune("Or :"))
	CharMenuText = append(CharMenuText, []rune("Tête"))
	CharMenuText = append(CharMenuText, []rune("Torse"))
	CharMenuText = append(CharMenuText, []rune("Mains"))
	CharMenuText = append(CharMenuText, []rune("Jambes"))
	CharMenuText = append(CharMenuText, []rune("Pieds"))
	CharMenuText = append(CharMenuText, []rune("Inventaire"))
	CharMenuText = append(CharMenuText, []rune("Équipement"))
	CharMenuText = append(CharMenuText, []rune("Retour"))
	CharMenuText = append(CharMenuText, []rune("Quitter"))

	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│ ├────────┬───────────────────────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┬───────────────┤ │"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│├──────────────────────┼─────┼─────┼─────────┼─────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰┴──────────────────────┴─────┴─────┴─────────┴─────╯"))

	CharInventoryGrid = append(CharInventoryGrid, []rune("─────────┴──────────────────────────────────────────────────────────────"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("──────────────────┬─────────────────────────────────────────────────────"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))

	CharInventoryText = append(CharInventoryText, []rune("╳  Esc"))
	CharInventoryText = append(CharInventoryText, []rune("Nom"))
	CharInventoryText = append(CharInventoryText, []rune("Type"))
	CharInventoryText = append(CharInventoryText, []rune("Quantité"))
	CharInventoryText = append(CharInventoryText, []rune("Utiliser"))
	CharInventoryText = append(CharInventoryText, []rune("Jeter"))
	CharInventoryText = append(CharInventoryText, []rune("Annuler"))
	CharInventoryText = append(CharInventoryText, []rune("Description :"))

	BottomBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	PrincipalMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharInventoryTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Inventaire ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationName = []rune("Nom du personnage (max 20 caractères) :")

	CharCreationNameError = []rune("Erreur : veuillez rentrer un nom valide !")

	CharCreationMenuCursor = []rune("⎯{===-")

	CharMenuCursor = []rune("⎯{==-")

	CharInventoryItemsCursor = []rune("⎯{====-")
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

	case CHAR_INVENTORY:
		CurrentTopBar = CharInventoryTitleBar

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
	if menuNb == CHAR_MENU || menuNb == PRINCIPAL_MENU {
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

	DisplayText(column, lines[option-1], string(CharCreationMenuCursor))

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

func displayCharMenuGrid() {

	var columns = []int{6, 40, 6, 6, 6, 23, 23}
	var lines = []int{1, 2, 3, 8, 9, 10, 16}
	var column int

	for i := range CharMenuMainGrid {
		column = columns[i]
		for _, char := range CharMenuMainGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}

	columns = []int{6, 51, 75, 6, 6, 6, 6, 24, 24, 24, 24, 40, 40, 40, 40, 42, 42, 42, 42, 51, 51, 51, 51, 75, 75, 75, 75, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 47, 47, 47, 47, 47, 53, 53, 53, 53, 53, 59, 59, 59, 59, 59, 69, 69, 69, 69, 69, 75, 75, 75, 75, 75}
	lines = []int{2, 3, 3, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15}

	for i, column := range columns {
		DisplayRune(column, lines[i], '│')
	}
}

func displayCharMenuText() {
	var column int
	var columns = []int{8, 8, 8, 8, 26, 26, 26, 26, 44, 44, 44, 44, 44, 12, 12, 12, 12}
	var lines = []int{4, 5, 6, 7, 4, 5, 6, 7, 3, 4, 5, 6, 7, 10, 12, 14, 15}

	for i := range CharMenuText {
		column = columns[i]
		for _, char := range CharMenuText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayCharMenuCursor(option, previousOption int) {
	column := 6
	var lines = []int{10, 12, 14, 15}

	DisplayText(column, lines[option-1], string(CharMenuCursor))

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ")
	}
}

func displayCharMenuStats(myChar *char.Character) {
	var column int
	var columns = []int{9, 32, 20, 20, 20, 20, 31, 31, 31, 31, 53, 53, 53, 53, 53, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71}
	var lines = []int{2, 2, 4, 5, 6, 7, 4, 5, 6, 7, 3, 4, 5, 6, 7, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15}
	var stats []string
	stats = append(stats, myChar.Name)
	stats = append(stats, myChar.Class.Name)
	stats = append(stats, strconv.Itoa(myChar.Level))
	stats = append(stats, strconv.Itoa(myChar.Attack))
	stats = append(stats, strconv.Itoa(myChar.Defense))
	stats = append(stats, strconv.Itoa(myChar.Agility))
	stats = append(stats, strconv.Itoa(myChar.Xp))
	stats = append(stats, strconv.Itoa(myChar.Hp)+"/"+strconv.Itoa(myChar.HpMax))
	stats = append(stats, strconv.Itoa(myChar.Mp)+"/"+strconv.Itoa(myChar.MpMax))
	stats = append(stats, strconv.Itoa(myChar.Gold))
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	for _, skill := range myChar.Skills {
		stats = append(stats, skill.Name)
		stats = append(stats, strconv.Itoa(skill.Attack))
		stats = append(stats, strconv.Itoa(skill.Defense))
		stats = append(stats, skill.StatBuffed)
		stats = append(stats, strconv.Itoa(skill.Buff))
	}
	for i, singleStat := range stats {
		column = columns[i]
		for _, char := range singleStat {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharMenu(myChar *char.Character) {
	DisplayBlankMenu(CHAR_MENU)
	displayCharMenuGrid()
	displayCharMenuText()
	displayCharMenuStats(myChar)
}

func displayCharInventoryGrid() {
	var columns = []int{5, 5, 14, 23, 23, 23}
	var lines = []int{2, 13, 1, 14, 15, 16}
	var column int

	for i := range CharInventoryGrid {
		column = columns[i]
		for _, char := range CharInventoryGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharInventoryText() {
	var columns = []int{6, 24, 49, 64, 12, 12, 12, 26}
	var lines = []int{1, 1, 1, 1, 14, 15, 16, 14}
	var column int

	for i := range CharInventoryText {
		column = columns[i]
		for _, char := range CharInventoryText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharInventoryItems(myChar *char.Character, itemOptions *int) []string {

	if len(myChar.Inventory) == 0 {
		DisplayText(17, 4, "Votre inventaire est vide !")
		time.Sleep(time.Second * 2)
		return nil
	}

	var inventory []string
	for item := range myChar.Inventory {
		inventory = append(inventory, item)
	}

	*itemOptions = len(inventory)

	line := 3
	for index, item := range inventory {
		if index == 10 {
			break
		}
		DisplayText(17, line+index, item)
		DisplayText(47, line+index, retreiveItemType(item))
		DisplayText(67, line+index, strconv.Itoa(myChar.Inventory[item]))
	}
	return inventory
}

func displayCharInventory(myChar char.Character, itemOptions *int) []string {
	DisplayBlankMenu(CHAR_INVENTORY)
	displayCharInventoryGrid()
	displayCharInventoryText()
	return displayCharInventoryItems(&myChar, itemOptions)
}

func displayCharInventoryItemsCursor(option, previousOption int) {
	column := 7
	var lines = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	DisplayText(column, lines[option-1], string(CharInventoryItemsCursor))

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "       ")
	}
}

func displayCharInventoryItemDescription(item string) {

	var description string

	switch retreiveItemType(item) {
	case "Équipement":
		for _, singleItem := range char.AllEquipement {
			if item == singleItem.Name {
				description = singleItem.Name + " est un équipement.       " // Add Description field in the Equipment struct!!!
			}
		}
	case "Potion":
		for _, singleItem := range char.AllPotion {
			if item == singleItem.Name {
				description = singleItem.Name + " est une potion.       " // Add Description field in the Potion struct!!!
			}
		}
	case "Livre de sort":
		for _, singleItem := range char.AllSpellBook {
			if item == singleItem.Name {
				description = singleItem.Name + " est un livre de sort.       " // Add Description field in the SpellBook struct!!!
			}
		}
	default:
		description = "Item inconnu"
	}

	DisplayText(29, 15, description)
}

func displayCharInventoryActionCursor(option, previousOption int) {
	column := 6
	var lines = []int{14, 15, 16}
	if option != 0 {
		DisplayText(column, lines[option-1], string(CharMenuCursor))
	}

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ")
	}
}

// Rajouter tous les types d'item dans la fonction afin de prendre en compte toutes les possibilités !
func retreiveItemType(s string) string {

	for _, name := range char.AllEquipement {
		if s == name.Name {
			return "Équipement"
		}
	}
	for _, name := range char.AllPotion {
		if s == name.Name {
			return "Potion"
		}
	}
	for _, name := range char.AllSpellBook {
		if s == name.Name {
			return "Livre de sort"
		}
	}
	return "Inconnu" // Dommage, cet item n'est par répertorié dans les slices présentes dans la fonction... :´(
}

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

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Inventaire ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╳  Esc  │         Nom                      Type           Quantité      /┃\
//  //┃\\─────────┴──────────────────────────────────────────────────────────────//┃\\
//  \\┃//  ⎯{====-   Heaume du chasseur            Équipement          2         \\┃//
//  /\┃/\            Armure de fantassin           Équipement          1         /\┃/\
//  \/┃\/            Épée longue en mithril        Équipement          1         \/┃\/
//  //┃\\            Jambières elfiques            Équipement          1         //┃\\
//  \\┃//            Chausses légères              Équipement          2         \\┃//
//  /\┃/\            Potion de soin                Potion              10        /\┃/\
//  \/┃\/            Potion de poison              Potion              24        \/┃\/
//  //┃\\            Boule de feu                  Livre de sort       1         //┃\\
//  \\┃//            Peau de mammouth              Matériau            35        \\┃//
//  /\┃/\            Morceau de mithril            Matériau            7         /\┃/\
//  \/┃\/──────────────────┬─────────────────────────────────────────────────────\/┃\/
//  //┃\\       Utiliser   │  Description :                                      //┃\\
//  \\┃// ⎯{==- Jeter      │     Le heaume du chasseur donne 35 pts de défense   \\┃//
//   \┃/        Annuler    │     et 15 pts d'attaque.                             \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Château ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\   MonPseudoEstTropLong    Humain         │ HP : 350/400 │  Or : 12.570   /┃\
//  //┃\\─────────────────────────────────────────┴──────────────┴───────────────//┃\\
//  \\┃//                                                                        \\┃//
//  /\┃/\                                                                        /\┃/\
//  \/┃\/                                                                        \/┃\/
//  //┃\\                                                                        //┃\\
//  \\┃//                                                                        \\┃//
//  /\┃/\                                                                        /\┃/\
//  \/┃\/                                                                        \/┃\/
//  //┃\\                                                                        //┃\\
//  \\┃//                                                                        \\┃//
//  /\┃/\                                                                        /\┃/\
//  \/┃\/──────────────────┬─────────────────────────────────────────────────────\/┃\/
//  //┃\\       Caserne    │  Description :                                      //┃\\
//  \\┃// ⎯{==- Ville      │     Le heaume du chasseur donne 35 pts de défense   \\┃//
//   \┃/        Menu       │     et 15 pts d'attaque.                             \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
