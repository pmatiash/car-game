package main

import (
	"embed"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"log"
	"runtime/debug"
	"time"
)

//go:embed assets/*
var assets embed.FS

var ScoreFont = mustLoadFont("assets/fonts/font.ttf")

type Vector struct {
	X float64
	Y float64
}

func main() {
	debug.SetGCPercent(200)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, Diana Game!")
	ebiten.SetTPS(30)
	//ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetVsyncEnabled(false)

	if err := ebiten.RunGame(&CarGame{
		player: &Player{
			position: Vector{X: 500, Y: 650},
			sprite:   mustLoadImage("assets/cars/car_red_small_5.png"),
		},
		land: &Land{
			position: Vector{X: 0, Y: 0},
			sprite:   mustLoadImage("assets/tiles/dirt/land_dirt12.png"),
		},
		road: &Road{
			positionL: Vector{X: 400, Y: 0},
			positionR: Vector{X: 500, Y: 0},
			spriteL:   mustLoadImage("assets/tiles/asphalt/road_asphalt21.png"),
			spriteR:   mustLoadImage("assets/tiles/asphalt/road_asphalt23.png"),
		},
		obstacle: &Obstacle{
			position: Vector{X: -39, Y: -65},
			sprite:   mustLoadImage("assets/cars/car_black_small_5.png"),
			visible:  false,
		},
		timer:   NewTimer(2 * time.Second),
		isCrash: false,
		level:   1,
		speed:   carSpeed,
	}); err != nil {
		log.Fatal(err)
	}
}
