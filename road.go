package main

import "github.com/hajimehoshi/ebiten/v2"

type Road struct {
	positionL Vector
	positionR Vector
	spriteL   *ebiten.Image
	spriteR   *ebiten.Image
}

func (r *Road) Update(level float64) {
	r.positionL.Y += carSpeed + level
	if r.positionL.Y >= screenHeight {
		r.positionR.Y = 0
	}

	r.positionR.Y += carSpeed + level
	if r.positionR.Y >= screenHeight {
		r.positionR.Y = 0
	}
}

func (r *Road) Draw(screen *ebiten.Image) {
	topY := r.positionL.Y
	bottomY := r.positionL.Y
	spriteH := float64(r.spriteL.Bounds().Dy())

	for {
		opL := &ebiten.DrawImageOptions{}
		opL.GeoM.Translate(r.positionL.X, topY)
		screen.DrawImage(r.spriteL, opL)

		opR := &ebiten.DrawImageOptions{}
		opR.GeoM.Translate(r.positionR.X, topY)
		screen.DrawImage(r.spriteR, opR)

		if topY > screenHeight {
			break
		}

		topY += spriteH
	}

	for {
		opL := &ebiten.DrawImageOptions{}
		opL.GeoM.Translate(r.positionL.X, bottomY)
		screen.DrawImage(r.spriteL, opL)

		opR := &ebiten.DrawImageOptions{}
		opR.GeoM.Translate(r.positionR.X, bottomY)
		screen.DrawImage(r.spriteR, opR)

		if bottomY < 0 {
			break
		}

		bottomY -= spriteH
	}
}
