package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
	"unicode"

	rwidth "github.com/mattn/go-runewidth"
	term "github.com/nsf/termbox-go"
)

var introductionStory [][]rune
var MainChar char.Character
var GameStatus int
var hourMerchantVisit, minutesMerchantVisit int

// Constantes pour spécifier les couleurs d'affichage et constante indiquant la fréquence d'AddItems du marchand
const (
	CYAN              = 0
	GREEN             = 1
	RED               = 2
	MERCHANT_INTERVAL = 2
)

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Run the game! It is the only function to call in "func main()" to play the game !
func Run() {
	LoadingScreen()
	PrincipalMenu()
}

// Display the introduction message. WARNING: call it after the character creation menu because it takes the name's user into account.
func PresentationCinematic(name string) {
	ClearTerminal()
	messageIntro(name)
	fmt.Println()
	for i := 0; i < len(introductionStory); i++ {
		for j := 1; j <= len(introductionStory[i]); j++ {
			fmt.Print("\r", string(introductionStory[i][:j]))
			time.Sleep(time.Millisecond * 40)
		}
		if i < len(introductionStory)-1 {
			fmt.Println()
			time.Sleep(time.Millisecond * 500)
		}
	}
	time.Sleep(time.Second * 3)
	ClearTerminal()
}

func PrincipalMenu() {

	pointingAt := 1
	previousPointingAt := 1
	selectedOption := 0
	options := 5

	//Initializing keyboard input variables --Use only once!--
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	term.Clear(term.ColorDefault, term.ColorDefault)
	DisplayBlankMenu(PRINCIPAL_MENU)
	DisplayMenuOptions(PRINCIPAL_MENU)
	DisplayPrincipalCursor(pointingAt, 0)

	for selectedOption == 0 {
		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {

			//Arrow up
			case term.KeyArrowUp:
				if pointingAt > 1 {
					pointingAt--
				}

			//Arrow down
			case term.KeyArrowDown:
				if pointingAt < options {
					pointingAt++
				}

			//Enter
			case term.KeyEnter:
				selectedOption = pointingAt

			}
		case term.EventError:
			panic(ev.Err)
		}

		//Change the display only if necessary.
		if previousPointingAt != pointingAt {
			DisplayPrincipalCursor(pointingAt, previousPointingAt)
			previousPointingAt = pointingAt
		}
	}

	term.Clear(term.ColorDefault, term.ColorDefault)

	switch selectedOption {
	case 1:
		ClearTerminal()
		CharacterCreationMenu()

	case 2:
		ClearTerminal()
		LoadingScreen()
		ClearTerminal()
		PrincipalMenu()

	case 3: // AFFICHER LE BONUS : Spielberg et ABBA ;)
		ClearTerminal()
		DisplayBlankMenu(BONUS_MENU)
		DisplayText(23, 5, "La réponse à la question bonus est :", CYAN)
		DisplayText(30, 10, "Spielberg,  ABBA et Queen", CYAN)
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 4:
		displayCredits()
		ClearTerminal()
		PrincipalMenu()

	case 5:
		ClearTerminal()
		term.Clear(term.ColorDefault, term.ColorDefault)
		Bye()
		time.Sleep(time.Second * 3)
		ClearTerminal()
		os.Exit(0)
	}
}

func DisplayRune(column, line int, char rune, colorNb int) {
	var color term.Attribute
	switch colorNb {
	case CYAN:
		color = term.ColorCyan
	case GREEN:
		color = term.ColorGreen
	case RED:
		color = term.ColorRed
	}
	term.SetCell(column, line, char, color, term.ColorDefault)
	term.Flush()
}

func DisplayText(column, line int, text string, colorNb int) {
	var color term.Attribute
	switch colorNb {
	case CYAN:
		color = term.ColorCyan
	case GREEN:
		color = term.ColorGreen
	case RED:
		color = term.ColorRed
	}
	for _, char := range text {
		term.SetCell(column, line, char, color, term.ColorDefault)
		column += rwidth.RuneWidth(char)
	}
	term.Flush()
}

func isAlpha(r rune) bool {
	if r < 48 || (r > 57 && r < 65) || (r > 90 && r < 97) || r > 122 {
		return false
	}
	return true
}

