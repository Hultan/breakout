package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

const playerStartingSize = 100.0

var playerColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}

type player struct {
	name string
	entity
	size float64
}

func newPlayer(name string, width, height float64) *player {
	return &player{name, entity{
		position:      position{(width - playerStartingSize) / 2, height - 50},
		speed:         speed{},
		collisionType: onCollisionBounce,
		color:         playerColor,
	}, playerStartingSize}
}

func (p *player) draw(ctx *cairo.Context, _ *game) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.position.x, p.position.y, p.size, 20)
	ctx.Fill()
}

func (p *player) update() {
}
