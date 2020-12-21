package puzzle8

import (
	"fmt"
	"math/rand"
	"time"
)

type Puzzle struct {
	board [3][3]int
	x int
	y int
}

// Creating puzzles
func NewPuzzle(init [3][3]int) Puzzle {
	var x, y int

	Loop:
	for i, row := range init {
		for j, val := range row {
			if val == 0 {
				x, y = i, j
				break Loop
			}
		}
	}

	for i := x; i < len(init); i++ {
		for j := y+1; j < len(init[i]); j++ {
			if init[i][j] == 0 {
				panic("Multiple Zeros")
			}
		}
	}

	return Puzzle {init, x, y}
}

func RandomPuzzle(l int) Puzzle {
	rand.Seed(time.Now().UnixNano())

	p := Puzzle {
		x: rand.Intn(3),
		y: rand.Intn(3),
	}

	for i := range(p.board) {
		for j := range(p.board[i]) {
			p.board[i][j] = rand.Intn(l) + 1
		}
	}

	p.board[p.x][p.y] = 0

	return p
}

// Manipulating puzzles
func Shuffle(p Puzzle, times int) Puzzle {
	for times > 0 {
		switch rand.Intn(4) {
		case 0:
			if n, ok := MoveUp(p); !ok {
				continue
			} else {
				p = n
				times--
			}
		case 1:
			if n, ok := MoveDown(p); !ok {
				continue
			} else {
				p = n
				times--
			}
		case 2:
			if n, ok := MoveLeft(p); !ok {
				continue
			} else {
				p = n
				times--
			}
		case 3:
			if n, ok := MoveRight(p); !ok {
				continue
			} else {
				p = n
				times--
			}
		}
	}
	
	return p
}

func MoveUp(p Puzzle) (Puzzle, bool) {
	if p.x == 0 {
		return p, false
	}
	
	p.board[p.x][p.y], p.board[p.x-1][p.y] = p.board[p.x-1][p.y], p.board[p.x][p.y]
	p.x--

	return p, true
}

func MoveDown(p Puzzle) (Puzzle, bool) {
	if p.x == len(p.board)-1 {
		return p, false
	}
	
	p.board[p.x][p.y], p.board[p.x+1][p.y] = p.board[p.x+1][p.y], p.board[p.x][p.y]
	p.x++

	return p, true
}

func MoveLeft(p Puzzle) (Puzzle, bool) {
	if p.y == 0 {
		return p, false
	}
	
	p.board[p.x][p.y], p.board[p.x][p.y-1] = p.board[p.x][p.y-1], p.board[p.x][p.y]
	p.y--

	return p, true
}

func MoveRight(p Puzzle) (Puzzle, bool) {
	if p.y == len(p.board)-1 {
		return p, false
	}
	
	p.board[p.x][p.y], p.board[p.x][p.y+1] = p.board[p.x][p.y+1], p.board[p.x][p.y]
	p.y++

	return p, true
}

// Other utilities
func (p Puzzle) Equals(g Puzzle) bool {
	if p.x != g.x || p.y != g.y {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (p.board[i][j] != g.board[i][j]) {
				return false
			}
		}
	}

	return true
}

func (p Puzzle) String() (s string) {
	s += fmt.Sprintf("x:%d\t\ty:%d\n", p.x, p.y)
	for _, row := range(p.board) {
		for j := range(row) {
			s += fmt.Sprintf("%d\t", row[j])
		}
		s += "\n"
	}

	return
}