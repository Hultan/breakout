package game

import (
	"image/color"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var backgroundColor = color.RGBA{
	R: 128,
	G: 128,
	B: 128,
	A: 255,
}

var entities []drawable

type BreakOut struct {
	win      *gtk.ApplicationWindow
	da       *gtk.DrawingArea
	ticker   ticker
	speed    time.Duration
	breakout *breakout
}

func NewBreakOut(win *gtk.ApplicationWindow, da *gtk.DrawingArea) *BreakOut {
	b := &BreakOut{
		win:   win,
		da:    da,
		speed: 20,
	}
	return b
}

func (b *BreakOut) StartGame() {
	b.win.Connect("key-press-event", b.onKeyPress)
	b.da.Connect("draw", b.onDraw)

	b.ticker.ticker = time.NewTicker(b.speed * time.Millisecond)
	b.ticker.tickerQuit = make(chan struct{})

	a := b.win.GetAllocation()
	b.breakout = newBreakout("per", float64(a.GetWidth()), float64(a.GetHeight()))

	entities = append(entities, b.breakout.p)

	go b.mainLoop()
}

func (b *BreakOut) mainLoop() {
	for {
		select {
		case <-b.ticker.ticker.C:
			b.breakout.update()
			b.da.QueueDraw()
		case <-b.ticker.tickerQuit:
			b.breakout.isActive = false
			b.ticker.ticker.Stop()
			return
		}
	}
}

func (b *BreakOut) onDraw(_ *gtk.DrawingArea, ctx *cairo.Context) {
	b.drawBackground(ctx, backgroundColor)
	for i := range entities {
		entities[i].draw(ctx)
	}
}

func (b *BreakOut) drawBackground(ctx *cairo.Context, color color.Color) {
	ctx.SetSourceRGBA(getColor(color))
	ctx.Rectangle(0, 0, b.breakout.width, b.breakout.height)
	ctx.Fill()
}

func (b *BreakOut) onKeyPress(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	switch ke.KeyVal() {
	case gdk.KEY_q, gdk.KEY_Q:
		b.quit()
		b.win.Close()
	case gdk.KEY_Left:
		b.breakout.p.position.x -= 10
	case gdk.KEY_Right:
		b.breakout.p.position.x += 10
	}
}

func (b *BreakOut) quit() {
	if b.breakout.isActive {
		b.breakout.isActive = false
		close(b.ticker.tickerQuit) // Stop ticker
	}
}
