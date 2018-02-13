package main

import (
	"math/rand"
	"time"
)

const (
	NORTH = "north"
	EAST  = "east"
	SOUTH = "south"
	WEST  = "west"
)

// Paths
type Paths struct {
	North *City
	East  *City
	South *City
	West  *City
}

// ExpelMonster off you pop
// Get all paths, if there are no paths
// Monster will die
func (p *Paths) ExpelMonster(monster *Monster) bool {
	// Choose random path
	// If no path return false
	paths := p.getCurrentPaths()

	pathLen := len(paths)

	if pathLen == 0 {
		return false
	}

	// Pink one at random
	rand.Seed(time.Now().Unix())
	random := rand.Intn(pathLen)

	p.sendMonsterPacking(monster, paths[random])

	return true
}

// sendMonsterPacking
// Send the monster along a path to a city
func (p *Paths) sendMonsterPacking(monster *Monster, direction string) {
	switch direction {
	case NORTH:
		p.North.AddMonster(monster)
		return
	case EAST:
		p.East.AddMonster(monster)
		return
	case SOUTH:
		p.South.AddMonster(monster)
		return
	case WEST:
		p.West.AddMonster(monster)
		return
	}
	panic("You have tried to send a direction that does not exist")
}

// getCurrentPaths
// Get all live paths
func (p *Paths) getCurrentPaths() []string {
	paths := []string{}
	if p.North != nil {
		paths = append(paths, NORTH)
	}
	if p.East != nil {
		paths = append(paths, EAST)
	}
	if p.South != nil {
		paths = append(paths, SOUTH)
	}
	if p.West != nil {
		paths = append(paths, WEST)
	}
	return paths
}
