package main

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
}

type Cell struct {
	object  int
	symbol  rune
	visited bool
	color   string
}

type Action struct {
	action_type int
	direction   int
}

const (
	MOVE = iota
	SHOT
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

const (
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
)
