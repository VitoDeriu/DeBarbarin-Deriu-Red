package menu

import (
	char "ProjetRed/Character"
	"fmt"
	rwidth "github.com/mattn/go-runewidth"
	"strconv"
	"time"
)

const (
	BONUS_MENU         = -1
	PRINCIPAL_MENU     = 0
	CHAR_CREATION_MENU = 1
	CHAR_MENU          = 2
	CHAR_INVENTORY     = 3
	STROLL_CASTLE      = 4
	STROLL_CITY        = 5
	STROLL_BARRACKS    = 6
	STROLL_OUTSIDE     = 7
	STROLL_MARKET      = 8
	STROLL_ARENA       = 9
	STROLL_MERCHANT    = 10
	STROLL_BLACKSMITH  = 11
	STROLL_MISSIONS    = 12
	BUY_MERCHANT       = 13
	SELL_MERCHANT      = 14
	COMBAT             = 15
	TOURNAMENT         = 16
	TRAINING           = 17
)

// Loading bar when you run the game
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

// Exit display
func Bye() {
	for i := range ByeDisplay {
		column := 10
		for _, char := range ByeDisplay[i] {
			DisplayRune(column, i+1, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the title of the corresponding menu (call it when building the display of the menu)
func displayTopBar(menuNb int) {
	var CurrentTopBar []rune

	switch menuNb {

	case PRINCIPAL_MENU:
		CurrentTopBar = PrincipalMenuTitleBar

	case BONUS_MENU:
		CurrentTopBar = BonusTitleBar

	case CHAR_CREATION_MENU:
		CurrentTopBar = CharCreationMenuTitleBar

	case CHAR_MENU:
		CurrentTopBar = CharMenuTitleBar

	case CHAR_INVENTORY:
		CurrentTopBar = CharInventoryTitleBar

	case STROLL_CASTLE:
		CurrentTopBar = StrollCastleTitleBar

	case STROLL_BARRACKS:
		CurrentTopBar = StrollBarracksTitleBar

	case STROLL_CITY:
		CurrentTopBar = StrollCityTitleBar

	case STROLL_MARKET:
		CurrentTopBar = StrollMarketTitleBar

	case STROLL_OUTSIDE:
		CurrentTopBar = StrollOutsideTitleBar

	case STROLL_MERCHANT, BUY_MERCHANT, SELL_MERCHANT:
		CurrentTopBar = StrollMerchantTitleBar

	case STROLL_BLACKSMITH:
		CurrentTopBar = BlacksmithMenuTitleBar

	case STROLL_ARENA:
		CurrentTopBar = StrollArenaTitleBar

	case STROLL_MISSIONS:
		//CurrentTopBar = StrollMissionsTitleBar

	case COMBAT, TOURNAMENT, TRAINING:
		CurrentTopBar = CombatTitleBar
	}
	column := 0
	for _, char := range CurrentTopBar {
		DisplayRune(column, 0, char, CYAN)
		column += rwidth.RuneWidth(char)
	}
}

// Display the lateral decorations of the menus (call when building a menu)
func displayMenuBars() {
	for i := range MenuLateralBar {
		column := 0
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
	for i := range MenuLateralBar {
		column := 77
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the bottom bar (call it when building a menu)
func displayBottomBar() {
	column := 0
	for _, char := range BottomBar {
		DisplayRune(column, 17, char, CYAN)
		column += rwidth.RuneWidth(char)
	}
}

// Display the best logo ever ;)
func displayLogo() {
	column := 69
	line := 13
	for i := range LogoLittle {
		column = 69
		for _, char := range LogoLittle[i] {
			DisplayRune(column, line+i, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display an ASCII Art soldier (currently used when creating the character)
func displaySoldier() {
	column := 57
	line := 3
	for i := range SoldierArt {
		column = 57
		for _, char := range SoldierArt[i] {
			DisplayRune(column, line+i, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Build the display of a blank menu (call it whith whichever constant int variable corresponding to a menu)
func DisplayBlankMenu(menuNb int) {
	displayTopBar(menuNb)
	displayMenuBars()
	displayBottomBar()
	if menuNb == CHAR_CREATION_MENU || menuNb == PRINCIPAL_MENU {
		displayLogo()
	}
	if menuNb == CHAR_CREATION_MENU {
		displaySoldier()
	}
}

// Display the possible options of the menu (currently only used for the principal menu...)
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
			DisplayRune(column, line+(2*i), char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the cursor of the principal menu (and erase the previous appearance of it)
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
				DisplayRune(column, line+i, char, CYAN)
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
				DisplayRune(column, line+i, char, CYAN)
				column += rwidth.RuneWidth(char)
			}
		}
	}
}

// Display of the character creation menu (when writing the custom name)
func CharCreationNameDisplay(err bool) {
	column := 8
	for _, char := range CharCreationName {
		DisplayRune(column, 3, char, CYAN)
		column += rwidth.RuneWidth(char)
	}
	if err {
		column = 8
		for _, char := range CharCreationNameError {
			DisplayRune(column, 2, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the complete character creation menu (with the name already typed)
func CharCreationMenuMainDisplay(name string) {

	DisplayText(13, 5, name, CYAN)

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
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the character creation menu changing stuff (cursor and race description)
func CharCreationMenuChangingDisplay(option, previousOption int) {

	column := 7
	var lines = []int{9, 10, 11}

	DisplayText(column, lines[option-1], string(CharCreationMenuCursor), CYAN)

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "      ", CYAN)
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
			DisplayRune(column, line+i, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the empty grid of the character menu
func displayCharMenuGrid() {

	var columns = []int{6, 40, 6, 6, 6, 23, 23}
	var lines = []int{1, 2, 3, 8, 9, 10, 16}
	var column int

	for i := range CharMenuMainGrid {
		column = columns[i]
		for _, char := range CharMenuMainGrid[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}

	columns = []int{6, 51, 75, 6, 6, 6, 6, 24, 24, 24, 24, 40, 40, 40, 40, 42, 42, 42, 42, 51, 51, 51, 51, 75, 75, 75, 75, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 47, 47, 47, 47, 47, 53, 53, 53, 53, 53, 59, 59, 59, 59, 59, 69, 69, 69, 69, 69, 75, 75, 75, 75, 75}
	lines = []int{2, 3, 3, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15}

	for i, column := range columns {
		DisplayRune(column, lines[i], '│', CYAN)
	}
}

// Display the unchanging text of the character menu
func displayCharMenuText() {
	var column int
	var columns = []int{8, 8, 8, 8, 26, 26, 26, 26, 44, 44, 44, 44, 44, 12, 12, 12}
	var lines = []int{4, 5, 6, 7, 4, 5, 6, 7, 3, 4, 5, 6, 7, 10, 13, 15}

	for i := range CharMenuText {
		column = columns[i]
		for _, char := range CharMenuText[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the cursor of the character menu (and erase the previous one)
func DisplayCharMenuCursor(option, previousOption int) {
	column := 6
	var lines = []int{10, 13, 15}

	DisplayText(column, lines[option-1], string(CharMenuCursor), CYAN)

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ", CYAN)
	}
}

// Display the changing text of the character menu (stats and other stuff related to the character)
func displayCharMenuStats(myChar *char.Character) {
	myEquipment := make([]string, 5)
	for _, equipment := range myChar.Equipement {
		switch equipment.Slot {
		case "head":
			myEquipment[0] = equipment.Name
		case "chest":
			myEquipment[1] = equipment.Name
		case "arms":
			myEquipment[2] = equipment.Name
		case "legs":
			myEquipment[3] = equipment.Name
		case "feet":
			myEquipment[4] = equipment.Name
		}
	}
	var column int
	var columns = []int{9, 32, 20, 20, 20, 20, 31, 31, 31, 31, 52, 52, 52, 52, 52, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71}
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
	stats = append(stats, myEquipment[0])
	stats = append(stats, myEquipment[1])
	stats = append(stats, myEquipment[2])
	stats = append(stats, myEquipment[3])
	stats = append(stats, myEquipment[4])
	for _, skill := range myChar.Skills {
		stats = append(stats, skill.Name)
		stats = append(stats, strconv.Itoa(skill.Attack))
		stats = append(stats, strconv.Itoa(skill.Defense))
		stats = append(stats, skill.Kind)
		stats = append(stats, strconv.Itoa(skill.MpCost))
	}
	for i, singleStat := range stats {
		column = columns[i]
		for _, char := range singleStat {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Build the display of the character menu
func displayCharMenu(myChar *char.Character) {
	DisplayBlankMenu(CHAR_MENU)
	displayCharMenuGrid()
	displayCharMenuText()
	displayCharMenuStats(myChar)
}

// Display the empty grid of the inventory
func displayCharInventoryGrid() {
	var columns = []int{5, 5, 14, 23, 23, 23}
	var lines = []int{2, 13, 1, 14, 15, 16}
	var column int

	for i := range CharInventoryGrid {
		column = columns[i]
		for _, char := range CharInventoryGrid[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the unchanging text of the inventory
func displayCharInventoryText(whichMenu int, blacksmithFacility bool) {

	var columns []int
	var lines []int
	var column int
	var inventoryText [][]rune

	switch whichMenu {
	case CHAR_INVENTORY:
		if blacksmithFacility {
			inventoryText = CharInventoryTextBlacksmithFacility
		} else {
			inventoryText = CharInventoryText
		}
		columns = []int{6, 24, 49, 64, 12, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 16, 14}
	case BUY_MERCHANT:
		inventoryText = MerchantInventoryText
		columns = []int{6, 24, 49, 64, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 14}
	case SELL_MERCHANT:
		inventoryText = SellMerchantInventoryText
		columns = []int{6, 24, 49, 64, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 14}
	}

	for i := range inventoryText {
		column = columns[i]
		for _, char := range inventoryText[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the character related stuff in the inventory (items)
func displayInventoryItems(myChar *char.Character, whichMenu, currentPage int) ([]string, int) {

	for line := 3; line < 13; line++ {
		DisplayText(17, line, "                                                          ", CYAN)
	}

	if len(myChar.Inventory) == 0 {
		DisplayText(17, 4, "Votre inventaire est vide !", CYAN)
		time.Sleep(time.Second * 2)
		return nil, 0
	}

	var inventory []string
	for item := range myChar.Inventory {
		inventory = append(inventory, item)
	}

	var nbPages int
	var currentItems []string
	if len(inventory) > 10 {
		if len(inventory)%10 > 0 {
			nbPages++
		}
		nbPages += len(inventory) / 10
		if currentPage < nbPages {
			currentItems = inventory[(currentPage*10)-10 : currentPage*10]
		} else {
			currentItems = inventory[(currentPage*10)-10:]
		}
	} else {
		currentPage = 1
		currentItems = inventory
	}

	line := 3
	for index, item := range currentItems {
		if index == 10 {
			break
		}
		DisplayText(17, line+index, item, CYAN)
		if whichMenu == CHAR_INVENTORY {
			DisplayText(47, line+index, retreiveItemType(item), CYAN)
		} else {
			switch retreiveItemType(item) {
			case "Potion":
				for _, potion := range char.AllPotion {
					if item == potion.Name {
						DisplayText(49, line+index, strconv.Itoa(potion.Price), CYAN)
					}
				}
			case "Livre de sort":
				for _, spellBook := range char.AllSpellBook {
					if item == spellBook.Name {
						DisplayText(49, line+index, strconv.Itoa(spellBook.Price), CYAN)
					}
				}
			case "Équipement":
				for _, equipement := range char.AllEquipement {
					if item == equipement.Name {
						DisplayText(49, line+index, strconv.Itoa(equipement.Price), CYAN)
					}
				}
			case "Ressource":
				for _, ressource := range char.AllRessources {
					if item == ressource.Name {
						DisplayText(49, line+index, strconv.Itoa(ressource.Price), CYAN)
					}
				}
			default:
				if item == char.EnhanceInventory.Name {
					DisplayText(49, line+index, strconv.Itoa(char.EnhanceInventory.Price), CYAN)
				}
			}

		}
		DisplayText(67, line+index, strconv.Itoa(myChar.Inventory[item]), CYAN)
	}
	return currentItems, nbPages
}

// Build the inventory menu
func displayInventory(myChar char.Character, whichMenu, currentPage int, blacksmithFacility bool) ([]string, int) {
	DisplayBlankMenu(whichMenu)
	displayCharInventoryGrid()
	displayCharInventoryText(whichMenu, blacksmithFacility)
	return displayInventoryItems(&myChar, whichMenu, currentPage)
}

// Display the cursor for the inventory
func displayCharInventoryItemsCursor(option, previousOption int) {
	column := 7
	var lines = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	DisplayText(column, lines[option-1], string(CharInventoryItemsCursor), CYAN)

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "       ", CYAN)
	}
}

// Display the item's description in the inventory (called each time the cursor changes item)
func displayItemDescription(item string) {

	DisplayText(29, 15, "                                                ", CYAN)
	DisplayText(29, 16, "                                                ", CYAN)

	var description string

	switch retreiveItemType(item) {
	case "Équipement":
		for _, singleItem := range char.AllEquipement {
			if item == singleItem.Name {
				description = singleItem.Name + " est un équipement.     " // Add Description field in the Equipment struct!!!
			}
		}
	case "Potion":
		for _, singleItem := range char.AllPotion {
			if item == singleItem.Name {
				description = singleItem.Name + " est une potion.          " // Add Description field in the Potion struct!!!
			}
		}
	case "Livre de sort":
		for _, singleItem := range char.AllSpellBook {
			if item == singleItem.Name {
				description = singleItem.Name + " est un livre de sort. " // Add Description field in the SpellBook struct!!!
			}
		}
	case "Ressource":
		for _, singleItem := range char.AllRessources {
			if item == singleItem.Name {
				description = singleItem.Name + " est une ressource.      " // Add Description field in the Ressource struct!!!
			}
		}
	case "Trésor":
		switch item {
		case char.WorldTreeLeaf.Name:
			description = item + " est un trésor elfe. À utiliser.  "
		case char.MountainHeart.Name:
			description = item + " est un trésor nain. À utiliser.  "
		case char.GraalFragment.Name:
			description = item + " est un trésor humain. À utiliser.  "
		}
	default:
		if item == char.EnhanceInventory.Name {
			description = "Augmenter la capacité de l'inventaire.  "
		}
		description = "Item inconnu"
	}

	DisplayText(29, 15, description, CYAN)
}

// Display the second cursor in the inventory (to select an action to do)
func displayCharInventoryActionCursor(option, previousOption, whichMenu int) {
	column := 6
	var lines []int
	if whichMenu == CHAR_INVENTORY {
		lines = []int{14, 15, 16}
	} else {
		lines = []int{14, 15}
	}
	if option != 0 {
		DisplayText(column, lines[option-1], string(CharMenuCursor), CYAN)
	}

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ", CYAN)
	}
}

// Useful function to retreive the item's name type (currently used for the inventory, merchant and blacksmith)
func retreiveItemType(s string) string {

	// Rajouter tous les types d'item dans la fonction afin de prendre en compte toutes les possibilités !
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
	for _, name := range char.AllRessources {
		if s == name.Name {
			return "Ressource"
		}
	}
	switch s {
	case char.WorldTreeLeaf.Name, char.MountainHeart.Name, char.GraalFragment.Name:
		return "Trésor"
	}
	return "Inconnu" // Dommage, cet item n'est par répertorié dans les slices présentes dans la fonction... :´(
}

// Display the empty grid of the stroll menu
func displayStrollGrid() {
	column := 5
	var lines = []int{2, 8, 10, 11, 15}
	for i := range StrollGrid {
		column = 5
		for _, char := range StrollGrid[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Fill the stroll menu's grid
func displayStrollText(myChar *char.Character, nbMenu int) {

	var text = []string{myChar.Name, myChar.Class.Name, "│ HP :", strconv.Itoa(myChar.Hp) + "/" + strconv.Itoa(myChar.HpMax), "│  Or :", strconv.Itoa(myChar.Gold), "│ Space :  Menu   "}

	switch nbMenu {
	case STROLL_CASTLE:
		text = append(text, "│  Caserne   │")
		text = append(text, "│   Ville    │")
		text = append(text, "│  Château  │")
	case STROLL_BARRACKS:
		text = append(text, "│Entrainement│")
		text = append(text, "│  Tournoi   │")
		text = append(text, "│  Château  │")
	case STROLL_CITY:
		text = append(text, "│   Dehors   │")
		text = append(text, "│   Marché   │")
		text = append(text, "│  Château  │")
	case STROLL_OUTSIDE:
		text = append(text, "│   Arène    │")
		text = append(text, "│  Mission   │")
		text = append(text, "│   Ville   │")
	case STROLL_MARKET:
		text = append(text, "│  Marchand  │")
		text = append(text, "│  Forgeron  │")
		text = append(text, "│   Ville   │")
	case STROLL_MERCHANT:
		text = append(text, "│  Acheter   │")
		text = append(text, "│   Vendre   │")
		text = append(text, "│  Marché   │")
	}

	var columns = []int{7, 31, 46, 53, 61, 69, 60, 11, 58, 34}
	var lines = []int{1, 1, 1, 1, 1, 1, 16, 9, 9, 16}
	var column int

	for i, str := range text {
		column = columns[i]
		for _, char := range str {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the "cursor" of the stroll menu (in a guy's shape)
func displayStrollCursor(pointingAt, previousPointingAt int) {

	var columns = []int{16, 39, 63}
	var column int
	line := 12

	for i := range StrollCursor {
		column = columns[pointingAt-1]
		for _, char := range StrollCursor[i] {
			DisplayRune(column, line+i, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}

	if previousPointingAt != 0 {
		for i := 0; i < 3; i++ {
			DisplayText(columns[previousPointingAt-1], line+i, "   ", CYAN)
		}
	}
}

// Build the stroll menu
func displayStrollMenu(myChar *char.Character, nbMenu int) {
	DisplayBlankMenu(nbMenu)
	displayStrollGrid()
	displayStrollText(myChar, nbMenu)
}

// Display the empty Blacksmith menu
func displayBlacksmithBasicMenu() {
	var columns = []int{5, 5, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 5, 12, 12, 23}
	var column int
	line := 1

	for i := range BlacksmithMenuText {
		column = columns[i]
		for _, char := range BlacksmithMenuText[i] {
			DisplayRune(column, line+i, char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display the Blacksmith's equipment
func displayBlacksmithEquipment() {
	column := 17
	line := 3

	for i, equipment := range char.BlacksmithEquipments {
		DisplayText(column, line+i, equipment.Name, CYAN)
	}
}

// Display the changing recipe of the equipment currently pointed at by the cursor
func displayBlacksmithRecipe(pointingAt int) {
	column := 50
	line := 4

	for i := 0; i < 9; i++ {
		DisplayText(column, line+i, "                           ", CYAN)
	}
	DisplayText(column, line, "10   Pièces d'or", CYAN)
	line += 2
	i := 0
	for ressource, nb := range char.BlacksmithEquipments[pointingAt-1].Recipe {
		DisplayText(column, line+i, strconv.Itoa(nb), CYAN)
		DisplayText(column+5, line+i, ressource.Name, CYAN)
		i += 2
	}
}

// Build the Blacksmith menu
func displayBlacksmithMenu() {
	DisplayBlankMenu(STROLL_BLACKSMITH)
	displayBlacksmithBasicMenu()
	displayBlacksmithEquipment()
}

// Display the empty grid of the combat menu
func displayCombatGrid() {
	columns := []int{5, 5, 5, 5, 23, 5, 23, 5, 23, 5, 23, 5, 23, 5, 23, 5, 12, 12, 12, 12, 12}
	lines := []int{2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 12, 13, 14, 15, 16}
	var column int

	for i := range CombatGrid {
		column = columns[i]
		for _, char := range CombatGrid[i] {
			DisplayRune(column, lines[i], char, CYAN)
			column += rwidth.RuneWidth(char)
		}
	}
}

// Display all stats in the combat menu
func displayCombatStats(myChar *char.Character, otherChar *char.Enemy) {
	var StatsToDisplay []string
	StatsToDisplay = append(StatsToDisplay, otherChar.Name)
	StatsToDisplay = append(StatsToDisplay, otherChar.Race)
	StatsToDisplay = append(StatsToDisplay, "Niveau :")
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(otherChar.Level))
	StatsToDisplay = append(StatsToDisplay, "HP :")
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(otherChar.Hp)+"/"+strconv.Itoa(otherChar.HpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Level))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Hp)+"/"+strconv.Itoa(myChar.HpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Mp)+"/"+strconv.Itoa(myChar.MpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Attack))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Defense))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Agility))

	columns := []int{6, 28, 45, 53, 62, 67, 18, 15, 15, 18, 18, 18}
	lines := []int{1, 1, 1, 1, 1, 1, 5, 6, 7, 8, 9, 10}
	var column int

	for i, str := range StatsToDisplay {
		column = columns[i]
		DisplayText(column, lines[i], str, CYAN)
	}
}

// Display the character's skills in the combat menu when the cursor is pointing at "Attaque"
func displayCombatSkills(myChar *char.Character) {
	line := 12
	for i, skill := range myChar.Skills {
		DisplayText(31, line+i, skill.Name, CYAN)
		DisplayText(57, line+i, skill.Kind, CYAN)
		DisplayText(68, line+i, strconv.Itoa(skill.MpCost), CYAN)
		DisplayText(73, line+i, "MP", CYAN)
	}
}

// Display the character's potions in the combat menu when the cursor is pointing at "Potions"
func displayCombatPotions(myChar *char.Character) []string {
	line := 12
	var myPotions []string
	for item := range myChar.Inventory {
		if retreiveItemType(item) == "Potion" {
			myPotions = append(myPotions, item)
		}
	}
	for i, potion := range myPotions {
		currentPotion := char.FindPotion(potion)
		DisplayText(31, line+i, currentPotion.Name, CYAN)
		if currentPotion.Buff != 0 {
			DisplayText(49, line+i, "+"+strconv.Itoa(currentPotion.Buff), CYAN)
			DisplayText(53, line+i, currentPotion.StatBuffed, CYAN)
		}
		if currentPotion.Debuff != 0 {
			DisplayText(62, line+i, "-"+strconv.Itoa(currentPotion.Debuff), CYAN)
			DisplayText(66, line+i, currentPotion.StatDebuffed, CYAN)
		}
		DisplayText(74, line+i, strconv.Itoa(myChar.Inventory[potion]), CYAN)
	}
	return myPotions
}

// Clear the skill or potions square's display
func clearDisplayCombatOptions() {
	for i := 12; i < 17; i++ {
		DisplayText(24, i, "                                                     ", CYAN)
	}
}

// Display the combat menu's first cursor and erase the previous one (when choosing a kind of action)
func displayCombatActionTypeCursor(option, previousOption int) {
	column := 6
	lines := []int{12, 13, 14, 16}

	if option != 0 {
		DisplayText(column, lines[option-1], string(CharMenuCursor), CYAN)
	}

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ", CYAN)
	}
}

// Display the combat menu's second cursor and erase the previous one (when choosing a specific action)
func displayCombatSpecificCursor(option, previousOption int) {
	column := 25
	line := 12
	if option != 0 {
		DisplayText(column, line+option-1, string(CharMenuCursor), CYAN)
	}

	if previousOption != 0 {
		DisplayText(column, line+previousOption-1, "     ", CYAN)
	}
}

// Build the combat menu
func DisplayCombatMenu(myChar *char.Character, enemy *char.Enemy, combatType int) {
	DisplayBlankMenu(combatType)
	displayCombatGrid()
	displayCombatStats(myChar, enemy)
}
