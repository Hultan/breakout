package game

import (
	"fmt"
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

const playerStartingWidth = 100.0
const playerStartingHeight = 20.0

var playerColor = color.RGBA{R: 50, G: 0, B: 0, A: 255}

type player struct {
	name string
	entity
	playerWidth    float64
	windowWidth    float64
	isLeftPressed  bool
	isRightPressed bool
}

func newPlayer(name string, winWidth, winHeight float64) *player {
	x, y := (winWidth-playerStartingWidth)/2, winHeight-50
	return &player{
		name: name,
		entity: entity{
			rectangle:     newRectangle(x, y, playerStartingWidth, playerStartingHeight),
			speed:         speed{},
			collisionType: onCollisionBounce,
			color:         playerColor,
		},
		playerWidth: playerStartingWidth,
		windowWidth: winWidth,
	}
}

func (p *player) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.rect())
	ctx.Fill()
}

func (p *player) update() {
	if p.isLeftPressed {
		p.x -= 7
		if p.x < 10 {
			p.x = 10
		}
	} else if p.isRightPressed {
		p.x += 7
		if p.x > p.windowWidth-p.playerWidth-10 {
			p.x = p.windowWidth - p.playerWidth - 10
		}
	}
}

func (p *player) collide(e gameObject) {
	fmt.Println("Collision with player")
}

func (p *player) typ() entityType {
	return entityTypePlayer
}
