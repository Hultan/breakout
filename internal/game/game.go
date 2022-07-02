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
var entities []gameObject
var nonGameEntities []gameObject

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
	width, height     float64
	level             int
	isPaused          bool
}

func newGame(da *gtk.DrawingArea, name string, w, h float64) *game {
	g := &game{
		da:           da,
		keyIsPressed: make(map[string]bool, 5),
		ball:         newBall(),
		player:       newPlayer(name, w, h),
		scorer:       newScore(),
		counter:      newBrickCounter(),
		width:        w,
		height:       h,
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

	// Non game entities, score etc
	nonGameEntities = append(nonGameEntities, g.scorer)
	nonGameEntities = append(nonGameEntities, g.counter)

	return g
}

func (g *game) initialize() {
	g.loadLevel()
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
	g.drawBackground(ctx)
	// Draw game entities
	for i := range entities {
		entities[i].draw(ctx)
	}
	// Draw non-game entities (like texts, etc)
	for i := range nonGameEntities {
		nonGameEntities[i].draw(ctx)
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
				entities = append(entities, b)
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
	nonGameEntities = append(nonGameEntities, g.pauser)
}

func (g *game) resume() {
	// Resume game
	g.isPaused = false

	// Find pause entity
	pauseIndex := -1
	for i, e := range nonGameEntities {
		if e == g.pauser {
			pauseIndex = i
			break
		}
	}

	// Remove pause entity
	if pauseIndex > 0 {
		nonGameEntities[pauseIndex] = nonGameEntities[len(nonGameEntities)-1]
		nonGameEntities = nonGameEntities[:len(nonGameEntities)-1]
	}
}
