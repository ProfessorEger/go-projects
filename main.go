package main

import "wumpus/configreader"

func main() {
	configPtr, err := configreader.LoadConstantsFromFile("config.json")
	if err != nil {
		panic(err)
	}
	config = *configPtr

	err = StartReadInput()
	if err != nil {
		panic(err)
	}
	defer StopReadInput()

	startField()
}

func startField() (stopGame bool) {
	var message string
	field, player := generateField(0, 0)
	checkPosition(field, &player, &message)
	calculateFeeling(field, player.position, &message)
	printGame(field, &message)

	for {
		var action Action
		action, stopGame = GetAction(player.armed > 0, config.QuickInput)
		if stopGame {
			ClearConsole()
			return
		}
		stopGame = calculateAction(field, &player, action, &message)

		if stopGame {
			MakeVisible(field)
			break
		}

		printGame(field, &message)
	}
	printGame(field, &message)

	if config.LastPrintDelay {
		WaitInput(config.QuickInput)
	}
	return
}
