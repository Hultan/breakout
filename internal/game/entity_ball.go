package game

import (
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
)

const ballSize = 10
const ballStartingSpeedX = 0
const ballStartingSpeedY = 3

var ballColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type ball struct {
	entity
}

func newBall(g *game) *ball {
	return &ball{entity{
		position:      position{g.width / 2, g.height * 2 / 3},
		speed:         speed{ballStartingSpeedX, ballStartingSpeedY},
		collisionType: onCollisionNone,
		color:         ballColor,
	}}
}

func (b *ball) draw(ctx *cairo.Context, g *game) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.Arc(b.position.x, b.position.y, ballSize, 0, math.Pi*2)
	ctx.Fill()
}

func (b *ball) update() {
	b.position.x += b.speed.dx
	b.position.y += b.speed.dy
}
