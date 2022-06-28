package game

import (
	"image/color"
	"math"

	"github.com/gotk3/gotk3/cairo"
)

const ballSize = 10
const ballStartingSpeedX = 0
const ballStartingSpeedY = 5

var ballColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type ball struct {
	entity
	g *game
}

func newBall(g *game) *ball {
	return &ball{
		entity{
			rectangle:     newRectangle(g.width/2, g.height*2/3, ballSize, ballSize),
			speed:         speed{ballStartingSpeedX, ballStartingSpeedY},
			collisionType: onCollisionNone,
			color:         ballColor,
		},
		g,
	}
}

func (b *ball) resetBallPosition() {
	b.rectangle = newRectangle(b.g.width/2, b.g.height*2/3, ballSize, ballSize)
	b.speed = speed{ballStartingSpeedX, ballStartingSpeedY}
}

func (b *ball) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.Arc(b.position.x, b.position.y, ballSize, 0, math.Pi*2)
	ctx.Fill()
}

func (b *ball) update() {
	b.position = b.position.move(b.speed)
}

func (b *ball) collide(e gameObject) {
	switch o := e.(type) {
	case *cage:
		if o.orientation == orientationHorizontal {
			b.speed.dy = -b.speed.dy
		} else {
			b.speed.dx = -b.speed.dx
		}
	case *player:
		// cos(theta) = adjacent/hypotenuse

		// Calculate the distance from the center of the paddle
		d := (o.x + o.w/2 - b.x) / o.w
		b.speed.dx = b.speed.dx - d*7
		b.speed.dy = -b.speed.dy
	case *cageBottom:
		// end of game, for now
		b.resetBallPosition()
	}
}

func (b *ball) collidesWith(e gameObject) bool {
	var r rectangle
	switch o := e.(type) {
	case *cage:
		r = o.rectangle
	case *player:
		r = o.rectangle
	case *cageBottom:
		r = o.rectangle
	}

	// Check if the ball intersects the rectangle
	// https://stackoverflow.com/questions/401847/circle-rectangle-collision-detection-intersection
	closestX := clamp(b.x, r.x, r.x+r.w)
	closestY := clamp(b.y, r.y, r.y+r.h)
	dx := b.x - closestX
	dy := b.y - closestY
	dsq := dx*dx + dy*dy
	return dsq < b.w*b.w
}

func (b *ball) typ() entityType {
	return entityTypeBall
}
