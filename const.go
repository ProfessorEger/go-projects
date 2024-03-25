package main

import "wumpus/configreader"

const (
	EMPTY = iota
	PLAYER
	HOLE
	WUMPUS
	GUN
)

type Player struct {
	position [2]int
	armed    int
	symbol   string
}

type Cell struct {
	object  int
	symbol  string
	visible bool
	color   string
}

type Action struct {
	ActionType int
	Direction  int
}

// action_type
const (
	MOVE = iota
	SHOT
)

// direction
const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

/*const (
	EMPTY_SYMBOL  = '·'
	PLAYER_SYMBOL = '□'
	HOLE_SYMBOL   = '⬤'
	WUMPUS_SYMBOL = '☠'
	GUN_SYMBOL    = '➽'
)

const (
	WIND_COLOR           = "\033[34m" // синий
	SMELL_COLOR          = "\033[31m" // желтый
	SMELL_AND_WIND_COLOR = "\033[35m" // фиолетовый

	COLOR_BOLD  = "\033[1m"
	COLOR_RESET = "\033[0m"
)*/

var config configreader.Config
