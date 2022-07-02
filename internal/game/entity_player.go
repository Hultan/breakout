package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

const playerStartingWidth = 100.0
const playerStartingHeight = 20.0

var playerColor = color.RGBA{R: 50, G: 0, B: 0, A: 255}

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
			rectangle:     newRectangle(x, y, playerStartingWidth, playerStartingHeight),
			speed:         speed{},
			collisionType: onCollisionBounce,
			color:         playerColor,
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
	if theGame.keysPressed["shift"] {
		turbo = 1.8
	}
	if theGame.keysPressed["a"] || theGame.keysPressed["left"] {
		p.x -= 6 * turbo
		if p.x < 10 {
			p.x = 10
		}
	} else if theGame.keysPressed["d"] || theGame.keysPressed["right"] {
		p.x += 6 * turbo
		if p.x > theGame.width-p.playerWidth-10 {
			p.x = theGame.width - p.playerWidth - 10
		}
	}
}

func (p *player) collide(e gameObject) {
}
