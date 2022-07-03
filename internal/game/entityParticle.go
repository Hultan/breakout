package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/gotk3/gotk3/cairo"
)

type particle struct {
	entity
	lifetime     int
	acceleration speed
}

func newParticle(rec rectangle, col color.Color, l int) *particle {
	// Calculate a random angle and speed
	angle := rand.Float64() * math.Pi * 2
	dx, dy := math.Cos(angle), math.Sin(angle)
	s := speed{dx, dy}
	a := speed{0, -.1}
	return &particle{
		entity: entity{
			rectangle: rec,
			speed:     s,
			color:     col,
		},
		lifetime:     l,
		acceleration: a,
	}
}

func (p *particle) update() {
	// Particles are affected by "gravity"
	p.speed.dx -= p.acceleration.dx
	p.speed.dy -= p.acceleration.dy

	// To create some randomness in particle movement
	f := rand.Float64()
	p.position.x += p.speed.dx * f * 2
	p.position.y += p.speed.dy * f * 2
	p.lifetime--
}

func (p *particle) draw(ctx *cairo.Context) {
	// Fade out the particles by lowering their alpha
	r, g, b, a := getColor(p.color)
	a = float64(p.lifetime) / 50.0
	ctx.SetSourceRGBA(r, g, b, a)

	x, y := p.position.x+p.w/2, p.position.y+p.h/2
	ctx.Arc(x, y, 5, 0, math.Pi*2)
	ctx.Fill()
}

func (p *particle) collide(e gameObject) {
	// To implement gameObject interface
}
