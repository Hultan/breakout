package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type cage struct {
	entity
	orientation
}

func newCage(x, y, w, h float64, o orientation) *cage {
	return &cage{
		entity{
			rectangle: newRectangle(x, y, w, h),
			color:     cageColor,
		},
		o,
	}
}

func (c *cage) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(c.color))
	ctx.Rectangle(c.rect())
	ctx.Fill()
}

func (c *cage) update() {
	// To implement gameObject interface
}

func (c *cage) collide(e gameObject) {
	// To implement gameObject interface
}
