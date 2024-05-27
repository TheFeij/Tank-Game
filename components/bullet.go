package components

// Bullet represents a bullet in the game
type Bullet struct {
	// component represents the general attributes and behaviors of a component in the game
	component
}

// NewBullet creates and returns a new bullet with the given configurations
func NewBullet(x, y, speed float64) Bullet {
	return Bullet{
		component: component{
			x:     x,
			y:     y,
			speed: speed,
		},
	}
}

// Move moves the bullet
func (b *Bullet) Move() {
	b.y -= b.speed
}

// IsWasted checks if the bullet is wasted or not
// in other words, if the bullet has gone out of the screen or not
func (b *Bullet) IsWasted() bool {
	return b.y < 0
}
