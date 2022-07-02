package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

var scoreColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type score struct {
	entity
	score int
}

func newScore() *score {
	return &score{
		entity{
			rectangle:     newRectangle(15, 593, 200, 20),
			speed:         speed{},
			collisionType: onCollisionNone,
			color:         scoreColor,
		},
		0,
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
}

func (s *score) collide(e gameObject) {
}

func (s *score) addScore(points int) {
	s.score += points
}
