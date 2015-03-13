package mesoelevator

import (
  "fmt"
  )

// interface
type ElevatorControlSystem interface {
  Status() ([][]int, bool)
  Update(id int, currentFloor int, goalFloor int) (bool)
  Pickup(pickupFloor int, direction int) (bool)
  Step() (bool)
}

// all structs
type ElevatorSystem struct{
  elevators Elevators
  pickupQueue *Deque
}

// NOTE: not thread-safe, but it does not need to be thread-safe.
// It's simple for loop to initialize the list.  This should get call at the very beginning to setup
// the elevator system.
func NewElevatorSystem(size int)(ElevatorSystem){
  es := NewElevators(size)
  pq := NewDeque()
  return ElevatorSystem{
    elevators: es,
    pickupQueue: pq,
  }
}

// NOTE: not thread-safe, but it does not need to be thread-safe.  It's a read-only function and
// should only be used for reporting purpose.
//
// def status(): Seq[(Int, Int, Int)]
func (es ElevatorSystem) Status()([][]int, bool){
  if (len(es.elevators) >= 0){
    return es.elevators.Flatten(), true
  }
  return nil, false
}

// NOTE: not thread-safe
// def update(Int, Int, Int)
func (es ElevatorSystem) Update(id int, currentFloor int )(bool) {

  // locate the elevator inside a list of elevators
  foundIndex := es.elevators.FindId(id)
  if foundIndex <0 { return false }
  foundElevator := es.elevators[foundIndex]

  if (foundElevator.goalFloors.Empty() == false) {
    nextFloor := foundElevator.removeFloor()

    if (foundElevator.currentFloor < nextFloor) {
      foundElevator.direction = 1 // elevator needs go up
    } else if (foundElevator.currentFloor > nextFloor) {
      foundElevator.direction = -1 // elevator needs go down
    } else {
      // next floor is the same as the current floor, stays
      foundElevator.direction = 0
    }
    foundElevator.currentFloor = nextFloor
  } else {
    // when no other goal floors left, let's move back to ground floor in idle
    foundElevator.currentFloor = 0
    foundElevator.direction = 1
  }
  return true
}

// NOTE: not thread-safe
// def pickup(Int, Int)
func (es ElevatorSystem) Pickup(pickupFloor int, direction int) (bool) {
  return es.pickupQueue.Append(Pickup{pickupFloor, direction})
}

func (es ElevatorSystem) Unload(elevator *Elevator)() {
  pickup := es.pickupQueue.First().(Pickup) //peak first element
  fmt.Printf("Elevator %d will pickup at floor %d\n", elevator.id, pickup.floor)

  // when elevator at ground floor & no queue, always pick them up!
  if (elevator.currentFloor == 0 && elevator.goalFloors.Empty()){
    elevator.addFloor(pickup.floor)
    es.pickupQueue.Shift()
    return
  }

  switch pickup.direction {
    case -1:
      if (elevator.direction == -1 && elevator.currentFloor >= pickup.floor) {
        elevator.addFloor(pickup.floor)
        es.pickupQueue.Shift()
        return
      }
    case 1:
      if (elevator.direction == 1 && elevator.currentFloor <= pickup.floor) {
        elevator.addFloor(pickup.floor)
        es.pickupQueue.Shift()
        return
      }
  }
}

func (es ElevatorSystem) Step() () {
  // first allocate all the pickup requests to all of the elevators before the actual stepping
  for _, elevator := range es.elevators {
    if (es.pickupQueue.Empty() == false) {
      es.Unload(elevator)
    } else {
      // now actually stepping it
      es.Update(elevator.id, elevator.currentFloor)
    }
  }
}
