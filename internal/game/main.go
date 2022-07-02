package game

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type BreakOut struct {
	win        *gtk.ApplicationWindow
	tickerQuit chan struct{}
	ticker     *time.Ticker
	speed      time.Duration
}

var theGame *game

func NewBreakOut(win *gtk.ApplicationWindow, da *gtk.DrawingArea) *BreakOut {
	b := &BreakOut{
		win:   win,
		speed: 20,
	}

	// Events
	b.win.Connect("key-press-event", b.onKeyPress)
	b.win.Connect("key-release-event", b.onKeyRelease)

	// Create game object
	a := win.GetAllocation()
	theGame = newGame(da, "per", float64(a.GetWidth()), float64(a.GetHeight()))
	theGame.initialize()

	// Create ticker object
	b.ticker = time.NewTicker(b.speed * time.Millisecond)
	b.tickerQuit = make(chan struct{})

	return b
}

func (b *BreakOut) Start() {
	go b.mainLoop()
}

func (b *BreakOut) mainLoop() {

	for {
		select {
		case <-b.ticker.C:
			if theGame.isPaused {
				theGame.draw()
				continue
			}
			theGame.update()
			theGame.checkCollision()
			theGame.draw()

			if theGame.counter.countBricks() == 0 {
				theGame.level++
				if theGame.level == len(levels) {
					// No more levels, quit game
					b.quit()
					return
				}
				err := theGame.loadLevel()
				if err != nil {
					panic(err)
				}
			}
		case <-b.tickerQuit:
			b.ticker.Stop()
			return
		}
	}
}

func (b *BreakOut) onKeyPress(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	theGame.keyIsPressedMutex.Lock()
	switch ke.KeyVal() {
	case gdk.KEY_a, gdk.KEY_A:
		theGame.keyIsPressed["a"] = true
	case gdk.KEY_d, gdk.KEY_D:
		theGame.keyIsPressed["d"] = true
	case gdk.KEY_Left:
		theGame.keyIsPressed["left"] = true
	case gdk.KEY_Right:
		theGame.keyIsPressed["right"] = true
	case gdk.KEY_Shift_L, gdk.KEY_Shift_R:
		theGame.keyIsPressed["shift"] = true
	case gdk.KEY_space:
		theGame.ball.startMoving()
	}
	theGame.keyIsPressedMutex.Unlock()
}

func (b *BreakOut) onKeyRelease(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	theGame.keyIsPressedMutex.Lock()
	switch ke.KeyVal() {
	case gdk.KEY_q, gdk.KEY_Q:
		b.quit()
	case gdk.KEY_a, gdk.KEY_A:
		theGame.keyIsPressed["a"] = false
	case gdk.KEY_d, gdk.KEY_D:
		theGame.keyIsPressed["d"] = false
	case gdk.KEY_p, gdk.KEY_P:
		pauseGame()
	case gdk.KEY_Left:
		theGame.keyIsPressed["left"] = false
	case gdk.KEY_Right:
		theGame.keyIsPressed["right"] = false
	case gdk.KEY_Shift_L, gdk.KEY_Shift_R:
		theGame.keyIsPressed["shift"] = false
	}
	theGame.keyIsPressedMutex.Unlock()
}

func pauseGame() {
	if theGame.isPaused {
		// Resume game
		theGame.isPaused = false

		// Find pause entity
		pauseIndex := -1
		for i, e := range nonGameEntities {
			if e == theGame.pause {
				pauseIndex = i
				break
			}
		}

		// Remove pause entity
		if pauseIndex > 0 {
			nonGameEntities[pauseIndex] = nonGameEntities[len(nonGameEntities)-1]
			nonGameEntities = nonGameEntities[:len(nonGameEntities)-1]
		}
	} else {
		// Pause game
		theGame.isPaused = true
		theGame.pause = newPause()
		nonGameEntities = append(nonGameEntities, theGame.pause)
	}
}

func (b *BreakOut) quit() {
	close(b.tickerQuit) // Stop ticker
	b.win.Close()
}
