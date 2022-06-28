package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

const playerStartingWidth = 100.0
const playerStartingHeight = 20.0

var playerColor = color.RGBA{R: 100, G: 0, B: 0, A: 255}

type player struct {
	name string
	entity
}

func newPlayer(name string, width, height float64) *player {
	x, y := (width-playerStartingWidth)/2, height-50
	return &player{name, entity{
		rectangle:     newRectangle(x, y, playerStartingWidth, playerStartingHeight),
		speed:         speed{},
		collisionType: onCollisionBounce,
		color:         playerColor,
	}}
}

func (p *player) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.rect())
	ctx.Fill()
}

func (p *player) update() {
}

func (p *player) collide(e gameObject) {
	fmt.Println("Collision with player")
}

func (p *player) typ() entityType {
	return entityTypePlayer
}
