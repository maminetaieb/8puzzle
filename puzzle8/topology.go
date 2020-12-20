package puzzle8

import "fmt"

type node struct {
	pzl Puzzle
	actions string
	nextNode *node		// horizontally
	nextLineNode *node	// Vertically
}
func (p Puzzle) GetPathTo(goal Puzzle) (s string) {
	n := (&node{p, "", nil, nil}).genNextLine(goal, true)
	s = fmt.Sprintf("%d actions performed to reach goal\n", len(n.actions))
	for i:=0; i < len(n.actions); i++ {
		switch n.actions[i] {
		case 'U':
			s += fmt.Sprintln("0 Moved up.")
		case 'D':
			s += fmt.Sprintln("0 Moved down.")
		case 'L':
			s += fmt.Sprintln("0 Moved left.")
		case 'R':
			s += fmt.Sprintln("0 Moved right.")
		}
	}
	return
}

func (parent *node) genNextLine(goal Puzzle, isLastNodeInline bool) *node {
	if up, okUp := MoveUp(parent.pzl); okUp {
		u := &node {up, parent.actions+"U", nil, nil}
		parent.nextLineNode = u
		if u.pzl.Equals(goal) {
			return u
		}
	}
	if down, okDown := MoveDown(parent.pzl); okDown {
		d := &node {down, parent.actions+"D", nil, nil}
		if (parent.nextLineNode == nil) {
			parent.nextLineNode = d
		} else {
			parent.nextLineNode.nextNode = d
		}
		if d.pzl.Equals(goal) {
			return d
		}
	}
	if left, okLeft := MoveLeft(parent.pzl); okLeft {
		l := &node {left, parent.actions+"L", nil, nil}
		var curr *node
		for curr = parent.nextLineNode; curr.nextNode != nil; {
			curr = curr.nextNode
		}
		curr.nextNode = l
		if l.pzl.Equals(goal) {
			return l
		}
	}
	if right, okRight := MoveRight(parent.pzl); okRight {
		r := &node {right, parent.actions+"R", nil, nil}
		var curr *node
		for curr = parent.nextLineNode; curr.nextNode != nil; {
			curr = curr.nextNode
		}
		curr.nextNode = r
		if r.pzl.Equals(goal) {
			return r
		}
	}

	if (isLastNodeInline) {
		return parent.nextLineNode.genNextLine(goal, isLastNodeInline && parent.nextLineNode.nextNode == nil)
	} else {
		return parent.nextNode.genNextLine(goal, isLastNodeInline && parent.nextNode.nextNode == nil)
	}
}