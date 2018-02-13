package main

import "fmt"

const ALIVE = "alive"
const DEAD = "dead"

type Monster struct {
	Name     string
	Status   string
	moves    int
	maxMoves int
}

// NewMonster create a new monster
func NewMonster(name string, max int) *Monster {
	return &Monster{
		Name:     name,
		Status:   ALIVE,
		moves:    0,
		maxMoves: max,
	}
}

// TryMove try to move the monster, increase moves if success
func (m *Monster) TryMove() bool {
	if m.moves == m.maxMoves {
		m.Status = DEAD
		return false
	}

	return true
}

// GetMessage
func (m *Monster) GetMessage() string {
	return fmt.Sprintf("monster `%s`", m.Name)
}

// KillMonster
func (m *Monster) KillMonster() string {
	m.Status = DEAD
	return m.GetMessage()
}
