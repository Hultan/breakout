package game

import (
	"fmt"
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
		speed: gameSpeed,
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

			if b.isLevelOver() {
				if b.loadNextLevel() == false {
					// No more levels, or error, quit game
					b.quit()
				}
			}
		case <-b.tickerQuit:
			b.ticker.Stop()
			return
		}
	}
}

func (b *BreakOut) isLevelOver() bool {
	return theGame.counter.countBricks() == 0
}

func (b *BreakOut) loadNextLevel() bool {
	theGame.level++
	if theGame.level == len(levels) {
		return false
	}
	err := theGame.loadLevel()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
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
		if theGame.isPaused {
			theGame.resume()
		} else {
			theGame.pause()
		}
	case gdk.KEY_Left:
		theGame.keyIsPressed["left"] = false
	case gdk.KEY_Right:
		theGame.keyIsPressed["right"] = false
	case gdk.KEY_Shift_L, gdk.KEY_Shift_R:
		theGame.keyIsPressed["shift"] = false
	case gdk.KEY_space:
		theGame.ball.startMoving()
	}
	theGame.keyIsPressedMutex.Unlock()
}

func (b *BreakOut) quit() {
	close(b.tickerQuit) // Stop ticker
	b.win.Close()
}
