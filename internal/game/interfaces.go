package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type gameObject interface {
	drawer
	updater
	collider
}

type drawer interface {
	draw(*cairo.Context)
}

type updater interface {
	update()
}

type collider interface {
	collide(object gameObject)
}
