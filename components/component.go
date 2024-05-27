package components

// component represents a component in the game
type component struct {
	x     float64
	y     float64
	speed float64
}

// X getter for x
func (c component) X() float64 {
	return c.x
}

// Y getter for y
func (c component) Y() float64 {
	return c.y
}

// Speed getter for speed
func (c component) Speed() float64 {
	return c.speed
}
