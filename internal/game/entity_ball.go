package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
)

const ballSize = 10
const ballStartingSpeedX = 0
const ballStartingSpeedY = 3

var ballColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type ball struct {
	entity
}

func newBall(g *game) *ball {
	return &ball{entity{
		rectangle:     newRectangle(g.width/2, g.height*2/3, ballSize, ballSize),
		speed:         speed{ballStartingSpeedX, ballStartingSpeedY},
		collisionType: onCollisionNone,
		color:         ballColor,
	}}
}

func (b *ball) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.Arc(b.position.x, b.position.y, ballSize, 0, math.Pi*2)
	ctx.Fill()
}

func (b *ball) update() {
	fmt.Println("ball")
	b.position = b.position.move(b.speed)
}

func (b *ball) collide(e gameObject) {
	fmt.Println(e)
}

func (b *ball) collidesWith(e gameObject) bool {
	switch o := e.(type) {
	case *cage:

		p := o.position
		s := o.size
		dx := math.Abs(b.position.x - p.x)
		dy := math.Abs(b.position.y - p.y)

		if dx > (s.w/2 + b.size.h) {
			return false
		}
		if dy > (s.h/2 + b.size.h) {
			return false
		}

		if dx <= (s.w / 2) {
			return true
		}
		if dy <= (s.h / 2) {
			return true
		}

		// Calculate corner distance
		d := math.Pow(dx-s.w/2, 2) +
			math.Pow(dy-s.h/2, 2)
		return d <= math.Pow(b.size.h, 2)

	}

	return false
}

func (b *ball) typ() entityType {
	return entityTypeBall
}
