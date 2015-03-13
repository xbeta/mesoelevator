package mesoelevator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewElevator(t *testing.T) {
	elevator := NewElevator(1)
	if assert.NotNil(t, elevator) {
		assert.Equal(t, 1, elevator.id)
		assert.Equal(t, 1, elevator.direction)
		assert.Equal(t, 0, elevator.currentFloor)
		assert.NotNil(t, elevator.goalFloors)
		assert.Equal(t, 0, elevator.goalFloors.Size())
	}
}

func TestElevatorAddFloor(t *testing.T) {
	elevator := NewElevator(1)
	elevator.addFloor(9)
	elevator.addFloor(10)
	assert.NotNil(t, elevator.goalFloors)
	assert.Equal(t, 2, elevator.goalFloors.Size())
	assert.Equal(t, 9, elevator.goalFloors.First())
}

func TestElevatorRemoveFloor(t *testing.T) {
	elevator := NewElevator(1)
	elevator.addFloor(9)
	elevator.addFloor(10)
	elevator.addFloor(11)
	floor := elevator.removeFloor()
	assert.NotNil(t, elevator.goalFloors)
	assert.Equal(t, 2, elevator.goalFloors.Size())
	assert.Equal(t, 9, floor)
}
