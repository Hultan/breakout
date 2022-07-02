package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type brickCounter struct {
	entity
	needCount bool
	count     int
}

var brickCounterColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func newBrickCounter() *brickCounter {
	return &brickCounter{
		entity: entity{
			rectangle: newRectangle(685, 593, 200, 20),
			speed:     speed{},
			color:     brickCounterColor,
		},
		needCount: true,
	}
}

func (b *brickCounter) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.SelectFontFace("Roboto Thin", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(b.h)
	ctx.MoveTo(b.x, b.y)
	text := fmt.Sprintf("Bricks : %d", b.countBricks())
	ctx.ShowText(text)
}

func (b *brickCounter) update() {
	// To implement gameObject interface
}

func (b *brickCounter) collide(e gameObject) {
	// To implement gameObject interface
}

func (b *brickCounter) countBricks() int {
	if !b.needCount {
		return b.count
	}

	count := 0
	for _, object := range theGame.entities {
		br, ok := object.(*brick)
		if ok && !br.dead {
			count++
		}
	}

	b.count = count
	b.needCount = false
	return count
}
