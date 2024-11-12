package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func (p *Player) Update(speed float64) {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if p.position.Y < 705 {
			p.position.Y += speed * 2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if p.position.Y > 0 {
			p.position.Y -= speed * 2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.position.X > 420 {
			p.position.X -= speed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.position.X < 570 {
			p.position.X += speed
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opP := &ebiten.DrawImageOptions{}
	opP.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, opP)
}

func (p *Player) IsCrash(o *Obstacle) bool {
	obstacleW := float64(o.sprite.Bounds().Dx())
	obstacleH := float64(o.sprite.Bounds().Dy())
	if p.position.X < o.position.X+obstacleW && p.position.X+obstacleW > o.position.X &&
		p.position.Y < o.position.Y+obstacleH && p.position.Y+obstacleH > o.position.Y {
		return true
	}

	return false
}
