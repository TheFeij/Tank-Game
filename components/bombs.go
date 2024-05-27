package components

// Bomb represents a bomb in the game
type Bomb struct {
	// component represents the general attributes and behaviors of a component in the game
	component
}

// NewBomb creates and returns a new Bomb instance with the given configurations
func NewBomb(x, y, speed float64) Bomb {
	return Bomb{
		component: component{
			x:     x,
			y:     y,
			speed: speed,
		},
	}
}

// Move moves the bomb
func (b *Bomb) Move() {
	b.y += b.speed
}
