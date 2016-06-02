package main

type EntityType int

const (
	typeHero EntityType = iota
)

type Entity struct {
	Name   string
	Health int
	Type   EntityType
}
