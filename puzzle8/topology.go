package puzzle8

import (
	"fmt"
)

// Actions performed on puzzles
type actions string
type action byte
const (
	up action = 'U'
	down = 'D'
	left = 'L'
	right = 'R'
)

// Queue
type qElement struct {
	rep Puzzle
	act action
	next *qElement
}
type queue struct {
	head *qElement
	tail *qElement
}
var q *queue

func pend(p Puzzle, act action) {	// aka add element
	q.tail.next = &qElement{p, act, nil}
	q.tail = q.tail.next
}

func popAction() action {
	defer func() {
		q.head = q.head.next
		if (q.head == nil) {
			q.tail = nil
		}
	}()

	return q.head.act
}

// Searching path
func (p Puzzle) ClosestPathTo(g Puzzle) actions {
	init := &qElement{p, ' ', nil}
	q = &queue{init, init}

	a := continueUntil(g)
	a = actions(a[1:])
	return a
}

func continueUntil(g Puzzle) actions {
	p, act := q.head.rep, q.head.act

	if p.Equals(g) {
		return actions(act)
	}
	if newPzl, okNew := MoveUp(p); okNew && act != down {
		pend(newPzl, 'U')
	}
	if newPzl, okNew := MoveDown(p); okNew && act != up {
		pend(newPzl, 'D')
	}
	if newPzl, okNew := MoveLeft(p); okNew && act != right {
		pend(newPzl, 'L')
	}
	if newPzl, okNew := MoveRight(p); okNew && act != left {
		pend(newPzl, 'R')
	}

	return actions(popAction()) + continueUntil(g)
}

// Printing
func (a actions) String() string {
	s := fmt.Sprintln("At least", len(a), "actions needed:")
	for i := 0; i < len(a); i++ {
		switch action(a[i]) {
		case up:
			s += fmt.Sprintln(i+1, "- Up")
		case down:
			s += fmt.Sprintln(i+1, "- Down")
		case left:
			s += fmt.Sprintln(i+1, "- Left")
		case right:
			s += fmt.Sprintln(i+1, "- Right")
		}
	}

	return s
}

func (p Puzzle) Trace(a actions) {
	fmt.Println(p)
	if len(a) != 0 {
		switch action(a[0]) {
		case up:
			np, _ := MoveUp(p)
			np.Trace(actions(a[1:]))
		case down:
			np, _ := MoveDown(p)
			np.Trace(actions(a[1:]))
		case left:
			np, _ := MoveLeft(p)
			np.Trace(actions(a[1:]))
		case right:
			np, _ := MoveRight(p)
			np.Trace(actions(a[1:]))
		}
	}
}