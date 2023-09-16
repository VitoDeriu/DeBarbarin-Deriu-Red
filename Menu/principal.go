package menu

import (
	char "ProjetRed/Character"
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
			time.Sleep(time.Millisecond * 45)
		}
		if i < len(introductionStory)-1 {
			fmt.Println()
			time.Sleep(time.Millisecond * 800)
		}
	}
	time.Sleep(time.Second * 3)
	PrincipalMenu()
}

func PrincipalMenu() {

	pointingAt := 1
	var selectedOption int

	//Initializing keyboard input variables
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	principalMenu1()

	for selectedOption == 0 {

		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {

			//Arrow up
			case term.KeyArrowUp:
				if pointingAt > 1 {
					pointingAt--
				} else {
					continue
				}

			//Arrow down
			case term.KeyArrowDown:
				if pointingAt < 5 {
					pointingAt++
				} else {
					continue
				}

			//Enter
			case term.KeyEnter:
				selectedOption = pointingAt

			//Autres touches
			default:
				continue
			}
		case term.EventError:
			panic(ev.Err)
		}

		switch pointingAt {
		case 1:
			principalMenu1()
		case 2:
			principalMenu2()
		case 3:
			principalMenu3()
		case 4:
			principalMenu4()
		case 5:
			principalMenu5()
		}
	}

	switch selectedOption {
	case 1:
		fmt.Println("Nouvelle partie")
		fmt.Println("Désolé, il n'y a rien pour le moment !")

		//À debugguer avant de réactiver
		// CharacterCreationMenu()

	case 2:
		ClearTerminal()
		LoadingScreen()
		PrincipalMenu()

	case 3:
		ClearTerminal()
		// fmt.Println("Bonus")
		// fmt.Println("Désolé, il n'y a rien pour le moment !")
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
		os.Exit(0)
	}
}

func CharacterCreationMenu() {
	var name string
	var nameRegistered bool
	CharacterCreationDisplayName(false)
	for !nameRegistered {
		fmt.Scan(&name)
		fmt.Print(name)
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
	var selectedOption int

	for selectedOption == 0 {

		CharacterCreationDisplayRace(name, nameCharSpaces, pointingAt)

		//Switch case expecting keyboard input
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {

			//Arrow up
			case term.KeyArrowUp:
				if pointingAt > 1 {
					pointingAt--
				} else {
					continue
				}

			//Arrow down
			case term.KeyArrowDown:
				if pointingAt < 3 {
					pointingAt++
				} else {
					continue
				}

			//Enter
			case term.KeyEnter:
				selectedOption = pointingAt

			//Autres touches
			default:
				continue
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
	char.CreateMainCharacter(name, selectedOption).DisplayInfo()
}