func TextInput() string {

	var enterKeyIsPressed bool
	var name []rune
	column := 13
	line := 5

	for !enterKeyIsPressed {

		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {

		case term.EventKey:
			if isAlpha(ev.Ch) {
				DisplayRune(column, line, ev.Ch, CYAN)
				name = append(name, ev.Ch)
				column += rwidth.RuneWidth(ev.Ch)
			}

			switch ev.Key {

			case term.KeyEnter:
				enterKeyIsPressed = true

			case term.KeyBackspace:
				if column > 13 {
					column -= rwidth.RuneWidth(name[len(name)-1])
					DisplayRune(column, line, ' ', CYAN)
					if len(name) < 2 {
						name = []rune("")
					} else {
						name = append(name[:len(name)-2], name[len(name)-2])
					}
				}

			case term.KeySpace:
				DisplayRune(column, line, ' ', CYAN)
				name = append(name, ' ')
				column += rwidth.RuneWidth(' ')

			}

		case term.EventError:
			panic(ev.Err)
		}
	}
	return string(name)
}

func CharacterCreationMenu() {

	var name string
	var nameRegistered bool

	DisplayBlankMenu(CHAR_CREATION_MENU)
	CharCreationNameDisplay(false)

	for !nameRegistered {

		name = TextInput()

		if len([]rune(name)) > 20 || len([]rune(name)) < 2 {
			term.Clear(term.ColorDefault, term.ColorDefault)
			DisplayBlankMenu(CHAR_CREATION_MENU)
			CharCreationNameDisplay(true)
			name = ""
		} else {
			chars := []rune(name)
			chars[0] = unicode.ToUpper(chars[0])
			name = string(chars)
			nameRegistered = true
		}
	}

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	options := 3

	CharCreationMenuMainDisplay(name)
	CharCreationMenuChangingDisplay(pointingAt, 0)

	for selectedOption == 0 {

		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {

			//Arrow up
			case term.KeyArrowUp:
				if pointingAt > 1 {
					pointingAt--
				}

			//Arrow down
			case term.KeyArrowDown:
				if pointingAt < options {
					pointingAt++
				}

			//Enter
			case term.KeyEnter:
				selectedOption = pointingAt
			}
		case term.EventError:
			panic(ev.Err)
		}

		//Change the display only if necessary.
		if previousPointingAt != pointingAt {
			CharCreationMenuChangingDisplay(pointingAt, previousPointingAt)
			previousPointingAt = pointingAt
		}
	}

	term.Clear(term.ColorDefault, term.ColorDefault)
	MainChar = char.CreateMainCharacter(name, selectedOption)
	GameStatus = STROLL_CASTLE
	hourMerchantVisit = time.Now().Hour()
	minutesMerchantVisit = time.Now().Minute()

	//Run the beginning of the game (like the cinematic introduction and then send to the first location)
	PresentationCinematic(MainChar.Name)
	Stroll(&MainChar, GameStatus)
}

