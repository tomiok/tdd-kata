package tdd_kata

import (
	"testing"
)

/*
The Mars Rover Kata (simplified)
Test-drive code to drive a rover over the idealised surface of Mars, represented as a 2-D grid of x, y
coordinates (e.g., [4, 7]).
The rover can be “dropped” on to the grid at a specific set of coordinates, facing in one of four
directions: North, East, South or West.
Sequences of instructions are sent to the rover telling it where to go. This sequence is a string of
characters, each representing a single instruction:

R = turn right
L = turn left
F = move on square forward
B = move one square back
The rover ignores invalid instructions.
How to tackle this exercise:
1. Make a test list for your rover based on these requirements
2. Work through your test list on test case at a time, remembering to Red-Green-Refactor with
each test
3. When refactoring, don’t forget that test code has to be maintained, too!


example
0,0 , north

	N
O	->	E
	S

0 = n
90 = e
180 = s
270 = o

N -> F = y + 1
N -> B = y - 1
E -> F = x + 1
E -> B = x - 1
*/

func Test_rover_initial_position(t *testing.T) {
	rover := newRover(0, 0, north)

	x, y, direction := rover.getPosition()

	assertEquals(t, x, 0, rover)

	assertEquals(t, y, 0, rover)

	assertEquals(t, direction, north, rover)
}

func Test_rover_move_f_north(t *testing.T) {
	rover := newRover(0, 0, north)

	rover.Command("F")

	x, y, direction := rover.getPosition()

	assertEquals(t, x, 0, rover)

	assertEquals(t, y, 1, rover)

	assertEquals(t, direction, north, rover)

	rover.printPosition()
}

func Test_rover_move_f_east(t *testing.T) {
	rover := newRover(0, 0, east)

	rover.Command("F")

	x, y, direction := rover.getPosition()

	assertEquals(t, x, 1, rover)

	assertEquals(t, y, 0, rover)

	assertEquals(t, direction, east, rover)

	rover.printPosition()
}

func Test_rover_move_f_south(t *testing.T) {
	rover := newRover(0, 0, east)

	rover.Command("F")

	x, y, direction := rover.getPosition()

	assertEquals(t, x, 1, rover)

	assertEquals(t, y, 0, rover)

	assertEquals(t, direction, east, rover)

	rover.printPosition()
}

func assertEquals(t *testing.T, actual, expected int, rover *Rover) {
	if actual != expected {
		rover.printPosition()
		t.FailNow()
	}
}
