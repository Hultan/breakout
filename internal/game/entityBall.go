package game

import (
	"math"

	"github.com/gotk3/gotk3/cairo"
)

type ball struct {
	entity
	isMoving bool
}

func newBall() *ball {
	return &ball{
		entity{
			color: ballColor,
		},
		false,
	}
}

func (b *ball) resetBallPosition() {
	b.rectangle = newRectangle(
		theGame.player.x+theGame.player.playerWidth/2,
		theGame.player.y-ballSize,
		ballSize,
		ballSize,
	)
	b.isMoving = false
	b.dx, b.dy = 0.0, 0.0
}

func (b *ball) startMoving() {
	if !b.isMoving {
		b.isMoving = true
		b.speed = speed{ballStartingSpeedX, ballStartingSpeedY}
	}
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
		// Ball bounces against a cage wall
		if o.orientation == orientationHorizontal {
			b.speed.dy = -b.speed.dy
		} else {
			b.speed.dx = -b.speed.dx
		}
	case *player:
		// Ball bounces against the player

		// Calculate the distance from the center of the paddle
		d := (o.x + o.w/2 - b.x) / o.w
		b.speed.dx = b.speed.dx - d*playerBounce
		b.speed.dy = -b.speed.dy
	case *cageBottom:
		// Player missed the ball, end of ball

		// end of game, for now
		// b.resetBallPosition()
		b.speed.dy = -b.speed.dy
	case *brick:
		// Ball bounces against a brick
		b.speed.dy = -b.speed.dy
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
	case *brick:
		if o.dead {
			return false
		}
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
