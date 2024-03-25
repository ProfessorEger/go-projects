## Hunt the Wumpus

Console version of the game Hunt the Wumpus.
To win you need to kill Wumpus. To do this, there is a gun hidden on the map. The gun shoots at an adjacent cell in the selected direction. Wumpus, Holes and gun are randomly placed on the map. On the cells adjacent to the wumpus you can smell its smell, and next to the holes you can feel the wind.
To move use W/A/S/D. To shoot Ctrl+W/A/S/D


## Settings
All user settings are stored in ```hunt-the-wumpus/config.json```

# Default settings
```
"EmptySymbol": " ",
"EmptyVisitedSymbol": "·",
"PlayerSymbol": "□",
"ArmedPlayerSymbol": "■",
"HoleSymbol": "⬤",
"WumpusSymbol": "☠",
"GunSymbol": "➽",
"WindColor": "\u001b[34m",
"SmellColor": "\u001b[31m",
"SmellAndWindColor": "\u001b[35m",
"ColorBold": "\u001b[1m",
"ColorReset": "\u001b[0m",
"MinFieldSize": 4,
"MaxFieldSize": 8,
"ChanceHole": 0.15,
"QuickInput": true,
"LastPrintDelay": true
```

# For Windows, the following settings are recommended:

``` json
{
	"WindColor": "",
	"SmellColor": "",
	"SmellAndWindColor": "",
	"ColorBold": "",
	"ColorReset": ""
}
```
