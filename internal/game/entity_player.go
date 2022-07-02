package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

const playerStartingWidth = 100.0
const playerStartingHeight = 15.0

var playerColor = color.RGBA{R: 200, G: 0, B: 0, A: 255}

type player struct {
	name string
	entity
	playerWidth float64
}

func newPlayer(name string, w, h float64) *player {
	x, y := (w-playerStartingWidth)/2, h-50
	return &player{
		name: name,
		entity: entity{
			rectangle: newRectangle(x, y, playerStartingWidth, playerStartingHeight),
			speed:     speed{},
			color:     playerColor,
		},
		playerWidth: playerStartingWidth,
	}
}

func (p *player) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.rect())
	ctx.Fill()
}

func (p *player) update() {
	turbo := 1.0

	theGame.keyIsPressedMutex.Lock()
	if theGame.keyIsPressed["shift"] {
		turbo = 2.0
	}
	if theGame.keyIsPressed["a"] || theGame.keyIsPressed["left"] {
		p.x -= 5 * turbo
		if p.x < 10 {
			p.x = 10
		}
	} else if theGame.keyIsPressed["d"] || theGame.keyIsPressed["right"] {
		p.x += 5 * turbo
		if p.x > theGame.width-p.playerWidth-10 {
			p.x = theGame.width - p.playerWidth - 10
		}
	}
	theGame.keyIsPressedMutex.Unlock()

	// If the ball is not moving, let it follow the player
	if !theGame.ball.isMoving {
		theGame.ball.x = p.x + p.playerWidth/2
	}
}

func (p *player) collide(e gameObject) {
	// To implement gameObject interface
}
