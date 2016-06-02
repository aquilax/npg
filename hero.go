package main

const (
	initialExperience = 0
)

type Hero struct {
	Entity
	Experience int
}

func NewHero() *Hero {
	var hero Hero
	hero.Experience = initialExperience
	hero.Type = typeHero
	return &hero
}