func CharMenu(myChar *char.Character, nbMenu int) {
	for {
		pointingAt := 1
		selectedOption := 0
		previousPointingAt := 1
		options := 3

		displayCharMenu(&MainChar)
		DisplayCharMenuCursor(pointingAt, 0)

		for selectedOption == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if pointingAt > 1 {
						pointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if pointingAt < options {
						pointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedOption = pointingAt
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousPointingAt != pointingAt {
				DisplayCharMenuCursor(pointingAt, previousPointingAt)
				previousPointingAt = pointingAt
			}
		}

		switch selectedOption {

		case 1: //Display the inventory
			var blacksmithFacility bool
			if nbMenu == STROLL_MARKET {
				blacksmithFacility = true
			}
			Inventory(myChar, CHAR_INVENTORY, blacksmithFacility)

		case 2: //return to where you called the function
			return

		case 3: //Exiting the program... bye!
			ClearTerminal()
			term.Clear(term.ColorDefault, term.ColorDefault)
			Bye()
			time.Sleep(time.Second * 3)
			ClearTerminal()
			os.Exit(0)
		}
	}
}

func Inventory(myChar *char.Character, whichMenu int, blacksmithFacility bool) {

	term.Clear(term.ColorDefault, term.ColorDefault)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	currentPage := 1
	previousPage := 1
	var currentItems []string
	var options, nbPages int
	for {
		ClearTerminal()
		term.Clear(term.ColorDefault, term.ColorDefault)
		currentItems, nbPages = displayInventory(*myChar, whichMenu, currentPage, blacksmithFacility)
		options = len(currentItems)

		if len(currentItems) == 0 {
			time.Sleep(time.Second * 1)
			ClearTerminal()
			term.Clear(term.ColorDefault, term.ColorDefault)
			return

		}

		if pointingAt > len(currentItems) {
			pointingAt = len(currentItems)
		}

		displayItemDescription(currentItems[pointingAt-1])
		displayCharInventoryItemsCursor(pointingAt, 0)

		for selectedOption == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if pointingAt > 1 {
						pointingAt--
					} else if currentPage > 1 {
						currentPage--
					}

				//Arrow down
				case term.KeyArrowDown:
					if pointingAt < options {
						pointingAt++
					} else if currentPage < nbPages {
						currentPage++
					}

				//Enter
				case term.KeyEnter:
					selectedOption = pointingAt

				//Escape
				case term.KeyEsc:
					term.Clear(term.ColorDefault, term.ColorDefault)
					return
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousPointingAt != pointingAt {
				displayItemDescription(currentItems[pointingAt-1])
				displayCharInventoryItemsCursor(pointingAt, previousPointingAt)
				previousPointingAt = pointingAt
			}
			if previousPage != currentPage {
				currentItems, nbPages = displayInventory(*myChar, whichMenu, currentPage, blacksmithFacility)
				options = len(currentItems)
				previousPage = currentPage
				if pointingAt > len(currentItems) {
					pointingAt = len(currentItems)
					displayCharInventoryItemsCursor(pointingAt, previousPointingAt)
					previousPointingAt = pointingAt
				}
			}
		}

		actionPointingAt := 1
		selectedAction := 0
		previousActionPointingAt := 1
		var actionOptions int

		if whichMenu == CHAR_INVENTORY {
			actionOptions = 3
		} else {
			actionOptions = 2
		}

		displayCharInventoryActionCursor(actionPointingAt, 0, whichMenu)

		for selectedAction == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if actionPointingAt > 1 {
						actionPointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if actionPointingAt < actionOptions {
						actionPointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedAction = actionPointingAt

				//Escape
				case term.KeyEsc:
					return
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousActionPointingAt != actionPointingAt {
				displayCharInventoryActionCursor(actionPointingAt, previousActionPointingAt, whichMenu)
				previousActionPointingAt = actionPointingAt
			}
		}

		if whichMenu == CHAR_INVENTORY {
			switch selectedAction {

			//Use the item!
			case 1:
				switch retreiveItemType(currentItems[pointingAt-1]) {

				case "Équipement":
					for _, equipment := range char.AllEquipement {
						if equipment.Name == currentItems[pointingAt-1] {
							myChar.Equiper(equipment)
							term.Clear(term.ColorDefault, term.ColorDefault)
							return
						}
					}

				case "Potion":
					for _, potion := range char.AllPotion {
						if currentItems[pointingAt-1] == potion.Name {
							message := myChar.TakePotion(potion)
							DisplayText(26, 14, "                                                   ", CYAN)
							DisplayText(26, 15, "                                                   ", CYAN)
							DisplayText(26, 16, "                                                   ", CYAN)
							DisplayText(30, 15, message, GREEN)
							time.Sleep(time.Second * 2)
							ClearTerminal()
							term.Clear(term.ColorDefault, term.ColorDefault)
							return
						}
					}

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if currentItems[pointingAt-1] == spellBook.Name {
							myChar.SpellBook(spellBook)
							term.Clear(term.ColorDefault, term.ColorDefault)
							return
						}
					}

					//Using the special treasure item of the race!
				case "Trésor":
					switch currentItems[pointingAt-1] {
					case "Feuille d'Yggdrasil":
						myChar.RemoveRessource(char.WorldTreeLeaf)
						myChar.UseSpecialItem()
					case "Coeur de montagne":
						myChar.RemoveRessource(char.MountainHeart)
						myChar.UseSpecialItem()
					case "Fragment du Graal":
						myChar.RemoveRessource(char.GraalFragment)
						myChar.UseSpecialItem()
					}
					term.Clear(term.ColorDefault, term.ColorDefault)
					return
				}

			//Throw away the item!
			case 2:

				switch retreiveItemType(currentItems[pointingAt-1]) {

				case "Équipement":
					for _, equipement := range char.AllEquipement {
						if currentItems[pointingAt-1] == equipement.Name {
							myChar.RemoveEquipement(equipement)

							//Sneaky possibility to recycle equipment (you need to visit the Blacksmith often enough to find it ;))
							if blacksmithFacility && equipement.Recipe != nil {
								var recycledRessource char.Ressource
								var quantity int
								for ressource := range equipement.Recipe {
									if recycledRessource.Price < ressource.Price {
										recycledRessource = ressource
										quantity = equipement.Recipe[ressource]
									}
								}
								for i := 0; i < quantity; i++ {
									myChar.AddRessource(recycledRessource)
								}
							}

						}
					}
					selectedOption = 0
					displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)

				case "Potion":
					for _, potion := range char.AllPotion {
						if currentItems[pointingAt-1] == potion.Name {
							myChar.RemovePotion(potion)
						}
					}

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if currentItems[pointingAt-1] == spellBook.Name {
							myChar.RemoveSpellBook(spellBook)
						}
					}

				case "Ressource":
					for _, ressource := range char.AllRessources {
						if currentItems[pointingAt-1] == ressource.Name {
							myChar.RemoveRessource(ressource)
						}
					}
				}

			//Return to item selection
			case 3:
				selectedOption = 0
				displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)
			}

		} else {
			//If it is the Merchant
			switch selectedAction {

			//Buy the item
			case 1:
				switch retreiveItemType(currentItems[pointingAt-1]) {

				case "Équipement":
					for _, equipement := range char.AllEquipement {
						if currentItems[pointingAt-1] == equipement.Name {
							if whichMenu == BUY_MERCHANT {
								MainChar.BuyEquipement(&char.Merchant, equipement)
							} else {
								char.Merchant.BuyEquipement(&MainChar, equipement)
							}
						}
					}
					selectedOption = 0
					displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)

				case "Potion":
					if currentItems[pointingAt-1] == char.EnhanceInventory.Name {
						myChar.UpgradeInventory()
					}
					for _, potion := range char.AllPotion {
						if currentItems[pointingAt-1] == potion.Name {
							if whichMenu == BUY_MERCHANT {
								MainChar.BuyPotion(&char.Merchant, potion)
							} else {
								char.Merchant.BuyPotion(&MainChar, potion)
							}
						}
					}
					selectedOption = 0
					displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if currentItems[pointingAt-1] == spellBook.Name {
							if whichMenu == BUY_MERCHANT {
								MainChar.BuySpellBook(&char.Merchant, spellBook)
							} else {
								char.Merchant.BuySpellBook(&MainChar, spellBook)
							}
						}
					}
					selectedOption = 0
					displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)

				case "Ressource":
					for _, ressource := range char.AllRessources {
						if currentItems[pointingAt-1] == ressource.Name {
							if whichMenu == BUY_MERCHANT {
								MainChar.BuyRessource(&char.Merchant, ressource)
							} else {
								char.Merchant.BuyRessource(&MainChar, ressource)
							}
						}
					}
				}
				selectedOption = 0
				displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)

			//return to item selection
			case 2:
				selectedOption = 0
				displayCharInventoryActionCursor(0, actionPointingAt, whichMenu)
			}
		}

	}
}

