package game

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
)

type particleEmitter struct {
	entity
	particles entityCollection
}

func newParticleEmitter(rec rectangle, col color.Color, size emitterSize) *particleEmitter {
	p := &particleEmitter{
		entity: entity{
			rectangle: rec,
			color:     col,
		},
	}

	for i := 0; i < int(size); i++ {
		p.particles = append(p.particles, newParticle(rec, col, int(size)*3))
	}

	return p
}

func (p *particleEmitter) update() {
	// Update particles
	for _, o := range p.particles {
		o.update()
	}
}

func (p *particleEmitter) draw(ctx *cairo.Context) {
	for i, o := range p.particles {
		// Draw particles that are still "alive"
		part := p.particles[i].(*particle)
		if part.lifetime > 0 {
			o.draw(ctx)
		}
	}
}

func (p *particleEmitter) collide(e gameObject) {
	// To implement gameObject interface
}

func (p *particleEmitter) alive() bool {
	// Check if there is still particles alive
	for i := range p.particles {
		part := p.particles[i].(*particle)
		if part.lifetime > 0 {
			return true
		}
	}
	return false
}
