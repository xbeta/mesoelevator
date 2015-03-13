package mesoelevator

import (
	"fmt"
)

type Elevator struct {
	id           int
	currentFloor int         // default 0
	direction    int         // -1:  down, 0: nodirection/idle,  1: up
	goalFloors   *Deque // all floors that this elevator needs to visit, floor number starts with positive 1
}

func NewElevator(id int) (*Elevator) {
	deque := NewDeque()
	elevator := Elevator{
		id:           id,
		direction:    1,
		currentFloor: 0,
		goalFloors:   deque}
	return &elevator
}

func (e *Elevator) addFloor(floor int) {
	e.goalFloors.Append(floor)
}

func (e *Elevator) removeFloor() (int) {
	floor := e.goalFloors.Shift()
	return floor.(int) //assuming the type is always int for now
}

func (e *Elevator) String() string {
	return fmt.Sprintf("%d: {%d, %d, %v}", e.id, e.currentFloor, e.direction, e.goalFloors)
}
