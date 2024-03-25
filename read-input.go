package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func StartReadInput() (err error) {
	err = keyboard.Open()
	return
}

func StopReadInput() {
	keyboard.Close()
}

func WaitInput(quickInput bool) {
	if quickInput {
		keyboard.GetKey()
	} else {
		fmt.Scanln()
	}
}

func GetAction(canShot bool, quickInput bool) (action Action, stopGame bool) {

	if quickInput {
		action, stopGame = quickGetAction(canShot)
	} else {
		action, stopGame = standardGetAction(canShot)
	}
	return
}
func quickGetAction(canShot bool) (action Action, stopGame bool) {
	action.ActionType = MOVE

WaitForInput:
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyCtrlQ || key == keyboard.KeyCtrlC {
			stopGame = true
			break WaitForInput
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

func standardGetAction(canShot bool) (action Action, stopGame bool) {
	return
}
