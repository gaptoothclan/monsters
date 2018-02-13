package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var cityMap = flag.String("citymap", "world_small_map.txt", "The name of the map you want to load")
var monsters = flag.Int("monsters", 100, "The number of monsters you would like to deploy")
var maxMonsterMoves = flag.Int("maxmoves", 100000, "The maximum amount of moves you want to make")

func main() {
	flag.Parse()

	world := NewWorld(*maxMonsterMoves)

	file, err := os.Open(fmt.Sprintf("./data/%s", *cityMap))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	world.LoadCities(file)
	world.UnleashMonsters(*monsters, *maxMonsterMoves)
	world.RunSimulation()
}
