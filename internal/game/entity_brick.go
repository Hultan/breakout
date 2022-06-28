package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type brick struct {
	entity
}

var brickColor = color.RGBA{R: 0, G: 0, B: 50, A: 255}

func newBrick(size, x, y float64) *brick {
	return &brick{
		entity{
			rectangle: newRectangle(x, y, size*40, 20),
			color:     brickColor,
		},
	}
}

func (b *brick) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.Rectangle(b.rect())
	ctx.Fill()
}

func (b *brick) update() {
}

func (b *brick) collide(e gameObject) {
	fmt.Println("Collision with brick")
}