func Stroll(myChar *char.Character, nbMenu int) {
	term.Clear(term.ColorDefault, term.ColorDefault)

	var options,
		pointingAt,
		selectedOption,
		previousPointingAt int

	for {
		options = 3
		pointingAt = 2
		selectedOption = 0
		previousPointingAt = 2

		term.Clear(term.ColorDefault, term.ColorDefault)
		displayStrollMenu(myChar, nbMenu)
		displayStrollCursor(pointingAt, 0)
	fromCharMenu:
		for selectedOption == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow left
				case term.KeyArrowLeft:
					if pointingAt > 1 {
						pointingAt--
					}

				//Arrow right
				case term.KeyArrowRight:
					if pointingAt < options {
						pointingAt++
					}

				//Space
				case term.KeySpace:
					selectedOption = 0
					term.Clear(term.ColorDefault, term.ColorDefault)
					CharMenu(myChar, nbMenu)
					term.Clear(term.ColorDefault, term.ColorDefault)
					displayStrollMenu(myChar, nbMenu)
					displayStrollCursor(pointingAt, 0)
					break fromCharMenu

				//Enter
				case term.KeyEnter:
					selectedOption = pointingAt

				//Escape
				case term.KeyEsc:
					term.Clear(term.ColorDefault, term.ColorDefault)
					return
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousPointingAt != pointingAt {
				displayStrollCursor(pointingAt, previousPointingAt)
				previousPointingAt = pointingAt
			}
		}

		switch nbMenu {

		case STROLL_CASTLE:
			switch selectedOption {

			case 1:
				nbMenu = STROLL_BARRACKS
			case 2:
				nbMenu = STROLL_CASTLE
			case 3:
				nbMenu = STROLL_CITY
			}
		case STROLL_BARRACKS:
			switch selectedOption {

			case 1:
				Combat(myChar, TRAINING)
			case 2:
				nbMenu = STROLL_CASTLE
			case 3:
				Combat(myChar, TOURNAMENT)
			}
		case STROLL_CITY:
			switch selectedOption {

			case 1:
				nbMenu = STROLL_OUTSIDE
			case 2:
				nbMenu = STROLL_CASTLE
			case 3:
				nbMenu = STROLL_MARKET
			}
		case STROLL_OUTSIDE:
			switch selectedOption {

			case 1:
				Combat(myChar, STROLL_ARENA)
			case 2:
				nbMenu = STROLL_CITY
			case 3:
				Combat(myChar, STROLL_MISSIONS)
			}
		case STROLL_MARKET:
			switch selectedOption {

			case 1:
				nbMenu = STROLL_MERCHANT
			case 2:
				nbMenu = STROLL_CITY
			case 3:
				BlacksmithMenu(myChar)
			}
		case STROLL_MERCHANT:
			switch selectedOption {

			case 1:
				merchantDeliveryTime()
				Inventory(&char.Merchant, BUY_MERCHANT, false)
			case 2:
				nbMenu = STROLL_MARKET
			case 3:
				merchantDeliveryTime()
				Inventory(&MainChar, SELL_MERCHANT, false)
			}
		}

	}

}

