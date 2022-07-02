package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type brick struct {
	entity
	points int
	dead   bool
}

var brickColors = []color.RGBA{
	{R: 0, G: 50, B: 0, A: 255},
	{R: 0, G: 100, B: 0, A: 255},
	{R: 0, G: 150, B: 0, A: 255},
}

func newBrick(col, size int, x, y float64) *brick {
	return &brick{
		entity: entity{
			rectangle: newRectangle(x, y, float64(size)*brickWidth, 15),
			color:     brickColors[col-1],
		},
		points: col * 10,
	}
}

func (b *brick) draw(ctx *cairo.Context) {
	if b.dead {
		return
	}
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.Rectangle(b.rect())
	ctx.Fill()
}

func (b *brick) update() {
	// To implement gameObject interface
}

func (b *brick) collide(e gameObject) {
	theGame.counter.needCount = true
	b.dead = true
	theGame.scorer.addScore(b.points)
}
