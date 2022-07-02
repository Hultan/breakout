package game

import (
	"image/color"
	"strconv"
	"strings"
	"sync"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{R: 28, G: 28, B: 28, A: 200}

const (
	brickWidth  = 20.0
	levelHeight = 30.0
)

type game struct {
	da                *gtk.DrawingArea
	keyIsPressed      map[string]bool
	keyIsPressedMutex sync.RWMutex
	player            *player
	ball              *ball
	scorer            *score
	pauser            *pause
	counter           *brickCounter
	entities          entityCollection
	nonGameEntities   entityCollection
	width, height     float64
	level             int
	isPaused          bool
}

func newGame(da *gtk.DrawingArea, name string, w, h float64) *game {
	g := &game{
		da:              da,
		keyIsPressed:    make(map[string]bool, 5),
		ball:            newBall(),
		player:          newPlayer(name, w, h),
		scorer:          newScore(),
		counter:         newBrickCounter(),
		entities:        newEntityCollection(),
		nonGameEntities: newEntityCollection(),
		width:           w,
		height:          h,
	}

	// Events
	g.da.Connect("draw", g.onDraw)

	// Entities
	g.entities.add(g.player)
	g.entities.add(newCage(0, 0, 10, h, orientationVertical))   // left cage
	g.entities.add(newCage(0, 0, w, 10, orientationHorizontal)) // top cage
	g.entities.add(newCage(w-10, 0, w, h, orientationVertical)) // right cage
	g.entities.add(newCageBottom(0, h-10, w, h))                // bottom cage
	g.entities.add(g.ball)

	// Non game entities, score etc
	g.nonGameEntities.add(g.scorer)
	g.nonGameEntities.add(g.counter)

	return g
}

func (g *game) initialize() {
	g.loadLevel()
	g.ball.resetBallPosition()
}

func (g *game) update() {
	for i := range g.entities {
		g.entities[i].update()
	}
}

func (g *game) checkCollision() {
	for i := range g.entities {
		e := g.entities[i]
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
	g.drawBackground(ctx)
	// Draw game entities
	for i := range g.entities {
		g.entities[i].draw(ctx)
	}
	// Draw non-game entities (like texts, etc)
	for i := range g.nonGameEntities {
		g.nonGameEntities[i].draw(ctx)
	}
}

func (g *game) drawBackground(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(backgroundColor))
	ctx.Rectangle(0, 0, g.width, g.height)
	ctx.Fill()
}

func (g *game) loadLevel() error {
	h := levelHeight
	for _, row := range levels[g.level].bricks {
		w := 50.0
		fields := strings.Fields(row)
		for _, s := range fields {
			brickType, err := strconv.Atoi(string(s[0]))
			if err != nil {
				return levelError{g.level, "unknown brick type"}
			}
			if brickType > 0 {
				b := newBrick(brickType, len(s), w, h)
				g.entities.add(b)
			}
			w += float64(len(s))*brickWidth + brickWidth
		}
		h += levelHeight
	}
	return nil
}

func (g *game) pause() {
	// Pause game
	g.isPaused = true
	g.pauser = newPause()
	g.nonGameEntities.add(g.pauser)
}

func (g *game) resume() {
	// Resume game
	g.isPaused = false

	// Find pause entity
	pauseIndex := -1
	for i, e := range g.nonGameEntities {
		if e == g.pauser {
			pauseIndex = i
			break
		}
	}

	// Remove pause entity
	if pauseIndex > 0 {
		g.nonGameEntities[pauseIndex] = g.nonGameEntities[len(g.nonGameEntities)-1]
		g.nonGameEntities = g.nonGameEntities[:len(g.nonGameEntities)-1]
	}
}
