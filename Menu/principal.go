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

func TextInput() string {

	var enterKeyIsPressed bool
	var name []rune
	writingLineBefore := "\\/┃\\/    "
	writingLineAfter := "(_/`-`\\_)           \\/┃\\/"
	column := 13
	line := 5

	for !enterKeyIsPressed {

		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {

		case term.EventKey:
			switch ev.Ch {

			case 'q':
				DisplayRune(column, line, 'q')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'q')
				column += rwidth.RuneWidth('q')

			case 'w':
				DisplayRune(column, line, 'w')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'w')
				column += rwidth.RuneWidth('w')

			case 'e':
				DisplayRune(column, line, 'e')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'e')
				column += rwidth.RuneWidth('e')

			case 'r':
				DisplayRune(column, line, 'r')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'r')
				column += rwidth.RuneWidth('r')

			case 't':
				DisplayRune(column, line, 't')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 't')
				column += rwidth.RuneWidth('t')

			case 'y':
				DisplayRune(column, line, 'y')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'y')
				column += rwidth.RuneWidth('y')

			case 'u':
				DisplayRune(column, line, 'u')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'u')
				column += rwidth.RuneWidth('u')

			case 'i':
				DisplayRune(column, line, 'i')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'i')
				column += rwidth.RuneWidth('i')

			case 'o':
				DisplayRune(column, line, 'o')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'o')
				column += rwidth.RuneWidth('o')

			case 'p':
				DisplayRune(column, line, 'p')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'p')
				column += rwidth.RuneWidth('p')

			case 'a':
				DisplayRune(column, line, 'a')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'a')
				column += rwidth.RuneWidth('a')

			case 's':
				DisplayRune(column, line, 's')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 's')
				column += rwidth.RuneWidth('s')

			case 'd':
				DisplayRune(column, line, 'd')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'd')
				column += rwidth.RuneWidth('d')

			case 'f':
				DisplayRune(column, line, 'f')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'f')
				column += rwidth.RuneWidth('f')

			case 'g':
				DisplayRune(column, line, 'g')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'g')
				column += rwidth.RuneWidth('g')

			case 'h':
				DisplayRune(column, line, 'h')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'h')
				column += rwidth.RuneWidth('h')

			case 'j':
				DisplayRune(column, line, 'j')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'j')
				column += rwidth.RuneWidth('j')

			case 'k':
				DisplayRune(column, line, 'k')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'k')
				column += rwidth.RuneWidth('k')

			case 'l':
				DisplayRune(column, line, 'l')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'l')
				column += rwidth.RuneWidth('l')

			case 'z':
				DisplayRune(column, line, 'z')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'z')
				column += rwidth.RuneWidth('z')

			case 'x':
				DisplayRune(column, line, 'x')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'x')
				column += rwidth.RuneWidth('x')

			case 'c':
				DisplayRune(column, line, 'c')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'c')
				column += rwidth.RuneWidth('c')

			case 'v':
				DisplayRune(column, line, 'v')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'v')
				column += rwidth.RuneWidth('v')

			case 'b':
				DisplayRune(column, line, 'b')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'b')
				column += rwidth.RuneWidth('b')

			case 'n':
				DisplayRune(column, line, 'n')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'n')
				column += rwidth.RuneWidth('n')

			case 'm':
				DisplayRune(column, line, 'm')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, 'm')
				column += rwidth.RuneWidth('m')

			case '0':
				DisplayRune(column, line, '0')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '0')
				column += rwidth.RuneWidth('0')

			case '1':
				DisplayRune(column, line, '1')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '1')
				column += rwidth.RuneWidth('1')

			case '2':
				DisplayRune(column, line, '2')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '2')
				column += rwidth.RuneWidth('2')

			case '3':
				DisplayRune(column, line, '3')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '3')
				column += rwidth.RuneWidth('3')

			case '4':
				DisplayRune(column, line, '4')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '4')
				column += rwidth.RuneWidth('4')

			case '5':
				DisplayRune(column, line, '5')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '5')
				column += rwidth.RuneWidth('5')

			case '6':
				DisplayRune(column, line, '6')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '6')
				column += rwidth.RuneWidth('6')

			case '7':
				DisplayRune(column, line, '7')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '7')
				column += rwidth.RuneWidth('7')

			case '8':
				DisplayRune(column, line, '8')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '8')
				column += rwidth.RuneWidth('8')

			case '9':
				DisplayRune(column, line, '9')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
				name = append(name, '9')
				column += rwidth.RuneWidth('9')

			}

			switch ev.Key {

			case term.KeyEnter:
				enterKeyIsPressed = true

			case term.KeyBackspace:
				if column > 13 {
					column -= rwidth.RuneWidth(name[len(name)-1])
					DisplayRune(column, line, ' ')
					DisplayText(0, line, writingLineBefore)
					DisplayText(57, line, writingLineAfter)
					if len(name) < 2 {
						name = []rune("")
					} else {
						name = append(name[:len(name)-2], name[len(name)-2])
					}
				}

			case term.KeySpace:
				DisplayRune(column, line, ' ')
				DisplayText(0, line, writingLineBefore)
				DisplayText(57, line, writingLineAfter)
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

	// nameCharSpaces := char.Spaces(name, 17)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 1
	options := 3

	CharCreationMenuMainDisplay(name)
	CharCreationMenuChangingDisplay(pointingAt, 0)
	// CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)

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
			// CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)
		}
	}

	term.Clear(term.ColorDefault, term.ColorDefault)
	MainChar = char.CreateMainCharacter(name, selectedOption)

	pointingAt = 1
	selectedOption = 0
	previousPointingAt = 1
	options = 4

	CharMenu(&MainChar)
	DisplayCharMenuCursor(pointingAt, 0)
	// CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)

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
			// CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)
		}
	}

	switch selectedOption {

	case 1:
		PrincipalMenu()

	case 2:
		PrincipalMenu()

	case 3:
		PrincipalMenu()

	case 4:
		ClearTerminal()
		fmt.Println("Bye bye !")
		time.Sleep(time.Second * 3)
		ClearTerminal()
		os.Exit(0)
	}
}
