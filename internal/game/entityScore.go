package game

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
)

type score struct {
	entity
	score int
}

func newScore() *score {
	return &score{
		entity: entity{
			rectangle: newRectangle(scoreLeft, scoreTop, scoreWidth, scoreHeight),
			color:     scoreColor,
		},
	}
}

func (s *score) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(s.color))
	ctx.SelectFontFace("Roboto Thin", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(s.h)
	ctx.MoveTo(s.x, s.y)
	text := fmt.Sprintf("Score : %d", s.score)
	ctx.ShowText(text)
}

func (s *score) update() {
	// To implement gameObject interface
}

func (s *score) collide(e gameObject) {
	// To implement gameObject interface
}

func (s *score) addScore(points int) {
	s.score += points
}
