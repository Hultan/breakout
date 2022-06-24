package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type player struct {
	name string
	entity
	size float64
}

func newPlayer(name string, width, height float64) *player {
	size := 100.0
	return &player{name, entity{
		position: position{(width - size) / 2, height - 50},
		speed:    speed{},
		color: color.RGBA{
			R: 100,
			G: 0,
			B: 0,
			A: 255,
		},
	}, size}
}

func (p *player) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.position.x, p.position.y, p.size, 20)
	ctx.Fill()
}

func (p *player) update() {
}
