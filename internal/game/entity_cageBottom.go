package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type cageBottom struct {
	entity
}

func newCageBottom() *cageBottom {
	return &cageBottom{
		entity{
			collisionType: onCollisionBallLost,
		},
	}
}

func (c *cageBottom) draw(_ *cairo.Context, _ *game) {
	// Does not need to be drawn
}

func (c *cageBottom) update() {
}
