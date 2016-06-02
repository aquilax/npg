package main

const (
	initialExperience = 0
)

type Hero struct {
	Entity
	Experience int
}

func NewHero() *Hero {
	return &Hero{
		Experience: initialExperience,
		Type:       typeHero,
	}
}
