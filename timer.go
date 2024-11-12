package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type Timer struct {
	currentTicks int
	targetTicks  float64
	level        float64
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  float64(int(d.Milliseconds()) * ebiten.TPS() / 1000),
	}
}

func (t *Timer) Update(level float64) {
	if t.currentTicks < int(t.targetTicks) {
		fmt.Println("level", level)
		if math.Mod(level, 10.0) == 0 {
			if t.level < level {
				t.level = level
				t.targetTicks = t.targetTicks * 0.5
				fmt.Println("targetTicks", t.targetTicks)
			}
		}

		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= int(t.targetTicks)
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}
