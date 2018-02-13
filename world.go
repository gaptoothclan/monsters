package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// World play ground of monsters
type World struct {
	monsters map[string]*Monster
	cities   map[string]*City
	maxMoves int
	curMoves int
}

// NewWorld create a new world
func NewWorld(maxMoves int) *World {
	return &World{
		monsters: make(map[string]*Monster),
		cities:   make(map[string]*City),
		maxMoves: maxMoves,
		curMoves: 0,
	}
}

// LoadCities build cities from a file
func (w *World) LoadCities(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	data := [][]string{}
	for scanner.Scan() {
		dataLine := strings.Split(scanner.Text(), " ")
		data = append(data, dataLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false
	}

	// Temp City data
	tempCityData := make(map[string][]string)
	for _, cityData := range data {
		// Add reference to world into each city so they can notify later
		w.cities[cityData[0]] = NewCity(cityData[0], w)
		tempCityData[cityData[0]] = cityData[0:]
	}

	for cityName, cityData := range tempCityData {
		for _, cityLink := range cityData {
			cityLinkSplit := strings.Split(cityLink, "=")
			switch cityLinkSplit[0] {
			case NORTH:
				w.cities[cityName].Paths.North = w.cities[cityLinkSplit[1]]
				break
			case EAST:
				w.cities[cityName].Paths.East = w.cities[cityLinkSplit[1]]
				break
			case SOUTH:
				w.cities[cityName].Paths.South = w.cities[cityLinkSplit[1]]
				break
			case WEST:
				w.cities[cityName].Paths.West = w.cities[cityLinkSplit[1]]
				break
			}

		}
	}

	return true
}

// UnleashMonsters unleash those beasts Mwah Ha Ha!!!
func (w *World) UnleashMonsters(hordeSize int, monsterMaxMoves int) {
	// Build quick slice of city names
	cityLen := len(w.cities)
	citySlice := []string{}
	for cityName := range w.cities {
		citySlice = append(citySlice, cityName)
	}
	rand.Seed(time.Now().Unix())

	// Build the Horde
	for i := 0; i < hordeSize; i++ {
		monsterName := fmt.Sprintf("MONSTER_%d", i)
		monster := NewMonster(monsterName, monsterMaxMoves)
		w.monsters[monsterName] = monster

		// Asign to random city
		random := rand.Intn(cityLen)
		w.cities[citySlice[random]].AddMonster(monster)
	}
}

// MonsterDeathNotification
// Find city and remove
func (w *World) MonsterDeathNotification(monster *Monster) {
	if w.monsters[monster.Name] != nil {
		delete(w.monsters, monster.Name)
	}
}

// CityDeathNotification
// Find city and remove
func (w *World) CityDeathNotification(city *City) {
	if w.cities[city.Name] != nil {
		delete(w.cities, city.Name)
	}
}

// Run simulation
func (w *World) RunSimulation() {
	for {
		if len(w.cities) == 0 {
			fmt.Println("No more cities")
			break
		}

		if len(w.monsters) == 0 {
			fmt.Println("No more monsters")
			break
		}

		if w.curMoves == w.maxMoves {
			fmt.Println("Max moves exceeded")
			break
		}

		for cityName := range w.cities {
			w.cities[cityName].Check()
			if w.cities[cityName] != nil {
				w.cities[cityName].MoveMonsters()
			}
		}

		w.curMoves = w.curMoves + 1
	}
	fmt.Println("Simulation Ended Write out current world status")
	w.GetWorldStatus()
}

// GetWorldStatus
func (w *World) GetWorldStatus() {
	for cityName := range w.cities {
		fmt.Printf("%s\n", w.cities[cityName].ToString())
	}
}
