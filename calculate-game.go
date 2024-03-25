package main

func calculateAction(field [][]Cell, ptrPlayer *Player, action Action, ptrMessage *string) (stopGame bool) {
	switch action.ActionType {
	case MOVE:
		stopGame = calculateMove(field, ptrPlayer, action.Direction, ptrMessage)
	case SHOT:
		stopGame = calculateShot(field, ptrPlayer, action.Direction, ptrMessage)
	}
	return
}

func calculateMove(field [][]Cell, ptrPlayer *Player, direction int, ptrMessage *string) (stopGame bool) {
	changePosition(field, &ptrPlayer.position, direction)
	stopGame = checkPosition(field, ptrPlayer, ptrMessage)
	if !stopGame {
		calculateFeeling(field, ptrPlayer.position, ptrMessage)
	}

	return
}

// Moves the player
func changePosition(field [][]Cell, ptrPlayerPosition *[2]int, direction int) (oldPlayerPosition [2]int) {
	oldPlayerPosition = *ptrPlayerPosition

	field[ptrPlayerPosition[0]][ptrPlayerPosition[1]].object = EMPTY
	field[ptrPlayerPosition[0]][ptrPlayerPosition[1]].symbol = config.EmptyVisitedSymbol

	switch direction {
	case UP:
		if ptrPlayerPosition[1] != 0 {
			ptrPlayerPosition[1]--
		}
	case RIGHT:
		if ptrPlayerPosition[0] != len(field)-1 {
			ptrPlayerPosition[0]++
		}
	case DOWN:
		if ptrPlayerPosition[1] != len(field[0])-1 {
			ptrPlayerPosition[1]++
		}
	case LEFT:
		if ptrPlayerPosition[0] != 0 {
			ptrPlayerPosition[0]--
		}
	}

	return
}

func checkPosition(field [][]Cell, ptrPlayer *Player, ptrMessage *string) (stopGame bool) {
	switch field[ptrPlayer.position[0]][ptrPlayer.position[1]].object {
	case HOLE:
		stopGame = true
		*ptrMessage = "You fell into a hole. Game over."
	case WUMPUS:
		stopGame = true
		*ptrMessage = "The Wumpus ate you. Game over."
	case GUN:
		ptrPlayer.armed++
		ptrPlayer.symbol = config.ArmedPlayerSymbol
		*ptrMessage = "You found a gun with one bullet."
	}

	if !stopGame {
		field[ptrPlayer.position[0]][ptrPlayer.position[1]] = Cell{
			object:  PLAYER,
			symbol:  ptrPlayer.symbol,
			visible: true,
		}
	}

	return
}

func calculateShot(field [][]Cell, ptrPlayer *Player, direction int, ptrMessage *string) (stopGame bool) {
	targetPosition := ptrPlayer.position
	ptrPlayer.armed--
	if ptrPlayer.armed <= 0 {
		ptrPlayer.symbol = config.PlayerSymbol
	}

	switch direction {
	case UP:
		if ptrPlayer.position[1] != 0 {
			targetPosition[1]--
		}
	case RIGHT:
		if ptrPlayer.position[0] != len(field)-1 {
			targetPosition[0]++
		}
	case DOWN:
		if ptrPlayer.position[1] != len(field[0])-1 {
			targetPosition[1]++
		}
	case LEFT:
		if ptrPlayer.position[0] != 0 {
			targetPosition[0]--
		}
	}

	if field[targetPosition[0]][targetPosition[1]].object == WUMPUS {
		stopGame = true
		*ptrMessage += " You killed Wumpus! Victory!"

		field[targetPosition[0]][targetPosition[1]] = Cell{
			object: EMPTY,
			symbol: config.EmptySymbol,
		}

	} else {
		*ptrMessage += " miss"
	}
	return
}

func calculateFeeling(field [][]Cell, position [2]int, ptrMessage *string) {
	var neighboringCells [4]Cell
	if position[0] != 0 {
		neighboringCells[0] = field[position[0]-1][position[1]]
	}
	if position[0] != len(field)-1 {
		neighboringCells[1] = field[position[0]+1][position[1]]
	}
	if position[1] != 0 {
		neighboringCells[2] = field[position[0]][position[1]-1]
	}
	if position[1] != len(field[0])-1 {
		neighboringCells[3] = field[position[0]][position[1]+1]
	}

	var wind, smell bool
	for _, cell := range neighboringCells {

		switch cell.object {
		case HOLE:
			wind = true
		case WUMPUS:
			smell = true
		}
	}

	if wind && smell {
		field[position[0]][position[1]].color = config.SmellAndWindColor
		*ptrMessage += " You feel the breeze and smell the Wumpus!"
	} else if wind {
		field[position[0]][position[1]].color = config.WindColor
		*ptrMessage += " You feel the breeze."
	} else if smell {
		field[position[0]][position[1]].color = config.SmellColor
		*ptrMessage += " You smell the Wumpus!"
	}
}
