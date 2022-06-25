package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}

type game struct {
	da            *gtk.DrawingArea
	play          *player
	width, height float64
}

var entities []drawable
var theBall *ball

func newGame(da *gtk.DrawingArea, name string, width, height float64) *game {
	g := &game{
		da:     da,
		play:   newPlayer(name, width, height),
		width:  width,
		height: height,
	}

	g.da.Connect("draw", g.onDraw)

	// Entities
	entities = append(entities, g.play)
	entities = append(entities, newCage())
	entities = append(entities, newCageBottom())
	theBall = newBall(g)
	entities = append(entities, theBall)

	return g
}

func (g *game) update() {
	for i := range entities {
		entities[i].update()
	}
}

func (g *game) draw() {
	// Triggers onDraw function
	g.da.QueueDraw()
}

func (g *game) onDraw(_ *gtk.DrawingArea, ctx *cairo.Context) {
	g.drawBackground(ctx, backgroundColor)
	for i := range entities {
		entities[i].draw(ctx, g)
	}
}

func (g *game) drawBackground(ctx *cairo.Context, color color.Color) {
	ctx.SetSourceRGBA(getColor(color))
	ctx.Rectangle(0, 0, g.width, g.height)
	ctx.Fill()
}
