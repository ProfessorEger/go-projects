package main

import (
	"math/rand"
)

func generateField(sizeX, sizeY int) (field [][]Cell, player Player) {
	// Generate random size if 0 is passed to the function
	if sizeX == 0 {
		sizeX = rand.Intn(config.MaxFieldSize-config.MinFieldSize+1) + config.MinFieldSize
	}
	if sizeY == 0 {
		sizeY = rand.Intn(config.MaxFieldSize-config.MinFieldSize+1) + config.MinFieldSize
	}

	// Generate an empty field
	field = make([][]Cell, sizeX)
	for i := range field {
		field[i] = make([]Cell, sizeY)
	}

	player.position = arrangeObjects(field)
	player.symbol = config.PlayerSymbol

	return
}

// Arrange obstacles and player
func arrangeObjects(field [][]Cell) (playerPosition [2]int) {
	fillEmpty(field)
	playerPosition = putObject(field, PLAYER, config.PlayerSymbol)
	putObject(field, WUMPUS, config.WumpusSymbol)
	putObject(field, GUN, config.GunSymbol)
	putHole(field)

	return
}

func fillEmpty(field [][]Cell) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = Cell{
				object: EMPTY,
				symbol: config.EmptySymbol,
			}
		}
	}
}

func putHole(field [][]Cell) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j].object == 0 && rand.Float64() < config.ChanceHole {
				field[i][j] = Cell{
					object: HOLE,
					symbol: config.HoleSymbol,
					//visited: true, //temp
				}
			}
		}
	}
}

func putObject(field [][]Cell, objectType int, objectSymbol string) (position [2]int) {
	for {
		position[0] = rand.Intn(len(field))
		position[1] = rand.Intn(len(field[0]))

		if field[position[0]][position[1]].object == EMPTY {
			field[position[0]][position[1]] = Cell{
				object: objectType,
				symbol: objectSymbol,
				//visited: true, // temp
			}
			break
		}
	}

	return
}

func MakeVisible(field [][]Cell) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			field[i][j].visible = true
		}
	}
}
