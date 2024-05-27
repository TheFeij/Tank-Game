package main

// gameStatus represents possible game statuses
type gameStatus int

const (
	// running when game is being played
	running = iota
	// paused when game is paused
	paused
	// victory when game is finished and player is victorious
	victory
	// gameOver when the game is lost
	gameOver
)
