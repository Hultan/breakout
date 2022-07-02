package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type brick struct {
	entity
	points int
	dead   bool
}

func newBrick(col, size int, x, y float64) *brick {
	return &brick{
		entity: entity{
			rectangle: newRectangle(x, y, float64(size)*brickWidth, brickHeight),
			color:     brickColors[col-1],
		},
		points: col * scoreMultiplier,
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
