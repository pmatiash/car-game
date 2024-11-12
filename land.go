package main

import "github.com/hajimehoshi/ebiten/v2"

type Land struct {
	position Vector
	sprite   *ebiten.Image
}

func (l *Land) Update(level float64) {
	l.position.Y += carSpeed + level
	if l.position.Y >= screenHeight {
		l.position.Y = 0
	}
}

func (l *Land) Draw(screen *ebiten.Image) {
	spriteW := l.sprite.Bounds().Dx()
	spriteH := float64(l.sprite.Bounds().Dy())
	for y := l.position.Y; y <= screenHeight; y += spriteH {
		for x := 0; x <= screenWidth; x += spriteW {
			opD := &ebiten.DrawImageOptions{}
			opD.GeoM.Translate(float64(x), y)
			screen.DrawImage(l.sprite, opD)
		}
	}

	for y := l.position.Y; y > (0 - spriteH); y -= spriteH {
		for x := 0; x <= screenWidth; x += spriteW {
			opD := &ebiten.DrawImageOptions{}
			opD.GeoM.Translate(float64(x), y)
			screen.DrawImage(l.sprite, opD)
		}
	}
}
