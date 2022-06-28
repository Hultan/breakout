package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type orientation int

const (
	orientationHorizontal orientation = iota
	orientationVertical
)

type cage struct {
	entity
	orientation
}

var cageColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}

func newCage(x, y, w, h float64, o orientation) *cage {
	return &cage{
		entity{
			rectangle:     newRectangle(x, y, w, h),
			collisionType: onCollisionBounce,
			color:         cageColor,
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
}

func (c *cage) collide(e gameObject) {
	fmt.Println("Collision with cage")
}
