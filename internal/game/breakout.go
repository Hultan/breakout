package game

type breakout struct {
	p             *player
	width, height float64
	isActive      bool
}

func newBreakout(name string, width, height float64) *breakout {
	return &breakout{
		p:        newPlayer(name),
		width:    width,
		height:   height,
		isActive: true,
	}
}

func (b *breakout) update() {

}
