package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// tankImage holds the image for the tank component
var tankImage *ebiten.Image

// bombImage holds the image for the bomb components
var bombImage *ebiten.Image

// initImages initializes image variables
func initImages() {
	var err error

	tankImage, _, err = ebitenutil.NewImageFromFile("./assets/tank.png")
	if err != nil {
		panic("cannot load tank image")
	}

	bombImage, _, err = ebitenutil.NewImageFromFile("./assets/bomb.png")
	if err != nil {
		panic("cannot load bomb image")
	}
}
