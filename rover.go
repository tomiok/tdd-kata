package tdd_kata

import "fmt"

const north = 0
const east = 90

type Rover struct {
	x         int
	y         int
	direction int
}

func (r *Rover) getPosition() (int, int, int) {
	return r.x, r.y, r.direction
}

func (r *Rover) printPosition() {
	fmt.Println(fmt.Sprintf("x:%d, y:%d, degrees:%d", r.x, r.y, r.direction))
}

func (r *Rover) Command(s string) {
	o := r.getOrientation()

	switch o {
	case north:
		r.y = r.y + 1
	case east:
		r.x = r.x + 1
	}
}

func (r *Rover) getOrientation() int {
	return r.direction
}

func newRover(x int, y int, direction int) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}
