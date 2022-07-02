package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

var pauseColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type pause struct {
	entity
}

func newPause() *pause {
	return &pause{
		entity{
			rectangle: newRectangle(240, 280, 100, 40),
			speed:     speed{},
			color:     pauseColor,
		},
	}
}

func (s *pause) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(s.color))
	ctx.SelectFontFace("Roboto Thin", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(s.h)
	ctx.MoveTo(s.x, s.y)
	text := "Game is paused!"
	ctx.ShowText(text)
}

func (s *pause) update() {
	// To implement gameObject interface
}

func (s *pause) collide(e gameObject) {
	// To implement gameObject interface
}
