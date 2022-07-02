package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type brick struct {
	entity
	dead bool
}

var brickColors = []color.RGBA{
	{R: 0, G: 50, B: 0, A: 255},
	{R: 0, G: 100, B: 0, A: 255},
	{R: 0, G: 150, B: 0, A: 255},
}

func newBrick(col, size int, x, y float64) *brick {
	return &brick{
		entity: entity{
			rectangle: newRectangle(x, y, float64(size)*brickWidth, 20),
			color:     brickColors[col],
		},
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
}

func (b *brick) collide(e gameObject) {
}
