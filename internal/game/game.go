package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}

type game struct {
	da          *gtk.DrawingArea
	keysPressed map[rune]bool
}

var entities []gameObject

func newGame(da *gtk.DrawingArea, name string) *game {
	m := make(map[rune]bool, 5)
	g := &game{
		da:          da,
		keysPressed: m,
	}
	g.da.Connect("draw", g.onDraw)

	// Entities
	thePlayer = newPlayer(name)
	theBall = newBall()

	entities = append(entities, thePlayer)
	entities = append(entities, newCage(0, 0, 10, windowHeight, orientationVertical))                       // left cage
	entities = append(entities, newCage(0, 0, windowWidth, 10, orientationHorizontal))                      // top cage
	entities = append(entities, newCage(windowWidth-10, 0, windowWidth, windowHeight, orientationVertical)) // right cage
	entities = append(entities, newCageBottom(0, windowHeight-10, windowWidth, windowHeight))               // bottom cage
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
	ctx.Rectangle(0, 0, windowWidth, windowHeight)
	ctx.Fill()
}
