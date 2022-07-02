package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type brickCounter struct {
	entity
	needCount  bool
	brickCount int
}

var brickCounterColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func newBrickCounter() *brickCounter {
	return &brickCounter{
		entity: entity{
			rectangle:     newRectangle(685, 593, 200, 20),
			speed:         speed{},
			collisionType: onCollisionNone,
			color:         brickCounterColor,
		},
		needCount: true,
	}
}

func (b *brickCounter) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(b.color))
	ctx.SelectFontFace("Roboto Thin", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(b.h)
	ctx.MoveTo(b.x, b.y)
	text := fmt.Sprintf("Bricks : %d", b.count())
	ctx.ShowText(text)
}

func (b *brickCounter) update() {
}

func (b *brickCounter) collide(e gameObject) {
}

func (b *brickCounter) count() int {
	if !b.needCount {
		return b.brickCount
	}

	count := 0
	for _, object := range entities {
		br, ok := object.(*brick)
		if ok && !br.dead {
			count++
		}
	}

	b.brickCount = count
	b.needCount = false
	return count
}