func BlacksmithMenu(myChar *char.Character) {
	term.Clear(term.ColorDefault, term.ColorDefault)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	options := len(char.BlacksmithEquipments)
	for {
		ClearTerminal()
		term.Clear(term.ColorDefault, term.ColorDefault)
		displayBlacksmithMenu()
		displayItemDescription(char.BlacksmithEquipments[pointingAt-1].Name)
		displayBlacksmithRecipe(pointingAt)
		displayCharInventoryItemsCursor(pointingAt, 0)

		for selectedOption == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if pointingAt > 1 {
						pointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if pointingAt < options {
						pointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedOption = pointingAt

				//Escape
				case term.KeyEsc:
					term.Clear(term.ColorDefault, term.ColorDefault)
					return
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousPointingAt != pointingAt {
				displayCharInventoryItemsCursor(pointingAt, previousPointingAt)
				displayBlacksmithRecipe(pointingAt)
				displayItemDescription(char.BlacksmithEquipments[pointingAt-1].Name)
				previousPointingAt = pointingAt
			}
		}

		actionPointingAt := 1
		selectedAction := 0
		previousActionPointingAt := 1
		actionOptions := 2

		displayCharInventoryActionCursor(actionPointingAt, 0, STROLL_BLACKSMITH)

		for selectedAction == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if actionPointingAt > 1 {
						actionPointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if actionPointingAt < actionOptions {
						actionPointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedAction = actionPointingAt

				//Escape
				case term.KeyEsc:
					return
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousActionPointingAt != actionPointingAt {
				displayCharInventoryActionCursor(actionPointingAt, previousActionPointingAt, STROLL_BLACKSMITH)
				previousActionPointingAt = actionPointingAt
			}
		}

		switch selectedAction {

		//Craft the item!
		case 1:
			switch char.BlacksmithEquipments[pointingAt-1].Name {
			case "Chapeau de l'aventurier":
				myChar.CraftChapeauDelAventurier()

			case "Tunique de l'aventurier":
				myChar.CraftTuniqueDelAventurier()

			case "Bottes de l'aventurier":
				myChar.CraftBotteDelAventurier()
			}
			selectedOption = 0
			displayCharInventoryActionCursor(0, actionPointingAt, STROLL_BLACKSMITH)

		//Return to item selection
		case 2:
			selectedOption = 0
			displayCharInventoryActionCursor(0, actionPointingAt, STROLL_BLACKSMITH)
		}

	}
}

func Combat(myChar *char.Character, combatType int) {
	var enemy char.Enemy
	switch combatType {
	case TRAINING:
		if myChar.Level < len(char.Trainers) {
			enemy = char.Trainers[myChar.Level-1]
		} else {
			enemy = char.Trainers[2]
		}

	case TOURNAMENT:
		enemy = char.EnemiesTournament[rand.Intn(len(char.EnemiesTournament))]

	case STROLL_ARENA:
		enemy = char.EnemiesArena[rand.Intn(len(char.EnemiesArena))]

	case STROLL_MISSIONS:
		enemy = char.EnemiesMission[rand.Intn(len(char.EnemiesMission))]
	}
	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	options := 4
	var myCharPotions []string
	var myCharBuffDefense, IABuffDefense int
	var turn, round int
	const (
		MYCHARTURN = 1
		IATURN     = 2
	)

	if myChar.Agility > enemy.Agility {
		turn = MYCHARTURN
	} else {
		turn = IATURN
	}

	term.Clear(term.ColorDefault, term.ColorDefault)
	DisplayCombatMenu(myChar, &enemy, combatType)
	displayCombatActionTypeCursor(pointingAt, 0)
	time.Sleep(time.Second / 2)
	for {
		round++

		if turn == IATURN {
			IABuffDefense = IACombatTurn(myChar, &enemy, myCharBuffDefense, round%3 == rand.Intn(100)%3)
			if isDead, xpLost := myChar.IsDead(enemy); isDead {
				DisplayText(7, 3, "Vous avez succombé à vos blessures!", RED)
				if xpLost > 0 {
					DisplayText(25, 5, "-"+strconv.Itoa(xpLost)+"XP", RED)
				}
				time.Sleep(time.Second * 1)
				return
			}
			if myCharBuffDefense != 0 {
				DisplayText(25, 9, "     ", RED)
			}
			myCharBuffDefense = 0
			turn = MYCHARTURN
			term.Clear(term.ColorDefault, term.ColorDefault)
			DisplayCombatMenu(myChar, &enemy, combatType)
			displayCombatActionTypeCursor(pointingAt, 0)
			time.Sleep(time.Second / 2)
			continue
		}

		switch pointingAt {
		case 1:
			displayCombatSkills(myChar)
		case 3:
			myCharPotions = displayCombatPotions(myChar)
		default:
			clearDisplayCombatOptions()
		}

		for selectedOption == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if pointingAt > 1 {
						pointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if pointingAt < options {
						pointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedOption = pointingAt

				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousPointingAt != pointingAt {
				displayCombatActionTypeCursor(pointingAt, previousPointingAt)
				switch pointingAt {
				case 1:
					displayCombatSkills(myChar)
				case 3:
					myCharPotions = displayCombatPotions(myChar)
				default:
					clearDisplayCombatOptions()
				}
				previousPointingAt = pointingAt
			}
		}

		actionPointingAt := 1
		selectedAction := 0
		previousActionPointingAt := 1
		var actionOptions int

		switch selectedOption {
		//Attack!
		case 1:
			actionOptions = len(myChar.Skills)

		//Absolute defense!
		case 2:
			myCharBuffDefense = myChar.Defense / 2
			DisplayText(25, 9, "+"+strconv.Itoa(myCharBuffDefense), GREEN)
			selectedAction = 1

		//Boire une potion
		case 3:
			actionOptions = len(myCharPotions)

		//Fuite!!
		case 4:
			if myChar.Hp <= enemy.Attack {
				DisplayText(25, 6, "-"+strconv.Itoa(myChar.Hp-1), RED)
				time.Sleep(time.Second * 1)
				DisplayText(25, 6, "     ", RED)
				myChar.Hp = 1
			} else {
				myChar.Hp -= enemy.Attack
				DisplayText(25, 6, "-"+strconv.Itoa(enemy.Attack), RED)
				time.Sleep(time.Second * 1)
				DisplayText(25, 6, "     ", RED)
			}
			return

		}

		displayCombatSpecificCursor(actionPointingAt, 0)

		for selectedAction == 0 {

			//Switch case expecting keyboard input
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {

				//Arrow up
				case term.KeyArrowUp:
					if actionPointingAt > 1 {
						actionPointingAt--
					}

				//Arrow down
				case term.KeyArrowDown:
					if actionPointingAt < actionOptions {
						actionPointingAt++
					}

				//Enter
				case term.KeyEnter:
					selectedAction = actionPointingAt

				//Escape
				case term.KeyEsc:
					selectedOption = 2
					selectedAction = 1
				}
			case term.EventError:
				panic(ev.Err)
			}

			//Change the display only if necessary.
			if previousActionPointingAt != actionPointingAt {
				displayCombatSpecificCursor(actionPointingAt, previousActionPointingAt)
				previousActionPointingAt = actionPointingAt
			}
		}

		switch selectedOption {

		//Use the skill!
		case 1:
			var hasAttacked bool
			if myChar.Mp >= myChar.Skills[selectedAction-1].MpCost {
				var damage int
				hasAttacked, myCharBuffDefense, damage = myChar.UseSkill(myChar.Skills[selectedAction-1], &enemy, IABuffDefense)
				DisplayText(67, 3, "-"+strconv.Itoa(damage), RED)
				if myCharBuffDefense != 0 {
					DisplayText(25, 9, "+"+strconv.Itoa(myCharBuffDefense), GREEN)
				}
				if myChar.Skills[selectedAction-1].MpCost == 0 && myChar.Mp < myChar.MpMax-myChar.MpMax/15 {
					myChar.Mp += myChar.MpMax / 15
					DisplayText(25, 7, "+"+strconv.Itoa(myChar.MpMax/15), GREEN)
				}
				time.Sleep(time.Second * 1)
				DisplayText(67, 4, "     ", GREEN)
				DisplayText(67, 3, "     ", RED)
				if victory, levelUp := enemy.EnemyDeath(myChar); victory {
					if levelUp {
						DisplayText(8, 3, "Vous avez gagné et augmenté de niveau!", GREEN)
					} else {
						DisplayText(8, 3, "Vous avez gagné!", GREEN)
					}
					time.Sleep(time.Second * 1)
					return
				}
			}
			if hasAttacked {
				selectedOption = 0
				IABuffDefense = 0
				turn = IATURN
			}

		//Return to the first menu
		case 2:
			selectedOption = 0
			turn = IATURN

		//Drink a potion!
		case 3:
			myChar.TakePotion(char.FindPotion(myCharPotions[selectedAction-1]))
			selectedOption = 0
			turn = IATURN
		}

	}
}

func IACombatTurn(myChar *char.Character, enemy *char.Enemy, buffDefense int, crit bool) int {
	var skill char.Skill
	var isSelected bool
	var outOfOptions int
	for !isSelected {
		if len(enemy.Skills) > 1 {
			skill = enemy.Skills[rand.Intn(len(enemy.Skills))]
			if enemy.Mp >= skill.MpCost {
				isSelected = true
			}
		} else {
			skill = enemy.Skills[0]
		}
		outOfOptions++
		if outOfOptions == 6 {
			skill = char.HumanPunch
			isSelected = true
		}
	}
	buff, damage := enemy.UseSkill(skill, myChar, buffDefense, crit)
	DisplayText(25, 6, "-"+strconv.Itoa(damage), RED)
	time.Sleep(time.Second * 1)
	DisplayText(25, 6, "     ", RED)
	return buff
}

func merchantDeliveryTime() {
	hourDiff := time.Now().Hour() - hourMerchantVisit
	currentMinutes := time.Now().Minute()
	if hourDiff < 0 {
		hourDiff = 24 + hourDiff
	}
	currentMinutes += 60 * hourDiff
	for i := (currentMinutes - minutesMerchantVisit) / MERCHANT_INTERVAL; i > 0; i-- {
		char.Merchant.AddItems()
	}
	hourMerchantVisit = time.Now().Hour()
	minutesMerchantVisit = time.Now().Minute()
}
