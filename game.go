package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"math/rand"
)

type CarGame struct {
	player   *Player
	land     *Land
	road     *Road
	obstacle *Obstacle
	timer    *Timer
	isCrash  bool
	level    float64
	speed    int
}

func (g *CarGame) Reset() {
	g.isCrash = false
	g.player.position = Vector{X: 500, Y: 650}
	g.obstacle.position = Vector{X: -39, Y: -65}
	g.obstacle.visible = false
	//g.level = 0.1 // reset level
}

func (g *CarGame) Update() error {
	if g.isCrash {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Reset()
			g.timer.Reset()
		}
	}

	if !g.isCrash {
		g.timer.Update(g.level)
		if g.timer.IsReady() {
			g.timer.Reset()
			g.level += 1

			if g.obstacle.visible == false {
				g.obstacle.position.X = float64(rand.Intn(570-420) + 420)
				g.obstacle.visible = true
			}
		}

		if g.obstacle.visible == true {
			g.obstacle.Update(g.level)
		}

		speed := 5.0

		g.land.Update(g.level)
		g.road.Update(g.level)
		g.player.Update(speed)

		if g.player.IsCrash(g.obstacle) {
			g.isCrash = true
		}
	}

	return nil
}

func (g *CarGame) Draw(screen *ebiten.Image) {
	g.land.Draw(screen)
	g.road.Draw(screen)
	g.obstacle.Draw(screen)
	g.player.Draw(screen)

	if g.isCrash {
		text.Draw(screen, "!!!! CRASH !!!!", ScoreFont, screenWidth/2-200, 200, color.RGBA{R: 255, G: 0, B: 0})
		text.Draw(screen, "Press \n`Enter`", ScoreFont, 70, 300, color.RGBA{R: 255, G: 0, B: 0})
		text.Draw(screen, "to \ncontinue", ScoreFont, 650, 350, color.RGBA{R: 255, G: 0, B: 0})
	}
}

func (g *CarGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
