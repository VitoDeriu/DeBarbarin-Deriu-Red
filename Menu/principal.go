package menu

import (
	char "ProjetRed/Character"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	term "github.com/nsf/termbox-go"
)

var introductionStory [][]rune

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

	principalMenuDisplay(pointingAt)

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
			previousPointingAt = pointingAt
			principalMenuDisplay(pointingAt)
		}
	}

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
		asciiArtTesting()
		time.Sleep(time.Second * 3)
		PrincipalMenu()

	case 4:
		displayCredits()
		PrincipalMenu()

	case 5:
		ClearTerminal()
		fmt.Println("Bye bye !")
		time.Sleep(time.Second * 3)
		ClearTerminal()
		os.Exit(0)
	}
}

func CharacterCreationMenu() {

	var name string
	var nameRegistered bool

	CharacterCreationDisplayName(false)

	for !nameRegistered {

		//To debug, it doesn't show the keyboard entry but does take it into account.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name = scanner.Text()

		// fmt.Scan(&name)
		// fmt.Print(name)

		isAlphaNumeric := true

		for _, char := range name {
			if char < 48 || (char > 57 && char < 65) || (char > 90 && char < 97) || char > 122 {
				isAlphaNumeric = false
			}
		}
		if len([]rune(name)) > 20 || len([]rune(name)) < 2 || !isAlphaNumeric {
			CharacterCreationDisplayName(true)
		} else {
			nameRegistered = true
		}
	}

	nameCharSpaces := char.Spaces(name, 17)

	pointingAt := 1
	selectedOption := 0
	previousPointingAt := 0
	options := 3

	CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)

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
			previousPointingAt = pointingAt
			CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)
		}
	}
	char.CreateMainCharacter(name, selectedOption).DisplayInfo()
	PrincipalMenu()
}
