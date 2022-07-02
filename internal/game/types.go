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

type speed struct {
	dx, dy float64
}

type entity struct {
	rectangle
	speed
	color color.Color
}
