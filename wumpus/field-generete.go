package main

import (
	"math/rand"
	"time"
)

func generate_field(size_x, size_y int) (field [][]Cell, player Player) {
	rand.Seed(time.Now().UnixNano())
	// Генерация случайного размера, если в функцию передали 0
	const MIN_FIELD_SIZE = 4
	const MAX_FIELD_SIZE = 8

	if size_x == 0 {
		size_x = rand.Intn(MAX_FIELD_SIZE-MIN_FIELD_SIZE+1) + MIN_FIELD_SIZE
	}
	if size_y == 0 {
		size_y = rand.Intn(MAX_FIELD_SIZE-MIN_FIELD_SIZE+1) + MIN_FIELD_SIZE
	}

	//Генерация пустого поля
	field = make([][]Cell, size_x)
	for i := range field {
		field[i] = make([]Cell, size_y)
	}

	player.position = arrange_objects(field)

	return
}

// Расстановка препятствий и игрока
func arrange_objects(field [][]Cell) (player_position [2]int) {
	fill_empty(field)
	player_position = put_object(field, PLAYER, PLAYER_SYMBOL)
	put_object(field, WUMPUS, WUMPUS_SYMBOL)
	put_object(field, GUN, GUN_SYMBOL)
	put_hole(field)

	return
}

func fill_empty(field [][]Cell) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = Cell{
				object:  EMPTY, //temp
				symbol:  EMPTY_SYMBOL,
				visited: true, //temp
			}
		}
	}
}

func put_hole(field [][]Cell) {
	const CHANCE_HOLE = 0.2

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j].object == 0 && rand.Float64() < CHANCE_HOLE {
				field[i][j] = Cell{
					object:  HOLE,
					symbol:  HOLE_SYMBOL,
					visited: true, //temp
				}
			}
		}
	}
}

func put_object(field [][]Cell, object_type int, object_symbol rune) (position [2]int) {
	for {
		position[0] = rand.Intn(len(field))
		position[1] = rand.Intn(len(field[0]))

		if field[position[0]][position[1]].object == 0 {
			field[position[0]][position[1]] = Cell{
				object:  object_type,
				symbol:  object_symbol,
				visited: true, // temp
			}
			break
		}
	}

	return
}
