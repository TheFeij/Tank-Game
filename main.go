package main

import (
	"Tank-Game/components"
	"Tank-Game/config"
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"log"
	"math/rand"
	"os"
)

// Game holds components, score, status of the game
// implements ebiten.Game interface
type game struct {
	// tank is the player
	tank components.Tank
	// bullets all the bullets in the game
	bullets []components.Bullet
	// bombs all the bombs in the game
	bombs []components.Bomb
	// score number of bombs shot down by player (by the tank)
	score int
	// tick game tick
	tick int
	// status holds status of the game
	status gameStatus
}

// Update updates game elements
func (g *game) Update() error {
	// increment game tick
	g.tick++

	// check if game the game is finished
	if g.status == victory || g.status == gameOver {
		if g.tick == 120 {
			os.Exit(0)
		}
		return nil
	}

	// check if the Escape Key is pressed to pause or unpause the game
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if g.status == paused {
			g.status = running
		} else if g.status == running {
			g.status = paused
		}
		return nil
	}

	// check if the game is paused
	if g.status == paused {
		return nil
	}

	// move tank
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.tank.X() > 0 {
		g.tank.Move(components.Left)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.tank.X() < config.ScreenWidth-config.TankWidth {
		g.tank.Move(components.Right)
	}

	// shoot bullets
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bullet, err := g.tank.Shoot()
		if err != nil {
			if errors.Is(err, components.ErrNotEnoughAmmunition) {
				g.status = victory
				g.tick = 0
				return nil
			}
		}
		g.bullets = append(g.bullets, bullet)
	}

	// move bullets
	for i := 0; i < len(g.bullets); i++ {
		g.bullets[i].Move()
		if g.bullets[i].IsWasted() {
			g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
			i--
		}
	}

	// spawn bombs
	if g.tick > 120 {
		g.tick = 0
		g.bombs = append(g.bombs, components.NewBomb(float64(rand.Intn(config.ScreenWidth-config.BombWidth-40)+20), 0, 2))
	}

	// move bombs
	for i := 0; i < len(g.bombs); i++ {
		g.bombs[i].Move()
		if g.bombs[i].Y() > config.ScreenHeight {
			g.status = gameOver
			g.tick = 0
			return nil
		}
	}

	// check collisions
	for i := 0; i < len(g.bombs); i++ {
		for j := 0; j < len(g.bullets); j++ {
			if g.bullets[j].X() < g.bombs[i].X()+config.BombWidth &&
				g.bullets[j].X()+config.BulletWidth > g.bombs[i].X() &&
				g.bullets[j].Y() < g.bombs[i].Y()+config.BombHeight/3 &&
				g.bullets[j].Y()+config.BulletHeight > g.bombs[i].Y() {
				g.bombs = append(g.bombs[:i], g.bombs[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
				i--
				break
			}
		}
	}

	return nil
}

// Draw draws different elements of the game
func (g *game) Draw(screen *ebiten.Image) {
	// Draw a game over message
	if g.status == victory {
		finalScoreText := fmt.Sprintf("Final Score: %d", g.score)
		victoryText := "You Win!"

		textWidth := text.BoundString(resultFont, victoryText).Max.X
		text.Draw(screen, victoryText, resultFont, (config.ScreenWidth-textWidth)/2, config.ScreenHeight/2-20, color.White)

		scoreWidth := text.BoundString(scoreFont, finalScoreText).Max.X
		text.Draw(screen, finalScoreText, scoreFont, (config.ScreenWidth-scoreWidth)/2, config.ScreenHeight/2+20, color.White)

		return
	} else if g.status == gameOver {
		gameOverText := "Game Over!"

		textWidth := text.BoundString(resultFont, gameOverText).Max.X
		text.Draw(screen, gameOverText, resultFont, (config.ScreenWidth-textWidth)/2, config.ScreenHeight/2-20, color.White)

		return
	} else if g.status == paused {
		scoreText := fmt.Sprintf("Score: %d", g.score)
		gamePausedText := "Game Paused"

		textWidth := text.BoundString(resultFont, gamePausedText).Max.X
		text.Draw(screen, gamePausedText, resultFont, (config.ScreenWidth-textWidth)/2, config.ScreenHeight/2-20, color.White)

		scoreWidth := text.BoundString(scoreFont, scoreText).Max.X
		text.Draw(screen, scoreText, scoreFont, (config.ScreenWidth-scoreWidth)/2, config.ScreenHeight/2+20, color.White)
		return
	}

	// Draw tank
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.tank.X(), g.tank.Y())
	screen.DrawImage(tankImage, op)

	// Draw bullets
	for _, bullet := range g.bullets {
		ebitenutil.DrawRect(screen, bullet.X(), bullet.Y(), config.BulletWidth, config.BulletHeight, color.White)
	}

	// Draw bombs
	for _, e := range g.bombs {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.X(), e.Y())
		screen.DrawImage(bombImage, op)
	}

	// Draw score
	statusText := fmt.Sprintf("Score: %d\tBullets: %d", g.score, g.tank.Ammo())
	text.Draw(screen, statusText, statusFont, 8, 30, color.White)
}

// Layout returns Screen dimensions
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	// init image variables
	initImages()
	// init font variables
	initFonts()

	// initialize game struct
	game := &game{
		tank: components.NewTank(
			float64(config.ScreenWidth-config.TankWidth)/2,
			float64(config.ScreenHeight-config.TankHeight),
			float64(config.TankSpeed),
			config.TankAmmo,
		),
	}

	// start a go routine to play game soundtrack
	go playSoundTrack()

	// start the game
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.GameTitle)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
