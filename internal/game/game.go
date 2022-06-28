package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}
var entities []gameObject

type game struct {
	da            *gtk.DrawingArea
	keysPressed   map[rune]bool
	player        *player
	ball          *ball
	width, height float64
}

func newGame(da *gtk.DrawingArea, name string, w, h float64) *game {
	g := &game{
		da:          da,
		keysPressed: make(map[rune]bool, 5),
		ball:        newBall(w, h),
		player:      newPlayer(name, w, h),
		width:       w,
		height:      h,
	}

	// Events
	g.da.Connect("draw", g.onDraw)

	// Entities
	entities = append(entities, g.player)
	entities = append(entities, newCage(0, 0, 10, h, orientationVertical))   // left cage
	entities = append(entities, newCage(0, 0, w, 10, orientationHorizontal)) // top cage
	entities = append(entities, newCage(w-10, 0, w, h, orientationVertical)) // right cage
	entities = append(entities, newCageBottom(0, h-10, w, h))                // bottom cage
	entities = append(entities, g.ball)

	entities = append(entities, newBrick(3, 100, 100))

	return g
}

func (g *game) initialize() {
	g.ball.resetBallPosition()
}

func (g *game) update() {
	for i := range entities {
		entities[i].update()
	}
}

func (g *game) checkCollision() {
	for i := range entities {
		e := entities[i]
		if g.ball.collidesWith(e) {
			g.ball.collide(e)
			e.collide(g.ball)
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
	ctx.Rectangle(0, 0, theGame.width, theGame.height)
	ctx.Fill()
}
