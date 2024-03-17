package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printGame(field [][]Cell, ptrMessage *string) {

	var buffer bytes.Buffer

	addFieldToBuffer(field, &buffer)
	buffer.WriteString(*ptrMessage)
	ClearConsole()
	fmt.Println(buffer.String())
	*ptrMessage = ""
}

func addFieldToBuffer(field [][]Cell, ptrBuffer *bytes.Buffer) {

	for j := 0; j <= len(field[0])*2; j++ {
		for i := 0; i <= len(field)*2; i++ {
			if i == 0 && j == 0 {
				ptrBuffer.WriteString("  ┌─")
			} else if i == 0 && j == len(field[0])*2 {
				ptrBuffer.WriteString("  └─")
			} else if i == len(field)*2 && j == 0 {
				ptrBuffer.WriteString("─┐")
			} else if i == len(field)*2 && j == len(field[0])*2 {
				ptrBuffer.WriteString("─┘")
			} else if i == 0 && j%2 == 0 {
				ptrBuffer.WriteString("  ├─")
			} else if i%2 == 0 && j == 0 {
				ptrBuffer.WriteString("─┬─")
			} else if i == len(field)*2 && j%2 == 0 {
				ptrBuffer.WriteString("─┤")
			} else if i%2 == 0 && j == len(field[0])*2 {
				ptrBuffer.WriteString("─┴─")
			} else if i%2 == 0 && j%2 == 0 {
				ptrBuffer.WriteString("─┼─")
			} else if i == 0 {
				ptrBuffer.WriteString("  │ ")
			} else if j == 0 {
				ptrBuffer.WriteString("─")
			} else if i == len(field)*2 {
				ptrBuffer.WriteString(" │")
			} else if j == len(field[0])*2 {
				ptrBuffer.WriteString("─")
			} else if i%2 == 0 {
				ptrBuffer.WriteString(" │ ")
			} else if j%2 == 0 {
				ptrBuffer.WriteString("─")
			} else if i%2 == 1 && j%2 == 1 {
				addCellToBuffer(field[(i-1)/2][(j-1)/2], ptrBuffer)
			}
		}
		ptrBuffer.WriteString("\n")
	}

}

func addCellToBuffer(cell Cell, ptrBuffer *bytes.Buffer) {
	if cell.visible {
		ptrBuffer.WriteString(fmt.Sprintf("%s%s%s%s", config.ColorBold, cell.color, cell.symbol, config.ColorReset))
	} else {
		ptrBuffer.WriteString(" ")
	}
}

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
