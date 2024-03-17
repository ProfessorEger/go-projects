package main

import (
	"github.com/eiannone/keyboard"
)

func StartReadInput() (err error) {
	err = keyboard.Open()
	return
}

func StopReadInput() {
	keyboard.Close()
}

func GetAction(canShot bool, QUICK_INPUT bool) (action Action) {

	if QUICK_INPUT {
		action = quickGetAction(canShot)
	} else {
		action = standardGetAction(canShot)
	}
	return
}
func quickGetAction(canShot bool) (action Action) {
	action.ActionType = MOVE

WaitForInput:
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'w', 'W':
			action.Direction = UP
			break WaitForInput
		case 'd', 'D':
			action.Direction = RIGHT
			break WaitForInput
		case 's', 'S':
			action.Direction = DOWN
			break WaitForInput
		case 'a', 'A':
			action.Direction = LEFT
			break WaitForInput
		}

		if canShot {
			switch key {
			case keyboard.KeyCtrlW:
				action.Direction = UP
				action.ActionType = SHOT
				break WaitForInput
			case keyboard.KeyCtrlD:
				action.Direction = RIGHT
				action.ActionType = SHOT
				break WaitForInput
			case keyboard.KeyCtrlS:
				action.Direction = DOWN
				action.ActionType = SHOT
				break WaitForInput
			case keyboard.KeyCtrlA:
				action.Direction = LEFT
				action.ActionType = SHOT
				break WaitForInput
			}
		}
	}

	return
}

func standardGetAction(canShot bool) (action Action) {
	return
}
