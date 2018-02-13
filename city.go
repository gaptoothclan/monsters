package main

import (
	"fmt"
	"strings"
)

// DeathNotifcation
type DeathNotifcation interface {
	MonsterDeathNotification(monster *Monster)
	CityDeathNotification(city *City)
}

// City
type City struct {
	Name         string
	Destroyed    bool
	monsterIn    []*Monster
	monsters     []*Monster
	Paths        Paths
	notification DeathNotifcation
}

// NewCity
func NewCity(name string, notification DeathNotifcation) *City {
	return &City{
		Name:         name,
		monsterIn:    []*Monster{},
		monsters:     []*Monster{},
		notification: notification,
	}
}

// Inform all connected cities of desctruction
// Write message
func (c *City) destroy() {
	// Check each path
	if c.Paths.North != nil {
		c.Paths.North.Notify(c, NORTH)
	}
	if c.Paths.East != nil {
		c.Paths.East.Notify(c, EAST)
	}
	if c.Paths.South != nil {
		c.Paths.South.Notify(c, SOUTH)
	}
	if c.Paths.West != nil {
		c.Paths.West.Notify(c, WEST)
	}

	monsterMessages := make([]string, len(c.monsters))
	for key, monster := range c.monsters {
		monsterMessages[key] = monster.GetMessage()
		c.notification.MonsterDeathNotification(monster)
	}

	monsterMessage := strings.Join(monsterMessages[:], " and ")
	fmt.Printf("City `%s` has been destroyed by %s\n", c.Name, monsterMessage)

	c.Destroyed = true
	c.notification.CityDeathNotification(c)
}

// Notify allow another city to notify me that it is dead
// We will kill the path to them
// Notification comes from the oppisite so destroy link
func (c *City) Notify(city *City, from string) {

	switch from {
	case NORTH:
		if c.Paths.South == city {
			c.Paths.South = nil
		}
		break
	case EAST:
		if c.Paths.West == city {
			c.Paths.West = nil
		}
		break
	case SOUTH:
		if c.Paths.North == city {
			c.Paths.North = nil
		}
		break
	case WEST:
		if c.Paths.East == city {
			c.Paths.East = nil
		}
		break
	}

}

// Check how many monsters are in the city
// Move monsters from in to main store
// More than 1 monster then kill destroy
func (c *City) Check() {
	if len(c.monsterIn) == 0 {
		return
	}

	for _, monster := range c.monsterIn {
		c.monsters = append(c.monsters, monster)
	}

	c.monsterIn = []*Monster{}

	if len(c.monsters) > 1 {
		c.destroy()
		return
	}
}

// AddMonster a new monster enters town!!
func (c *City) AddMonster(monster *Monster) {
	c.monsterIn = append(c.monsterIn, monster)
}

// MoveMonsters Send monster on way
// Only one monster should be in the city
// otherwise game rules are running wrong
// If monster can not move, he is going to die
// notify the city and monster as dead
func (c *City) MoveMonsters() {
	// Check monster can move if not monster is dead so remove
	if len(c.monsters) != 1 {
		return
	}

	if c.monsters[0].TryMove() == false {
		// Notify map
		c.notification.MonsterDeathNotification(c.monsters[0])
		return
	}

	if c.Paths.ExpelMonster(c.monsters[0]) == false {
		// Inform map monster was trapped and died
		msg := c.monsters[0].KillMonster()
		c.notification.MonsterDeathNotification(c.monsters[0])
		c.notification.CityDeathNotification(c)
		fmt.Printf("%s became trapped and has now died in %s\n", msg, c.Name)
	}

	// Unset the array
	c.monsters = []*Monster{}
}

// ToString
// List out the paths to the city and return string
func (c *City) ToString() string {
	paths := []string{}
	if c.Paths.North != nil {
		paths = append(paths, fmt.Sprintf("north=%s", c.Paths.North.Name))
	}
	if c.Paths.East != nil {
		paths = append(paths, fmt.Sprintf("east=%s", c.Paths.East.Name))
	}
	if c.Paths.South != nil {
		paths = append(paths, fmt.Sprintf("south=%s", c.Paths.South.Name))
	}
	if c.Paths.West != nil {
		paths = append(paths, fmt.Sprintf("west=%s", c.Paths.West.Name))
	}

	return fmt.Sprintf("%s %s", c.Name, strings.Join(paths, " "))

}
