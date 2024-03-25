package configreader

import (
	"encoding/json"
	"os"
)

type Config struct {
	EmptySymbol,
	EmptyVisitedSymbol,
	PlayerSymbol,
	ArmedPlayerSymbol,
	HoleSymbol,
	WumpusSymbol,
	GunSymbol,

	BaseColor,
	WindColor,
	SmellColor,
	SmellAndWindColor,
	ColorBold,
	ColorReset string

	MinFieldSize,
	MaxFieldSize int

	ChanceHole float64

	QuickInput bool

	LastPrintDelay bool
}

func LoadConstantsFromFile(filename string) (config *Config, err error) {

	config = newConfig()
	file, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	return
}

func newConfig() *Config {
	return &Config{
		EmptySymbol:        " ",
		EmptyVisitedSymbol: "·",
		PlayerSymbol:       "□",
		ArmedPlayerSymbol:  "■",
		HoleSymbol:         "⬤",
		WumpusSymbol:       "☠",
		GunSymbol:          "➽",

		BaseColor:         "\033[30m", // черный
		WindColor:         "\033[34m", // синий
		SmellColor:        "\033[31m", // красный
		SmellAndWindColor: "\033[35m", // фиолетовый
		ColorBold:         "\033[1m",
		ColorReset:        "\033[0m",

		QuickInput: true,

		MinFieldSize: 4,
		MaxFieldSize: 8,

		ChanceHole: 0.15,

		LastPrintDelay: true,
	}
}
