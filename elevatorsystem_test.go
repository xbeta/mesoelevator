package mesoelevator

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func PrintStatus(es ElevatorSystem){

	status, ok := es.Status()
	if(ok) {
		fmt.Printf("Status: %v\n", status)
	}
}

func TestNewElevatorSystem(t *testing.T) {
	const numElevators = 8
	elevatorSystem := NewElevatorSystem(8)
	if assert.NotNil(t, elevatorSystem) {
		assert.Equal(t, numElevators, len(elevatorSystem.elevators))
		assert.Equal(t, 0, elevatorSystem.pickupQueue.Size())
	}
}

func TestStatus(t *testing.T) {
	const numElevators = 4
	elevatorSystem := NewElevatorSystem(numElevators)
	if assert.NotNil(t, elevatorSystem) {
		// build the test slice to check against
		row := 4
		checkSlice := make([][]int, row)
		for i := range checkSlice {
			checkSlice[i] = []int{i+1, 0, 1}
		}

		slice, ok := elevatorSystem.Status()
		assert.Equal(t, true, ok)
		assert.Equal(t, checkSlice, slice)
	}
}

func TestStep(t *testing.T) {
	const numElevators = 2
	elevatorSystem := NewElevatorSystem(numElevators)
	if assert.NotNil(t, elevatorSystem) {

		elevatorSystem.Step()
		elevatorSystem.Pickup(2, 1)
		elevatorSystem.Step()
		PrintStatus(elevatorSystem)

		elevatorSystem.Step()

		elevatorSystem.Pickup(3, -1)
		elevatorSystem.Pickup(3, 1)
		elevatorSystem.Pickup(1, 1)

		elevatorSystem.Step()
		PrintStatus(elevatorSystem)

		elevatorSystem.Step()
		PrintStatus(elevatorSystem)

		elevatorSystem.Step()
		PrintStatus(elevatorSystem)

		elevatorSystem.Step()
		PrintStatus(elevatorSystem)

		elevatorSystem.Pickup(5, -1)
		elevatorSystem.Step()
		PrintStatus(elevatorSystem)


		elevatorSystem.Step()
		elevatorSystem.Step()

		elevatorSystem.Step()
		elevatorSystem.Step()
		elevatorSystem.Step()
		elevatorSystem.Step()
		elevatorSystem.Step()

		PrintStatus(elevatorSystem)

	}
}
