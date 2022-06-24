package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

type game struct {
	da            *gtk.DrawingArea
	play          *player
	width, height float64
}

var entities []drawable

func newGame(da *gtk.DrawingArea, name string, width, height float64) *game {
	g := &game{
		da:     da,
		play:   newPlayer(name, width, height),
		width:  width,
		height: height,
	}

	g.da.Connect("draw", g.onDraw)
	entities = append(entities, g.play)

	return g
}

func (b *game) update() {

}

func (b *game) draw() {
	b.da.QueueDraw()
}

func (b *game) onDraw(_ *gtk.DrawingArea, ctx *cairo.Context) {
	b.drawBackground(ctx, backgroundColor)
	for i := range entities {
		entities[i].draw(ctx)
	}
}

func (b *game) drawBackground(ctx *cairo.Context, color color.Color) {
	ctx.SetSourceRGBA(getColor(color))
	ctx.Rectangle(0, 0, b.width, b.height)
	ctx.Fill()
}
