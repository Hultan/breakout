package game

import (
	"image/color"
)

type orientation int

const (
	orientationHorizontal orientation = iota
	orientationVertical
)

var backgroundColor = color.RGBA{R: 28, G: 28, B: 28, A: 200}
var ballColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var playerColor = color.RGBA{R: 200, G: 0, B: 0, A: 255}
var cageColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}
var scoreColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var pauseColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var brickCounterColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var brickColors = []color.RGBA{
	{R: 0, G: 50, B: 0, A: 255},
	{R: 0, G: 100, B: 0, A: 255},
	{R: 0, G: 150, B: 0, A: 255},
}

// Misc constants
const (
	gameSpeed               = 20
	cageWidth               = 10.0
	brickWidth              = 20.0
	brickHeight             = 15.0
	brickLeftMargin         = 50.0
	levelHeight             = 30.0
	scoreMultiplier         = 10
	newEntityCollectionSize = 10
)

// Score text position
const (
	scoreLeft   = 20
	scoreTop    = 593
	scoreWidth  = 200
	scoreHeight = 20
)

// Pause text position
const (
	pauseLeft   = 240
	pauseTop    = 280
	pauseWidth  = 100
	pauseHeight = 40
)

// Brick counter text position
const (
	brickCounterTop    = 585
	brickCounterLeft   = 593
	brickCounterWidth  = 200
	brickCounterHeight = 20
)

// Ball constants
const (
	ballSize           = 10
	ballStartingSpeedX = 0
	ballStartingSpeedY = 7
)

// Player constants
const (
	playerStartingWidth        = 100.0
	playerStartingHeight       = 15.0
	playerBounce               = 10
	playerBorderOffset         = 50
	playerSpeed                = 5
	playerSpeedTurboMultiplier = 2.0
)
