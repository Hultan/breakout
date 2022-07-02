package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type pause struct {
	entity
}

func newPause() *pause {
	return &pause{
		entity{
			rectangle: newRectangle(pauseLeft, pauseTop, pauseWidth, pauseHeight),
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
