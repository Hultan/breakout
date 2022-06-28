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
			theGame.update()
			theGame.checkCollision()
			theGame.draw()
		case <-b.tickerQuit:
			b.ticker.Stop()
			return
		}
	}
}

func (b *BreakOut) onKeyPress(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	switch ke.KeyVal() {
	case gdk.KEY_q, gdk.KEY_Q:
		b.quit()
		b.win.Close()
	case gdk.KEY_a, gdk.KEY_A:
		theGame.keysPressed["a"] = true
	case gdk.KEY_d, gdk.KEY_D:
		theGame.keysPressed["d"] = true
	case gdk.KEY_Left:
		theGame.keysPressed["left"] = true
	case gdk.KEY_Right:
		theGame.keysPressed["right"] = true
	case gdk.KEY_Shift_L, gdk.KEY_Shift_R:
		theGame.keysPressed["shift"] = true
	}
}

func (b *BreakOut) onKeyRelease(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	switch ke.KeyVal() {
	case gdk.KEY_q, gdk.KEY_Q:
		b.quit()
		b.win.Close()
	case gdk.KEY_a, gdk.KEY_A:
		theGame.keysPressed["a"] = false
	case gdk.KEY_d, gdk.KEY_D:
		theGame.keysPressed["d"] = false
	case gdk.KEY_Left:
		theGame.keysPressed["left"] = false
	case gdk.KEY_Right:
		theGame.keysPressed["right"] = false
	case gdk.KEY_Shift_L, gdk.KEY_Shift_R:
		theGame.keysPressed["shift"] = false
	}
}

func (b *BreakOut) quit() {
	close(b.tickerQuit) // Stop ticker
}
