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
	PrincipalMenu()
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

	case 3:
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

	//Run the beginning of the game (like the cinematic introduction and then send to the first location)
	CharMenu(MainChar)
}

func CharMenu(myChar char.Character) {
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
		CharInventory(myChar)

	case 2: //Need to implement that display (Equipment managing menu)
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

func CharInventory(myChar char.Character) {

	term.Clear(term.ColorDefault, term.ColorDefault)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	var options int
	for {
		items := displayCharInventory(myChar, &options)
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
		actionOptions := 3

		displayCharInventoryActionCursor(actionPointingAt, 0)

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
				displayCharInventoryActionCursor(actionPointingAt, previousActionPointingAt)
				previousActionPointingAt = actionPointingAt
			}
		}
		// returnToInventory:
		switch selectedAction {
		case 1:
			switch retreiveItemType(items[pointingAt-1]) {
			case "Équipement":
				//Equip the item!
			case "Potion":
				myChar.TakePotionSoin() //Change to an universal function for potions (one that checks automatically which potion it is and then, what to do with it and do it)
			case "Livre de sort":
				for _, spellBook := range char.AllSpellBook {
					if items[pointingAt] == spellBook.Name {
						myChar.SpellBook(spellBook)
					}
				}
			}
		case 2:
			//Remove the item from the inventory (all of it)
		case 3:
			selectedOption = 0
			displayCharInventoryActionCursor(0, actionPointingAt)
			// break returnToInventory
		}
	}
}
