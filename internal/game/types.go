package game

import (
	"image/color"
)

type position struct {
	x, y float64
}

func newPosition(x, y float64) position {
	return position{x, y}
}

func (p position) move(s speed) position {
	return newPosition(p.x+s.dx, p.y+s.dy)
}

type size struct {
	w, h float64
}

func newSize(w, h float64) size {
	return size{w, h}
}

type rectangle struct {
	position
	size
}

func newRectangle(x, y, w, h float64) rectangle {
	return rectangle{newPosition(x, y), newSize(w, h)}
}

func (r rectangle) rect() (float64, float64, float64, float64) {
	return r.x, r.y, r.w, r.h
}

func (r rectangle) center() (float64, float64) {
	return r.x + r.w/2, r.y + r.h/2
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

type entityType int

const (
	entityTypePlayer entityType = iota
	entityTypeBall
	entityTypeCage
	entityTypeCageBottom
)

type entity struct {
	rectangle
	speed
	collisionType
	entityType
	color color.Color
}
