package main

import (
	"fmt"
)

func main() {

	field, player := generate_field(0, 0)
	print_game(field, "gggg")
	fmt.Println(field[0][0], player)
}
