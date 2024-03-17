package main

import "fmt"

func printGame(field [][]Cell, ptrMessage *string) {
	printField(field)
	fmt.Println(*ptrMessage)
	*ptrMessage = ""
}

func printField(field [][]Cell) {

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
				printCell(field[(i-1)/2][(j-1)/2])
			}
		}
		fmt.Print("\n")
	}
}

func printCell(cell Cell) {
	var cellContents string

	if cell.visible {
		cellContents = fmt.Sprintf("%s%s%s%s", config.ColorBold, cell.color, cell.symbol, config.ColorReset)
	} else {
		cellContents = " "
	}

	fmt.Print(cellContents)
}
