package game

import (
	"github.com/gotk3/gotk3/cairo"
)

type player struct {
	name string
	entity
	playerWidth float64
}

func newPlayer(name string, w, h float64) *player {
	x, y := (w-playerStartingWidth)/2, h-playerBorderOffset
	return &player{
		name: name,
		entity: entity{
			rectangle: newRectangle(x, y, playerStartingWidth, playerStartingHeight),
			color:     playerColor,
		},
		playerWidth: playerStartingWidth,
	}
}

func (p *player) draw(ctx *cairo.Context) {
	ctx.SetSourceRGBA(getColor(p.color))
	ctx.Rectangle(p.rect())
	ctx.Fill()
}

func (p *player) update() {
	turbo := 1.0

	theGame.keyIsPressedMutex.Lock()
	if theGame.keyIsPressed["shift"] {
		turbo = playerSpeedTurboMultiplier
	}
	if theGame.keyIsPressed["a"] || theGame.keyIsPressed["left"] {
		p.x -= playerSpeed * turbo
		if p.x < cageWidth {
			p.x = cageWidth
		}
	} else if theGame.keyIsPressed["d"] || theGame.keyIsPressed["right"] {
		p.x += playerSpeed * turbo
		if p.x > theGame.w-p.playerWidth-cageWidth {
			p.x = theGame.w - p.playerWidth - cageWidth
		}
	}
	theGame.keyIsPressedMutex.Unlock()

	// If the ball is not moving, let it follow the player
	if !theGame.ball.isMoving {
		theGame.ball.x = p.x + p.playerWidth/2
	}
}

func (p *player) collide(e gameObject) {
	// To implement gameObject interface
}
