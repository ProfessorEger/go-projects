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

	return
}

// делает шаг
func change_position(field [][]Cell, ptr_player_position *[2]int, direction int) {
	switch direction {
	case UP:
		if (*ptr_player_position)[1] != 0 {
			(*ptr_player_position)[1]--
		}
	case RIGHT:
		if (*ptr_player_position)[0] != len(field) {
			(*ptr_player_position)[0]++
		}
	case DOWN:
		if (*ptr_player_position)[1] != len(field[0]) {
			(*ptr_player_position)[1]++
		}
	case LEFT:
		if (*ptr_player_position)[0] != 0 {
			(*ptr_player_position)[0]--
		}
	}
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

	return
}

func calculate_shot(field [][]Cell, ptr_player *Player, direction int, ptr_massage *string) (stop_game bool) {

	return
}
