package main

import (
	"fmt"
	p8 "puzzle8"
)

func main() {
	var (
		p = p8.RandomPuzzle(50)
		g = p.Randomize(10)
	)

	fmt.Println("Generated:"); fmt.Println(p)
	fmt.Println("Randomized:"); fmt.Println(g)

	fmt.Println(p.ClosestPathTo(g))
}