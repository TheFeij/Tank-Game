package main

import (
	"Tank-Game/config"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
	"log"
)

// resultFont for the end of the game message indicating win or loss
var resultFont font.Face

// scoreFont for displaying the score while game is paused or finished
var scoreFont font.Face

// statusFont for displaying status information on the top left of the screen during the game
var statusFont font.Face

// initFonts initializes different fonts of the game
func initFonts() {
	var err error

	resultFont, err = loadFont(config.FontPath, 36)
	if err != nil {
		log.Fatal(err)
	}

	scoreFont, err = loadFont(config.FontPath, 48)
	if err != nil {
		log.Fatal(err)
	}

	statusFont, err = loadFont(config.FontPath, 20)
	if err != nil {
		log.Fatal(err)
	}
}

// loadFont loads a font from a given file path and returns it as a font.Face with the specified size.
func loadFont(filePath string, size float64) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	const dpi = 72
	return opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
