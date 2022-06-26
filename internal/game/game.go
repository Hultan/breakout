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

var entities []gameObject
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
	entities = append(entities, newCage(0, 0, 10, g.height))                      // left cage
	entities = append(entities, newCage(0, 0, g.width, 10))                       // top cage
	entities = append(entities, newCage(g.width-10, 0, g.width, g.height))        // right cage
	entities = append(entities, newCageBottom(0, g.height-10, g.width, g.height)) // bottom cage
	theBall = newBall(g)
	entities = append(entities, theBall)

	return g
}

func (g *game) update() {
	for i := range entities {
		entities[i].update()
	}
}

func (g *game) checkCollision() {
	for i := range entities {
		e := entities[i]
		if theBall.collidesWith(e) {
			theBall.collide(e)
			e.collide(theBall)
		}
	}
}

func (g *game) draw() {
	// Triggers onDraw function
	g.da.QueueDraw()
}

func (g *game) onDraw(_ *gtk.DrawingArea, ctx *cairo.Context) {
	g.drawBackground(ctx, backgroundColor)
	for i := range entities {
		entities[i].draw(ctx)
	}
}

func (g *game) drawBackground(ctx *cairo.Context, color color.Color) {
	ctx.SetSourceRGBA(getColor(color))
	ctx.Rectangle(0, 0, g.width, g.height)
	ctx.Fill()
}
