package main

import (
	"fmt"
	p8 "puzzle8"
)

func main() {
	var (
		p = p8.RandomPuzzle(50)
		g = p.Randomize(50)
	)
	fmt.Println("Generated", p)
	fmt.Println("Randomized to", g)
	fmt.Println(p.GetPathTo(g))
}