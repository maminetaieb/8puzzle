package main

import (
	"fmt"
	p8 "puzzle8"
)

func main() {
	var (
		p = p8.RandomPuzzle(50)
		g = p8.Shuffle(p, 10)
	)

	fmt.Println("Generated:"); fmt.Println(p)
	fmt.Println("Shuffled:"); fmt.Println(g)

	p.Trace(p.ClosestPathTo(g))
}