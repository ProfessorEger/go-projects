package main

func calculate_action(field [][]Cell, ptr_player *Player, action Action, ptr_massage *string) (stop_game bool) {
	switch action.action_type {
	case MOVE:
		stop_game = calculate_move(field, ptr_player, action.direction, ptr_massage)
	case SHOT:
		stop_game = calculate_shot(field, ptr_player, action.direction, ptr_massage)
	}
	return
}

func calculate_move(field [][]Cell, ptr_player *Player, direction int, ptr_massage *string) (stop_game bool) {
	change_position(field, &ptr_player.position, direction)
	stop_game = check_position(field, ptr_player, ptr_massage)
	if !stop_game {
		calculate_feeling(field, ptr_player.position, ptr_massage)
	}

	return
}

// делает шаг
func change_position(field [][]Cell, ptr_player_position *[2]int, direction int) (old_player_position [2]int) {
	old_player_position = *ptr_player_position

	field[ptr_player_position[0]][ptr_player_position[1]].object = EMPTY
	field[ptr_player_position[0]][ptr_player_position[1]].symbol = EMPTY_SYMBOL

	switch direction {
	case UP:
		if (*ptr_player_position)[1] != 0 {
			(*ptr_player_position)[1]--
		}
	case RIGHT:
		if (*ptr_player_position)[0] != len(field)-1 {
			(*ptr_player_position)[0]++
		}
	case DOWN:
		if (*ptr_player_position)[1] != len(field[0])-1 {
			(*ptr_player_position)[1]++
		}
	case LEFT:
		if (*ptr_player_position)[0] != 0 {
			(*ptr_player_position)[0]--
		}
	}

	return
}

func check_position(field [][]Cell, ptr_player *Player, ptr_massage *string) (stop_game bool) {
	switch field[ptr_player.position[0]][ptr_player.position[1]].object {
	case HOLE:
		stop_game = true
		*ptr_massage = "Вы упали в яму. Игра окончена"
	case WUMPUS:
		stop_game = true
		*ptr_massage = "Вас съел Вампус. Игра окончена"
	case GUN:
		ptr_player.armed++
		*ptr_massage = "Вы нашли пистолет с одним патроном"
	}

	if !stop_game {
		field[ptr_player.position[0]][ptr_player.position[1]] = Cell{
			object:  PLAYER,
			symbol:  PLAYER_SYMBOL,
			visited: true,
		}
	}

	return
}

func calculate_shot(field [][]Cell, ptr_player *Player, direction int, ptr_massage *string) (stop_game bool) {

	return
}

func calculate_feeling(field [][]Cell, position [2]int, ptr_massage *string) {
	var neighboring_cells [4]Cell
	if position[0] != 0 {
		neighboring_cells[0] = field[position[0]-1][position[1]]
	}
	if position[0] != len(field)-1 {
		neighboring_cells[1] = field[position[0]+1][position[1]]
	}
	if position[1] != 0 {
		neighboring_cells[2] = field[position[0]][position[1]-1]
	}
	if position[1] != len(field[0])-1 {
		neighboring_cells[3] = field[position[0]][position[1]+1]
	}

	var wind, smell bool
	for _, cell := range neighboring_cells {

		switch cell.object {
		case HOLE:
			wind = true
		case WUMPUS:
			smell = true
		}
	}

	if wind && smell {
		field[position[0]][position[1]].color = SMELL_AND_WIND_COLOR
		*ptr_massage += " " + "Вы чувствуете ветер и запах Вампуса!"
	} else if wind {
		field[position[0]][position[1]].color = WIND_COLOR
		*ptr_massage += " " + "Вы чувствуете ветер"
	} else if smell {
		field[position[0]][position[1]].color = SMELL_COLOR
		*ptr_massage += " " + "Вы чувствуете запах вампуса!"
	}
}
