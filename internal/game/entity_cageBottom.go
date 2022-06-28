package game

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
)

type cageBottom struct {
	entity
}

func newCageBottom(x, y, w, h float64) *cageBottom {
	return &cageBottom{
		entity{
			rectangle:     newRectangle(x, y, w, h),
			collisionType: onCollisionBallLost,
		},
	}
}

func (c *cageBottom) draw(ctx *cairo.Context) {
	// Does not need to be drawn

	// ctx.SetSourceRGBA(getColor(color.Black))
	// ctx.Rectangle(c.rect())
	// ctx.Fill()
}

func (c *cageBottom) update() {
}

func (c *cageBottom) collide(e gameObject) {
	fmt.Println("Game over!")
}

func (c *cageBottom) typ() entityType {
	return entityTypeCageBottom
}
