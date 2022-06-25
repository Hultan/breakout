package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type cage struct {
	entity
}

var cageColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}

func newCage() *cage {
	return &cage{
		entity{
			collisionType: onCollisionBounce,
			color:         cageColor,
		},
	}
}

func (c *cage) draw(ctx *cairo.Context, g *game) {
	ctx.SetSourceRGBA(getColor(c.color))
	ctx.Rectangle(0, 0, 10, g.height)
	ctx.Rectangle(0, 0, g.width, 10)
	ctx.Rectangle(g.width-10, 0, g.width, g.height)
	ctx.Fill()
}

func (c *cage) update() {
}
