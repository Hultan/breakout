package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type drawable interface {
	draw(*cairo.Context, *game)
	update()
}
