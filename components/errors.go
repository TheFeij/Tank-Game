package components

import "errors"

// defines possible errors for this package
var (
	ErrNotEnoughAmmunition = errors.New("not enough ammunition")
	ErrReloading           = errors.New("reloading")
)
