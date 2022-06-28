package game

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{R: 128, G: 128, B: 128, A: 255}
var entities []gameObject

const (
	brickWidth  = 30.0
	levelHeight = 40.0
)

type game struct {
	da            *gtk.DrawingArea
	keysPressed   map[string]bool
	player        *player
	ball          *ball
	width, height float64
}

func newGame(da *gtk.DrawingArea, name string, w, h float64) *game {
	g := &game{
		da:          da,
		keysPressed: make(map[string]bool, 5),
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

	g.loadLevel(1)

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

func (g *game) loadLevel(level int) {
	h := levelHeight
	for _, row := range levels[level].bricks {
		w := 50.0
		fields := strings.Fields(row)
		for _, s := range fields {
			col, err := strconv.Atoi(string(s[0]))
			if err != nil {
				panic(err)
			}
			entities = append(entities, newBrick(col-1, len(s), w, h))
			w += float64(len(s))*brickWidth + brickWidth
		}
		h += levelHeight
	}
}
