package main

import "time"

const sleepSeconds = 1

type Game struct {
	tick int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Run() {
	for {
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
		println(g.tick)
		g.tick += 1
	}
}
