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

type entity struct {
	position
	speed
	color color.Color
}

type ball struct {
	entity
}

type enemy struct {
	entity
}
