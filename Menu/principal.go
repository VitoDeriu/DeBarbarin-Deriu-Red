package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unicode"

	rwidth "github.com/mattn/go-runewidth"
	term "github.com/nsf/termbox-go"
)

var introductionStory [][]rune
var MainChar char.Character
var GameStatus int

func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PresentationCinematic() {
	ClearTerminal()
	messageIntro()
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
			// DisplayBlankMenu(PRINCIPAL_MENU)
			// DisplayMenuOptions(PRINCIPAL_MENU)
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
		PrincipalMenu()

	case 3: // AFFICHER LE BONUS : Spielberg et ABBA ;)
		ClearTerminal()
		DisplayBlankMenu(PRINCIPAL_MENU)
		DisplayMenuOptions(PRINCIPAL_MENU)
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 4:
		displayCredits()
		ClearTerminal()
		PrincipalMenu()

	case 5:
		ClearTerminal()
		fmt.Println("Bye bye !")
		time.Sleep(time.Second * 3)
		ClearTerminal()
		os.Exit(0)
	}
}

func DisplayRune(column, line int, char rune) {
	term.SetCell(column, line, char, term.ColorDefault, term.ColorDefault)
	term.Flush()
}

func DisplayText(column, line int, text string) {
	for _, char := range text {
		term.SetCell(column, line, char, term.ColorDefault, term.ColorDefault)
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
				DisplayRune(column, line, ev.Ch)
				name = append(name, ev.Ch)
				column += rwidth.RuneWidth(ev.Ch)
			}

			switch ev.Key {

			case term.KeyEnter:
				enterKeyIsPressed = true

			case term.KeyBackspace:
				if column > 13 {
					column -= rwidth.RuneWidth(name[len(name)-1])
					DisplayRune(column, line, ' ')
					if len(name) < 2 {
						name = []rune("")
					} else {
						name = append(name[:len(name)-2], name[len(name)-2])
					}
				}

			case term.KeySpace:
				DisplayRune(column, line, ' ')
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

	//Run the beginning of the game (like the cinematic introduction and then send to the first location)
	// PresentationCinematic()
	Stroll(&MainChar, GameStatus)
}

func CharMenu(myChar *char.Character) {
	for {
		pointingAt := 1
		selectedOption := 0
		previousPointingAt := 1
		options := 4

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
			Inventory(myChar, false)

		case 2: //Need to implement that display (Equipment managing menu)!!
			PrincipalMenu()

		case 3: //return to where you called the function
			return

		case 4: //Exiting the program... bye!
			ClearTerminal()
			fmt.Println("Bye bye !")
			time.Sleep(time.Second * 3)
			ClearTerminal()
			os.Exit(0)
		}
	}
}

func Inventory(myChar *char.Character, isMerchant bool) {

	term.Clear(term.ColorDefault, term.ColorDefault)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	var options int
	for {
		ClearTerminal()
		term.Clear(term.ColorDefault, term.ColorDefault)
		items := displayInventory(*myChar, &options, isMerchant)
		if pointingAt > len(items) {
			pointingAt = len(items) - 1
		}
		if len(items) == 0 {
			time.Sleep(time.Second * 1)
			ClearTerminal()
			term.Clear(term.ColorDefault, term.ColorDefault)
			return
		}
		displayCharInventoryItemDescription(items[pointingAt-1])
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
				displayCharInventoryItemDescription(items[pointingAt-1])
				previousPointingAt = pointingAt
			}
		}

		actionPointingAt := 1
		selectedAction := 0
		previousActionPointingAt := 1
		var actionOptions int

		if !isMerchant {
			actionOptions = 3
		} else {
			actionOptions = 2
		}

		displayCharInventoryActionCursor(actionPointingAt, 0, isMerchant)

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
				displayCharInventoryActionCursor(actionPointingAt, previousActionPointingAt, isMerchant)
				previousActionPointingAt = actionPointingAt
			}
		}

		if !isMerchant {
			switch selectedAction {

			//Use the item!
			case 1:
				switch retreiveItemType(items[pointingAt-1]) {

				case "Équipement":
					//Equip the item!

				case "Potion":
					for _, potion := range char.AllPotion {
						if items[pointingAt-1] == potion.Name {
							myChar.TakePotion(potion)
							time.Sleep(time.Second * 2)
							ClearTerminal()
							term.Clear(term.ColorDefault, term.ColorDefault)
							return
						}
					}

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if items[pointingAt-1] == spellBook.Name {
							myChar.SpellBook(spellBook)
							term.Clear(term.ColorDefault, term.ColorDefault)
							return
						}
					}
				}

			//Throw away the item!
			case 2:

				switch retreiveItemType(items[pointingAt-1]) {

				case "Équipement":
					for _, equipement := range char.AllEquipement {
						if items[pointingAt-1] == equipement.Name {
							myChar.RemoveEquipement(equipement)
						}
					}

				case "Potion":
					for _, potion := range char.AllPotion {
						if items[pointingAt-1] == potion.Name {
							myChar.RemovePotion(potion)
						}
					}

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if items[pointingAt-1] == spellBook.Name {
							myChar.RemoveSpellBook(spellBook)
						}
					}
				}

			//Return to item selection
			case 3:
				selectedOption = 0
				displayCharInventoryActionCursor(0, actionPointingAt, isMerchant)
			}

		} else {
			//If it is the Merchant
			switch selectedAction {

			//Buy the item
			case 1:
				switch retreiveItemType(items[pointingAt-1]) {

				case "Équipement":
					for _, equipement := range char.AllEquipement {
						if items[pointingAt-1] == equipement.Name {
							MainChar.BuyEquipement(equipement)
						}
					}

				case "Potion":
					for _, potion := range char.AllPotion {
						if items[pointingAt-1] == potion.Name {
							MainChar.BuyPotion(potion)
						}
					}

				case "Livre de sort":
					for _, spellBook := range char.AllSpellBook {
						if items[pointingAt-1] == spellBook.Name {
							MainChar.BuySpellBook(spellBook)
						}
					}

				case "Ressource":
					for _, ressource := range char.AllRessources {
						if items[pointingAt-1] == ressource.Name {
							MainChar.BuyRessource(ressource)
						}
					}
				}

			//return to item selection
			case 2:
				selectedOption = 0
				displayCharInventoryActionCursor(0, actionPointingAt, isMerchant)
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
					CharMenu(myChar)
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
				// Tournoi ou entrainement
			case 2:
				nbMenu = STROLL_CASTLE
			case 3:
				// Tournoi ou entrainement
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
				nbMenu = STROLL_ARENA
			case 2:
				nbMenu = STROLL_CITY
			case 3:
				nbMenu = STROLL_MISSIONS
			}
		case STROLL_MARKET:
			switch selectedOption {

			case 1:
				nbMenu = STROLL_MERCHANT
			case 2:
				nbMenu = STROLL_CITY
			case 3:
				// Go to the blacksmith
			}
		case STROLL_MERCHANT:
			switch selectedOption {

			case 1:
				Inventory(&char.Merchant, true)
			case 2:
				nbMenu = STROLL_MARKET
			case 3:
				// Sell items from the merchant
			}
		}

	}

}
