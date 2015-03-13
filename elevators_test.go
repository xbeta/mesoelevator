package mesoelevator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewElevators(t *testing.T) {
	const numElevators = 16
	elevators := NewElevators(numElevators)
	if assert.NotNil(t, elevators) {
		assert.Equal(t, numElevators, len(elevators))
	}
}

func TestFindId(t *testing.T) {
	const numElevators = 4
	elevators := NewElevators(numElevators)
	if assert.NotNil(t, elevators) {
		assert.Equal(t, 3, elevators.FindId(4))
		assert.Equal(t, 1, elevators.FindId(2))
	}
}

func TestFlatten(t *testing.T) {
	const numElevators = 4
	elevators := NewElevators(numElevators)
	if assert.NotNil(t, elevators) {
		assert.Equal(t, numElevators, len(elevators))

		// build the test slice to check against
		row := 4
		checkSlice := make([][]int, row)
		for i := range checkSlice {
			checkSlice[i] = []int{i+1, 0, 1}
		}
		assert.Equal(t, checkSlice,elevators.Flatten() )
	}
}
