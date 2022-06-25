package game

import (
	"image/color"
)

type position struct {
	x, y float64
}

type speed struct {
	dx, dy float64
}

type collisionType int

const (
	onCollisionNone collisionType = iota
	onCollisionBounce
	onCollisionExplode
	onCollisionBallLost
)

type entity struct {
	position
	speed
	collisionType
	color color.Color
}
