package game

import (
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
	// To implement gameObject interface
}

func (c *cageBottom) update() {
	// To implement gameObject interface
}

func (c *cageBottom) collide(e gameObject) {
	// To implement gameObject interface
}
