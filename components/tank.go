package components

import (
	"Tank-Game/config"
	"time"
)

// Tank represents the player's tank
type Tank struct {
	// component represents the general attributes and behaviors of a component in the game
	component

	// lastShot records the time of the last shot fired by the tank. (for handling reloading logic)
	lastShot time.Time

	// ammo represents the current ammunition count of the tank.
	ammo int
}

func (t *Tank) Ammo() int {
	return t.ammo
}

func NewTank(x float64, y float64, speed float64, ammo int) Tank {
	return Tank{
		component: component{
			x:     x,
			y:     y,
			speed: speed,
		},
		lastShot: time.Now(),
		ammo:     ammo,
	}
}

// Move moves the tank
func (t *Tank) Move(direction Direction) {
	switch direction {
	case Right:
		t.x += t.speed
	case Left:
		t.x -= t.speed
	}
}

func (t *Tank) Shoot() (Bullet, error) {
	if t.ammo == 0 {
		return Bullet{}, ErrNotEnoughAmmunition
	}
	if t.lastShot.Add(500 * time.Millisecond).Before(time.Now()) {
		t.lastShot = time.Now()
		t.ammo--

		return NewBullet(t.x+config.TankWidth/2-config.BulletWidth/2, t.y, 5), nil
	}

	return Bullet{}, ErrReloading
}
