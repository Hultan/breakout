package game

import (
	"fmt"
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
	if theGame.keysPressed['a'] {
		p.x -= 6
		if p.x < 10 {
			p.x = 10
		}
	} else if theGame.keysPressed['d'] {
		p.x += 6
		if p.x > theGame.width-p.playerWidth-10 {
			p.x = theGame.width - p.playerWidth - 10
		}
	}
}

func (p *player) collide(e gameObject) {
	fmt.Println("Collision with player")
}
