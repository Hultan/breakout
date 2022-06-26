package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type cage struct {
	entity
}

var cageColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}

func newCage(x, y, w, h float64) *cage {
	return &cage{
		entity{
			rectangle:     newRectangle(x, y, w, h),
			collisionType: onCollisionBounce,
			color:         cageColor,
		},
	}
}

func (c *cage) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(c.color))
	ctx.Rectangle(c.rect())
	ctx.Fill()
}

func (c *cage) update() {
	fmt.Println("cage")
}

func (c *cage) collide(e gameObject) {
}

func (c *cage) typ() entityType {
	return entityTypeCage
}
