package main

import "github.com/hajimehoshi/ebiten/v2"

type Obstacle struct {
	position Vector
	sprite   *ebiten.Image
	visible  bool
}

func (o *Obstacle) Update(level float64) {
	o.position.Y += carSpeed + level/2
	if o.position.Y > screenHeight {
		o.position.Y = -65
		o.visible = false
	}
}

func (o *Obstacle) Draw(screen *ebiten.Image) {
	opO := &ebiten.DrawImageOptions{}
	opO.GeoM.Translate(o.position.X, o.position.Y)
	screen.DrawImage(o.sprite, opO)
}
