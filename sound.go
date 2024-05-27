package main

import (
	"bytes"
	"fmt"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"io"
	"os"
	"time"
)

// playSoundTrack plays game soundtrack repeatedly
func playSoundTrack() {
	// Read the mp3 file into memory
	fileBytes, err := os.ReadFile("./assets/sound.mp3")
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.
	op := &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(decodedMp3)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

	// Start the loop
	for {
		if !player.IsPlaying() {
			// Now that the playSoundTrack finished playing, we can restart from the beginning
			_, err := player.Seek(0, io.SeekStart)
			if err != nil {
				panic(fmt.Errorf("player.Seek failed: %w", err))
			}
			player.Play()
		}
		time.Sleep(time.Second)
	}
}
