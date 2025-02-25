package definitions

type ColorSquare struct {
	Color string

	// meta square
	ShortcutTarget	int
	ShortcutName	string
	IconSquare		string
	StuckSquare		string
}

type Card struct {
	Color	string
	Icon	string
	Double	bool
}

type Player struct {
	Name	string
	Square	int
}

var BoardSquares = []ColorSquare{
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange", ShortcutName: "Rainbow Trail", ShortcutTarget: 60},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{IconSquare: "Sugar Plum"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{IconSquare: "Candy Cane"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "purple"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "purple"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple", ShortcutName: "Gumdrop Pass", ShortcutTarget: 48},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{IconSquare: "Gum Drop"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"}, // shortcut target
	{Color: "yellow", StuckSquare: "Gooey Gumdrops"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"}, // shortcut target
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "orange"},
	{IconSquare: "Gramma Nut"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue", StuckSquare: "Lost in Lollipop Woods"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{IconSquare: "Lollipop"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{IconSquare: "Frostine"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red", StuckSquare: "Stuck in Molasses Swamp"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
	{Color: "yellow"},
	{Color: "blue"},
	{Color: "orange"},
	{Color: "green"},
	{Color: "red"},
	{Color: "purple"},
}

var Deck = []Card{
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red"},
	{Color: "red", Double: true},
	{Color: "red", Double: true},

	{Color: "orange"},
	{Color: "orange"},
	{Color: "orange"},
	{Color: "orange"},
	{Color: "orange"},
	{Color: "orange"},
	{Color: "orange", Double: true},
	{Color: "orange", Double: true},

	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow"},
	{Color: "yellow", Double: true},
	{Color: "yellow", Double: true},

	{Color: "green"},
	{Color: "green"},
	{Color: "green"},
	{Color: "green"},
	{Color: "green"},
	{Color: "green"},
	{Color: "green", Double: true},
	{Color: "green", Double: true},

	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue"},
	{Color: "blue", Double: true},
	{Color: "blue", Double: true},

	{Color: "purple"},
	{Color: "purple"},
	{Color: "purple"},
	{Color: "purple"},
	{Color: "purple"},
	{Color: "purple"},
	{Color: "purple", Double: true},
	{Color: "purple", Double: true},

	{Icon: "Sugar Plum"},
	{Icon: "Candy Cane"},
	{Icon: "Gum Drop"},
	{Icon: "Gramma Nut"},
	{Icon: "Lollipop"},
	{Icon: "Frostine"},
}