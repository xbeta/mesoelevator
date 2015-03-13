package mesoelevator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	deque := NewDeque()
	deque.Append(1)
	deque.Append(2)
	deque.Append(3)
	if (assert.NotNil(t, deque) && (assert.Equal(t, 3, deque.Size()))){
		checkList := []interface{}{1,2,3}
		assert.Equal(t, checkList, deque.List())
	}
}
