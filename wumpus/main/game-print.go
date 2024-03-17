package main

import "fmt"

func print_game(field [][]Cell, ptr_massage *string) {
	print_field(field)
	fmt.Println(*ptr_massage)
	*ptr_massage = ""
}

func print_field(field [][]Cell) {

	for j := 0; j <= len(field[0])*2; j++ {
		for i := 0; i <= len(field)*2; i++ {
			if i == 0 && j == 0 {
				fmt.Print("  ┌─")
			} else if i == 0 && j == len(field[0])*2 {
				fmt.Print("  └─")
			} else if i == len(field)*2 && j == 0 {
				fmt.Print("─┐")
			} else if i == len(field)*2 && j == len(field[0])*2 {
				fmt.Print("─┘")
			} else if i == 0 && j%2 == 0 {
				fmt.Print("  ├─")
			} else if i%2 == 0 && j == 0 {
				fmt.Print("─┬─")
			} else if i == len(field)*2 && j%2 == 0 {
				fmt.Print("─┤")
			} else if i%2 == 0 && j == len(field[0])*2 {
				fmt.Print("─┴─")
			} else if i%2 == 0 && j%2 == 0 {
				fmt.Print("─┼─")
			} else if i == 0 {
				fmt.Print("  │ ")
			} else if j == 0 {
				fmt.Print("─")
			} else if i == len(field)*2 {
				fmt.Print(" │")
			} else if j == len(field[0])*2 {
				fmt.Print("─")
			} else if i%2 == 0 {
				fmt.Print(" │ ")
			} else if j%2 == 0 {
				fmt.Print("─")
			} else if i%2 == 1 && j%2 == 1 {
				print_cell(field[(i-1)/2][(j-1)/2])
			}
		}
		fmt.Print("\n")
	}
}

func print_cell(cell Cell) {
	var cell_contents string

	if cell.visited {
		cell_contents = fmt.Sprintf("%s%s%c%s", COLOR_BOLD, cell.color, cell.symbol, COLOR_RESET)
	} else {
		cell_contents = " "
	}

	fmt.Print(cell_contents)
}
