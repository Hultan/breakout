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
	game       *game
}

var windowWidth, windowHeight float64

func NewBreakOut(win *gtk.ApplicationWindow, da *gtk.DrawingArea) *BreakOut {
	b := &BreakOut{
		win:   win,
		speed: 20,
	}

	b.win.Connect("key-press-event", b.onKeyPress)
	b.win.Connect("key-release-event", b.onKeyRelease)
	a := b.win.GetAllocation()
	windowWidth, windowHeight = float64(a.GetWidth()), float64(a.GetHeight())
	b.game = newGame(da, "per")

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
			b.game.update()
			b.game.checkCollision()
			b.game.draw()
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
	case gdk.KEY_Left:
		thePlayer.isLeftPressed = true
	case gdk.KEY_Right:
		thePlayer.isRightPressed = true
	}
}

func (b *BreakOut) onKeyRelease(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	switch ke.KeyVal() {
	case gdk.KEY_q, gdk.KEY_Q:
		b.quit()
		b.win.Close()
	case gdk.KEY_Left:
		thePlayer.isLeftPressed = false
	case gdk.KEY_Right:
		thePlayer.isRightPressed = false
	}
}

func (b *BreakOut) quit() {
	close(b.tickerQuit) // Stop ticker
}
