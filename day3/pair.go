package main

import "fmt"

type Pair struct {
	x, y int
}

func (p Pair) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
