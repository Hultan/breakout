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

func NewBreakOut(win *gtk.ApplicationWindow, da *gtk.DrawingArea) *BreakOut {
	b := &BreakOut{
		win:   win,
		speed: 20,
	}

	b.win.Connect("key-press-event", b.onKeyPress)
	a := b.win.GetAllocation()
	b.game = newGame(da, "per", float64(a.GetWidth()), float64(a.GetHeight()))

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
		b.game.play.position.x -= 10
	case gdk.KEY_Right:
		b.game.play.position.x += 10
	}
}

func (b *BreakOut) quit() {
	close(b.tickerQuit) // Stop ticker
}
