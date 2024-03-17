package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func start_read_input() (err error) {
	err = keyboard.Open()
	return
}

func stop_read_input() {
	keyboard.Close()
}

func get_action() (action Action) {
	action.action_type = MOVE

	/*var input string

	fmt.Scanln(&input)

	switch input {
	case "w", "W":
		action.direction = UP
	case "d", "D":
		action.direction = RIGHT
	case "s", "S":
		action.direction = DOWN
	case "a", "A":
		action.direction = LEFT
	}*/

	/*err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()*/

	char, _, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	switch char {
	case 'w', 'W':
		fmt.Println("Нажата клавиша 'w'.")
		action.direction = UP
	case 'd', 'D':
		action.direction = RIGHT
	case 's', 'S':
		action.direction = DOWN
	case 'a', 'A':
		action.direction = LEFT
	}

	return
}
